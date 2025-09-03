package serviceaccount

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// createServiceAccount creates a new service account using the /key/service-account/generate endpoint
func CreateServiceAccount(ctx context.Context, c *litellm.Client, request *ServiceAccountGenerateRequest) (*ServiceAccountGenerateResponse, error) {
	// Check if service_account_id exists in metadata, if not generate one
	if request.Metadata == nil {
		request.Metadata = make(map[string]interface{})
	}

	if _, exists := request.Metadata["service_account_id"]; !exists {
		// Generate UUIDv7 for service_account_id
		serviceAccountUUID, err := uuid.NewV7()
		if err != nil {
			return nil, fmt.Errorf("failed to generate service account ID: %w", err)
		}
		request.Metadata["service_account_id"] = serviceAccountUUID.String()
	}

	response, err := litellm.SendRequestTyped[ServiceAccountGenerateRequest, ServiceAccountGenerateResponse](
		ctx, c, http.MethodPost, "/key/service-account/generate", request,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create service account: %w", err)
	}

	return response, nil
}

// getServiceAccount retrieves information about a service account using the /key/info endpoint
func GetServiceAccount(ctx context.Context, c *litellm.Client, keyID string) (*ServiceAccountInfoResponse, error) {
	response, err := litellm.SendRequestTyped[interface{}, ServiceAccountInfoResponse](
		ctx, c, http.MethodGet, fmt.Sprintf("/key/info?key=%s", keyID), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get service account: %w", err)
	}

	return response, nil
}

// updateServiceAccount updates a service account using the /key/update endpoint
func UpdateServiceAccount(ctx context.Context, c *litellm.Client, keyID string, request *ServiceAccountUpdateRequest) (*ServiceAccountGenerateResponse, error) {
	// Set the token for the update request
	request.Token = keyID

	response, err := litellm.SendRequestTyped[ServiceAccountUpdateRequest, ServiceAccountGenerateResponse](
		ctx, c, http.MethodPost, "/key/update", request,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update service account: %w", err)
	}

	return response, nil
}

// deleteServiceAccount deletes a service account using the /key/delete endpoint
func DeleteServiceAccount(ctx context.Context, c *litellm.Client, keyID string) error {
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
		return fmt.Errorf("failed to delete service account: %w", err)
	}

	return nil
}

// isServiceAccount checks if a key is a service account by examining the metadata
func IsServiceAccount(metadata map[string]interface{}) bool {
	if metadata == nil {
		return false
	}

	_, exists := metadata["service_account_id"]
	return exists
}

// getServiceAccountID extracts the service account ID from metadata
func GetServiceAccountID(metadata map[string]interface{}) (string, bool) {
	if metadata == nil {
		return "", false
	}

	if serviceAccountID, ok := metadata["service_account_id"]; ok {
		if id, ok := serviceAccountID.(string); ok {
			return id, true
		}
	}

	return "", false
}
