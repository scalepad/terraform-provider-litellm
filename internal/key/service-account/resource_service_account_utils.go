package serviceaccount

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

// buildServiceAccountGenerateRequest creates a ServiceAccountGenerateRequest directly from ResourceData
func buildServiceAccountGenerateRequest(d *schema.ResourceData) *ServiceAccountGenerateRequest {
	request := &ServiceAccountGenerateRequest{}

	// Required field
	if v, ok := d.GetOk("team_id"); ok {
		request.TeamID = v.(string)
	}

	// Optional fields
	if v, ok := d.GetOk("key_alias"); ok {
		request.KeyAlias = v.(string)
	}
	if v, ok := d.GetOk("models"); ok {
		request.Models = interfaceSliceToStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("key_type"); ok {
		request.KeyType = v.(string)
	}

	// Handle metadata - ensure service_account_id is included if provided
	metadata := make(map[string]interface{})
	if v, ok := d.GetOk("metadata"); ok {
		for k, val := range v.(map[string]interface{}) {
			metadata[k] = val
		}
	}

	// If service_account_id is provided directly, add it to metadata
	if v, ok := d.GetOk("service_account_id"); ok && v.(string) != "" {
		metadata["service_account_id"] = v.(string)
	}

	// Only set metadata if it has values
	if len(metadata) > 0 {
		request.Metadata = metadata
	}

	return request
}

// buildServiceAccountUpdateRequest creates a ServiceAccountUpdateRequest with only changed fields
func buildServiceAccountUpdateRequest(d *schema.ResourceData) *ServiceAccountUpdateRequest {
	request := &ServiceAccountUpdateRequest{}

	// Always include the token (required for updates)
	request.Token = d.Id()

	// Include changed fields
	if d.HasChange("key_alias") {
		if v, ok := d.GetOk("key_alias"); ok {
			request.KeyAlias = v.(string)
		}
	}
	if d.HasChange("models") {
		if v, ok := d.GetOk("models"); ok {
			request.Models = interfaceSliceToStringSlice(v.([]interface{}))
		}
	}
	if d.HasChange("max_budget") {
		if v, ok := d.GetOk("max_budget"); ok {
			request.MaxBudget = utils.FloatPtr(v.(float64))
		}
	}
	if d.HasChange("budget_duration") {
		if v, ok := d.GetOk("budget_duration"); ok {
			request.BudgetDuration = v.(string)
		}
	}
	if d.HasChange("tpm_limit") {
		if v, ok := d.GetOk("tpm_limit"); ok {
			request.TPMLimit = utils.IntPtr(v.(int))
		}
	}
	if d.HasChange("rpm_limit") {
		if v, ok := d.GetOk("rpm_limit"); ok {
			request.RPMLimit = utils.IntPtr(v.(int))
		}
	}
	if d.HasChange("max_parallel_requests") {
		if v, ok := d.GetOk("max_parallel_requests"); ok {
			request.MaxParallelRequests = utils.IntPtr(v.(int))
		}
	}
	if d.HasChange("guardrails") {
		if v, ok := d.GetOk("guardrails"); ok {
			request.Guardrails = interfaceSliceToStringSlice(v.([]interface{}))
		}
	}
	if d.HasChange("prompts") {
		if v, ok := d.GetOk("prompts"); ok {
			request.Prompts = interfaceSliceToStringSlice(v.([]interface{}))
		}
	}
	if d.HasChange("team_id") {
		if v, ok := d.GetOk("team_id"); ok {
			request.TeamID = v.(string)
		}
	}
	if d.HasChange("metadata") {
		if v, ok := d.GetOk("metadata"); ok {
			request.Metadata = v.(map[string]interface{})
		}
	}

	return request
}

// interfaceSliceToStringSlice converts []interface{} to []string
func interfaceSliceToStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		if s, ok := v.(string); ok {
			result[i] = s
		}
	}
	return result
}

// setServiceAccountResourceData sets resource data from a ServiceAccountGenerateResponse
func setServiceAccountResourceData(d *schema.ResourceData, response *ServiceAccountGenerateResponse) error {
	// Map of all possible fields from API response
	apiFields := map[string]interface{}{
		// Sensitive fields - only available during creation
		"key":      response.Key,
		"token":    response.Token,
		"token_id": response.TokenID,
		"key_name": response.KeyName,

		// Configuration fields
		"models":                 response.Models,
		"spend":                  response.Spend,
		"max_budget":             response.MaxBudget,
		"team_id":                response.TeamID,
		"max_parallel_requests":  response.MaxParallelRequests,
		"metadata":               response.Metadata,
		"tpm_limit":              response.TPMLimit,
		"rpm_limit":              response.RPMLimit,
		"budget_duration":        response.BudgetDuration,
		"allowed_cache_controls": response.AllowedCacheControls,
		"key_alias":              response.KeyAlias,
		"aliases":                response.Aliases,
		"permissions":            response.Permissions,
		"model_max_budget":       response.ModelMaxBudget,
		"model_rpm_limit":        response.ModelRPMLimit,
		"model_tpm_limit":        response.ModelTPMLimit,
		"guardrails":             response.Guardrails,
		"prompts":                response.Prompts,
		"tags":                   response.Tags,
		"expires":                formatTimePtr(response.Expires),
		"created_by":             response.CreatedBy,
		"updated_by":             response.UpdatedBy,
		"created_at":             formatTime(response.CreatedAt),
		"updated_at":             formatTime(response.UpdatedAt),
		"enforced_params":        response.EnforcedParams,

		// Extract service_account_id from metadata if present
		"service_account_id": extractServiceAccountID(response.Metadata),
	}

	// Set fields from API if they have values
	for field, apiValue := range apiFields {
		if utils.ShouldUseAPIValue(apiValue) {
			if err := d.Set(field, apiValue); err != nil {
				log.Printf("[WARN] Error setting %s: %s", field, err)
				return fmt.Errorf("error setting %s: %s", field, err)
			}
		}
	}

	return nil
}

// setServiceAccountResourceDataFromInfo sets resource data from a ServiceAccountInfoResponse
func setServiceAccountResourceDataFromInfo(d *schema.ResourceData, response *ServiceAccountInfoResponse) error {
	info := response.Info

	// Map of all possible fields from API response
	apiFields := map[string]interface{}{
		"key_name":   info.KeyName,
		"key_alias":  info.KeyAlias,
		"spend":      info.Spend,
		"models":     info.Models,
		"team_id":    info.TeamID,
		"metadata":   info.Metadata,
		"max_budget": info.MaxBudget,
		"created_at": formatTime(info.CreatedAt),
		"updated_at": formatTime(info.UpdatedAt),

		// Extract service_account_id from metadata if present
		"service_account_id": extractServiceAccountID(info.Metadata),
	}

	// Set fields from API if they have values
	for field, apiValue := range apiFields {
		if utils.ShouldUseAPIValue(apiValue) {
			if err := d.Set(field, apiValue); err != nil {
				log.Printf("[WARN] Error setting %s: %s", field, err)
				return fmt.Errorf("error setting %s: %s", field, err)
			}
		}
	}

	return nil
}

// extractServiceAccountID extracts the service_account_id from metadata
func extractServiceAccountID(metadata map[string]interface{}) string {
	if metadata == nil {
		return ""
	}

	if serviceAccountID, ok := metadata["service_account_id"]; ok {
		if id, ok := serviceAccountID.(string); ok {
			return id
		}
	}

	return ""
}

// Helper functions for time formatting
func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

func formatTimePtr(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}
