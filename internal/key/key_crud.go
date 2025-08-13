package key

import (
	"context"
	"fmt"
	"net/http"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

func createKey(ctx context.Context, c *litellm.Client, request *KeyGenerateRequest) (*KeyGenerateResponse, error) {
	response, err := litellm.SendRequestTyped[KeyGenerateRequest, KeyGenerateResponse](
		ctx, c, http.MethodPost, "/key/generate", request,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create key: %w", err)
	}

	return response, nil
}

func getKey(ctx context.Context, c *litellm.Client, keyID string) (*KeyInfoResponse, error) {
	response, err := litellm.SendRequestTyped[interface{}, KeyInfoResponse](
		ctx, c, http.MethodGet, fmt.Sprintf("/key/info?key=%s", keyID), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get key: %w", err)
	}

	return response, nil
}

func updateKey(ctx context.Context, c *litellm.Client, keyID string, request *KeyGenerateRequest) (*KeyGenerateResponse, error) {
	// Add the key ID to the request for updates
	updateRequest := *request
	updateRequest.Key = &keyID

	response, err := litellm.SendRequestTyped[KeyGenerateRequest, KeyGenerateResponse](
		ctx, c, http.MethodPost, "/key/update", &updateRequest,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update key: %w", err)
	}

	return response, nil
}

func deleteKey(ctx context.Context, c *litellm.Client, keyID string) error {
	deleteRequest := struct {
		Keys []string `json:"keys"`
	}{
		Keys: []string{keyID},
	}

	_, err := litellm.SendRequestTyped[struct {
		Keys []string `json:"keys"`
	}, interface{}](
		ctx, c, http.MethodPost, "/key/delete", &deleteRequest,
	)
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}

	return nil
}

// convertInfoResponseToKey converts a KeyInfoResponse to the internal Key struct
func convertInfoResponseToKey(response *KeyInfoResponse) *Key {
	if response == nil {
		return nil
	}

	info := response.Info
	key := &Key{
		Key:      response.Key, // Use the key from the top level
		KeyAlias: safeStringDeref(info.KeyAlias),
		KeyName:  info.KeyName,
		Token:    response.Key, // The key field in info response is the token
		TokenID:  response.Key, // Use key as token ID for consistency
		BudgetID: safeStringDeref(info.BudgetID),

		Models:               info.Models,
		Duration:             "", // Not available in info response
		UserID:               info.UserID,
		TeamID:               safeStringDeref(info.TeamID),
		MaxParallelRequests:  safeIntDeref(info.MaxParallelRequests),
		Metadata:             info.Metadata,
		TPMLimit:             safeIntDeref(info.TPMLimit),
		RPMLimit:             safeIntDeref(info.RPMLimit),
		BudgetDuration:       safeStringDeref(info.BudgetDuration),
		AllowedCacheControls: info.AllowedCacheControls,
		AllowedRoutes:        info.AllowedRoutes,
		KeyType:              "", // Not available in info response

		Spend:      info.Spend,
		MaxBudget:  safeFloat64Deref(info.MaxBudget),
		SoftBudget: 0, // Not directly available, could derive from soft_budget_cooldown

		Aliases:          info.Aliases,
		Config:           info.Config,
		Permissions:      info.Permissions,
		ObjectPermission: info.ObjectPermission,
		ModelMaxBudget:   info.ModelMaxBudget,
		ModelRPMLimit:    nil, // Not available in info response
		ModelTPMLimit:    nil, // Not available in info response
		EnforcedParams:   nil, // Not available in info response

		Guardrails:      nil, // Not available in info response
		Prompts:         nil, // Not available in info response
		Blocked:         safeBoolDeref(info.Blocked),
		Tags:            nil, // Not available in info response
		SendInviteEmail: false,

		Expires:   info.Expires,
		CreatedBy: info.CreatedBy,
		UpdatedBy: info.UpdatedBy,
		CreatedAt: &info.CreatedAt,
		UpdatedAt: &info.UpdatedAt,

		LitellmBudgetTable: info.LitellmBudgetTable,
	}

	return key
}

// convertKeyToRequest converts the internal Key struct to a KeyGenerateRequest
func convertKeyToRequest(key *Key) *KeyGenerateRequest {
	if key == nil {
		return nil
	}

	request := &KeyGenerateRequest{
		Duration: safeStringPtr(key.Duration),
		KeyAlias: safeStringPtr(key.KeyAlias),
		Key:      safeStringPtr(key.Key),
		TeamID:   safeStringPtr(key.TeamID),
		UserID:   safeStringPtr(key.UserID),
		BudgetID: safeStringPtr(key.BudgetID),
		KeyType:  safeStringPtr(key.KeyType),

		Models:               key.Models,
		Aliases:              key.Aliases,
		Permissions:          key.Permissions,
		AllowedCacheControls: key.AllowedCacheControls,
		Guardrails:           key.Guardrails,
		Prompts:              key.Prompts,
		Tags:                 key.Tags,

		Spend:               safeFloat64Ptr(key.Spend),
		MaxBudget:           safeFloat64Ptr(key.MaxBudget),
		SoftBudget:          safeFloat64Ptr(key.SoftBudget),
		BudgetDuration:      safeStringPtr(key.BudgetDuration),
		MaxParallelRequests: safeIntPtr(key.MaxParallelRequests),
		RPMLimit:            safeIntPtr(key.RPMLimit),
		TPMLimit:            safeIntPtr(key.TPMLimit),
		ModelMaxBudget:      key.ModelMaxBudget,
		ModelRPMLimit:       key.ModelRPMLimit,
		ModelTPMLimit:       key.ModelTPMLimit,

		Metadata:        key.Metadata,
		SendInviteEmail: key.SendInviteEmail,
		Blocked:         key.Blocked,
		EnforcedParams:  key.EnforcedParams,
	}

	return request
}

// Helper functions for safe pointer operations
func safeStringDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func safeStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func safeIntDeref(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func safeIntPtr(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func safeFloat64Deref(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func safeFloat64Ptr(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

func safeBoolDeref(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func safeBoolPtr(b bool) *bool {
	if !b {
		return nil
	}
	return &b
}

func safeStringSliceDeref(s *[]string) []string {
	if s == nil {
		return nil
	}
	return *s
}

func safeMapDeref(m *map[string]interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}
	return *m
}
