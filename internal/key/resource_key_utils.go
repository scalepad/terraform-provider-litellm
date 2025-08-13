package key

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// buildKeyGenerateRequest creates a KeyGenerateRequest directly from ResourceData
func buildKeyGenerateRequest(d *schema.ResourceData) *KeyGenerateRequest {
	request := &KeyGenerateRequest{}

	// String fields
	if v, ok := d.GetOk("duration"); ok {
		request.Duration = stringPtr(v.(string))
	}
	if v, ok := d.GetOk("key_alias"); ok {
		request.KeyAlias = stringPtr(v.(string))
	}
	if v, ok := d.GetOk("user_id"); ok {
		request.UserID = stringPtr(v.(string))
	}
	if v, ok := d.GetOk("team_id"); ok {
		request.TeamID = stringPtr(v.(string))
	}
	if v, ok := d.GetOk("budget_id"); ok {
		request.BudgetID = stringPtr(v.(string))
	}
	if v, ok := d.GetOk("key_type"); ok {
		request.KeyType = stringPtr(v.(string))
	}
	if v, ok := d.GetOk("budget_duration"); ok {
		request.BudgetDuration = stringPtr(v.(string))
	}

	// String slice fields
	if v, ok := d.GetOk("models"); ok {
		request.Models = interfaceSliceToStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("allowed_cache_controls"); ok {
		request.AllowedCacheControls = interfaceSliceToStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("guardrails"); ok {
		request.Guardrails = interfaceSliceToStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("prompts"); ok {
		request.Prompts = interfaceSliceToStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("tags"); ok {
		request.Tags = interfaceSliceToStringSlice(v.([]interface{}))
	}

	// Float64 fields
	if v, ok := d.GetOk("max_budget"); ok {
		request.MaxBudget = floatPtr(v.(float64))
	}
	if v, ok := d.GetOk("soft_budget"); ok {
		request.SoftBudget = floatPtr(v.(float64))
	}
	if v, ok := d.GetOk("spend"); ok {
		request.Spend = floatPtr(v.(float64))
	}

	// Int fields
	if v, ok := d.GetOk("max_parallel_requests"); ok {
		request.MaxParallelRequests = intPtr(v.(int))
	}
	if v, ok := d.GetOk("tpm_limit"); ok {
		request.TPMLimit = intPtr(v.(int))
	}
	if v, ok := d.GetOk("rpm_limit"); ok {
		request.RPMLimit = intPtr(v.(int))
	}

	// Bool fields - use Get() to handle false values properly
	request.Blocked = d.Get("blocked").(bool)
	request.SendInviteEmail = d.Get("send_invite_email").(bool)

	// Map fields
	if v, ok := d.GetOk("metadata"); ok {
		request.Metadata = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("aliases"); ok {
		request.Aliases = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("permissions"); ok {
		request.Permissions = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_max_budget"); ok {
		request.ModelMaxBudget = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_rpm_limit"); ok {
		request.ModelRPMLimit = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_tpm_limit"); ok {
		request.ModelTPMLimit = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("enforced_params"); ok {
		request.EnforcedParams = v.(map[string]interface{})
	}

	return request
}

// buildKeyUpdateRequest creates a KeyGenerateRequest with only changed fields
// This is used for updates to avoid sending unchanged fields to the API
func buildKeyUpdateRequest(d *schema.ResourceData) *KeyGenerateRequest {
	request := &KeyGenerateRequest{}

	// String fields - only include if changed
	if d.HasChange("duration") {
		if v, ok := d.GetOk("duration"); ok {
			request.Duration = stringPtr(v.(string))
		}
	}
	if d.HasChange("key_alias") {
		if v, ok := d.GetOk("key_alias"); ok {
			request.KeyAlias = stringPtr(v.(string))
		}
	}
	if d.HasChange("user_id") {
		if v, ok := d.GetOk("user_id"); ok {
			request.UserID = stringPtr(v.(string))
		}
	}
	if d.HasChange("team_id") {
		if v, ok := d.GetOk("team_id"); ok {
			request.TeamID = stringPtr(v.(string))
		}
	}
	if d.HasChange("budget_id") {
		if v, ok := d.GetOk("budget_id"); ok {
			request.BudgetID = stringPtr(v.(string))
		}
	}
	if d.HasChange("key_type") {
		if v, ok := d.GetOk("key_type"); ok {
			request.KeyType = stringPtr(v.(string))
		}
	}
	if d.HasChange("budget_duration") {
		if v, ok := d.GetOk("budget_duration"); ok {
			request.BudgetDuration = stringPtr(v.(string))
		}
	}

	// String slice fields - only include if changed
	if d.HasChange("models") {
		if v, ok := d.GetOk("models"); ok {
			request.Models = interfaceSliceToStringSlice(v.([]interface{}))
		}
	}
	if d.HasChange("allowed_cache_controls") {
		if v, ok := d.GetOk("allowed_cache_controls"); ok {
			request.AllowedCacheControls = interfaceSliceToStringSlice(v.([]interface{}))
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
	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			request.Tags = interfaceSliceToStringSlice(v.([]interface{}))
		}
	}

	// Float64 fields - only include if changed
	if d.HasChange("max_budget") {
		if v, ok := d.GetOk("max_budget"); ok {
			request.MaxBudget = floatPtr(v.(float64))
		}
	}
	if d.HasChange("soft_budget") {
		if v, ok := d.GetOk("soft_budget"); ok {
			request.SoftBudget = floatPtr(v.(float64))
		}
	}
	if d.HasChange("spend") {
		if v, ok := d.GetOk("spend"); ok {
			request.Spend = floatPtr(v.(float64))
		}
	}

	// Int fields - only include if changed
	if d.HasChange("max_parallel_requests") {
		if v, ok := d.GetOk("max_parallel_requests"); ok {
			request.MaxParallelRequests = intPtr(v.(int))
		}
	}
	if d.HasChange("tpm_limit") {
		if v, ok := d.GetOk("tpm_limit"); ok {
			request.TPMLimit = intPtr(v.(int))
		}
	}
	if d.HasChange("rpm_limit") {
		if v, ok := d.GetOk("rpm_limit"); ok {
			request.RPMLimit = intPtr(v.(int))
		}
	}

	// Bool fields - only include if changed
	if d.HasChange("blocked") {
		request.Blocked = d.Get("blocked").(bool)
	}
	if d.HasChange("send_invite_email") {
		request.SendInviteEmail = d.Get("send_invite_email").(bool)
	}

	// Map fields - only include if changed
	if d.HasChange("metadata") {
		if v, ok := d.GetOk("metadata"); ok {
			request.Metadata = v.(map[string]interface{})
		}
	}
	if d.HasChange("aliases") {
		if v, ok := d.GetOk("aliases"); ok {
			request.Aliases = v.(map[string]interface{})
		}
	}
	if d.HasChange("permissions") {
		if v, ok := d.GetOk("permissions"); ok {
			request.Permissions = v.(map[string]interface{})
		}
	}
	if d.HasChange("model_max_budget") {
		if v, ok := d.GetOk("model_max_budget"); ok {
			request.ModelMaxBudget = v.(map[string]interface{})
		}
	}
	if d.HasChange("model_rpm_limit") {
		if v, ok := d.GetOk("model_rpm_limit"); ok {
			request.ModelRPMLimit = v.(map[string]interface{})
		}
	}
	if d.HasChange("model_tpm_limit") {
		if v, ok := d.GetOk("model_tpm_limit"); ok {
			request.ModelTPMLimit = v.(map[string]interface{})
		}
	}
	if d.HasChange("enforced_params") {
		if v, ok := d.GetOk("enforced_params"); ok {
			request.EnforcedParams = v.(map[string]interface{})
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

// Helper functions for creating pointers
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPtr(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func floatPtr(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

func boolPtr(b bool) *bool {
	return &b
}

// setKeyResourceDataFromGenerate sets resource data from a KeyGenerateResponse
// This is used during creation when we have the full response including the sensitive key
// Logic: Only set fields that aren't empty from the API Response, similar to setKeyResourceDataFromInfo
func setKeyResourceDataFromGenerate(d *schema.ResourceData, response *KeyGenerateResponse) error {
	// Map of all possible fields from API response
	// Only set fields that have values from the API response
	apiFields := map[string]interface{}{
		// Sensitive fields - only available during creation
		"key":       response.Key,
		"token":     response.Token,
		"token_id":  response.TokenID,
		"key_name":  response.KeyName,
		"budget_id": response.BudgetID,

		// Configuration fields from response
		"models":                 response.Models,
		"spend":                  response.Spend,
		"max_budget":             response.MaxBudget,
		"user_id":                response.UserID,
		"team_id":                response.TeamID,
		"max_parallel_requests":  response.MaxParallelRequests,
		"metadata":               response.Metadata,
		"tpm_limit":              response.TPMLimit,
		"rpm_limit":              response.RPMLimit,
		"budget_duration":        response.BudgetDuration,
		"allowed_cache_controls": response.AllowedCacheControls,
		"soft_budget":            response.SoftBudget,
		"key_alias":              response.KeyAlias,
		"duration":               response.Duration,
		"aliases":                response.Aliases,
		"permissions":            response.Permissions,
		"model_max_budget":       response.ModelMaxBudget,
		"model_rpm_limit":        response.ModelRPMLimit,
		"model_tpm_limit":        response.ModelTPMLimit,
		"guardrails":             response.Guardrails,
		"prompts":                response.Prompts,
		"blocked":                response.Blocked,
		"tags":                   response.Tags,
		"send_invite_email":      response.SendInviteEmail,
		"key_type":               response.KeyType,
		"expires":                formatTimePtr(response.Expires),
		"created_by":             response.CreatedBy,
		"updated_by":             response.UpdatedBy,
		"created_at":             formatTime(response.CreatedAt),
		"updated_at":             formatTime(response.UpdatedAt),
		"enforced_params":        response.EnforcedParams,
	}

	// Set fields from API if they have values
	for field, apiValue := range apiFields {
		// If API has a value, use it; otherwise don't set it
		if shouldUseAPIValue(apiValue) {
			if err := d.Set(field, apiValue); err != nil {
				log.Printf("[WARN] Error setting %s: %s", field, err)
				return fmt.Errorf("error setting %s: %s", field, err)
			}
		}
		// If API doesn't have a value, we don't set it (preserves defaults or existing state)
	}

	return nil
}

// setKeyResourceDataFromInfo sets resource data from a KeyInfoResponse
// Logic: If API has the field, use it; otherwise preserve what's in the state
func setKeyResourceDataFromInfo(d *schema.ResourceData, response *KeyInfoResponse) error {
	info := response.Info

	// Map of all possible fields from API response that exist in the schema
	// If the API field is not nil/empty, we use it; otherwise we preserve state
	apiFields := map[string]interface{}{
		"key_name":               info.KeyName,
		"key_alias":              info.KeyAlias,
		"spend":                  info.Spend,
		"models":                 info.Models,
		"aliases":                info.Aliases,
		"user_id":                info.UserID,
		"team_id":                info.TeamID,
		"permissions":            info.Permissions,
		"max_parallel_requests":  info.MaxParallelRequests,
		"metadata":               info.Metadata,
		"blocked":                info.Blocked,
		"tpm_limit":              info.TPMLimit,
		"rpm_limit":              info.RPMLimit,
		"max_budget":             info.MaxBudget,
		"budget_duration":        info.BudgetDuration,
		"allowed_cache_controls": info.AllowedCacheControls,
		"model_max_budget":       info.ModelMaxBudget,
		"budget_id":              info.BudgetID,
		"expires":                formatTimePtr(info.Expires),
		"created_by":             info.CreatedBy,
		"updated_by":             info.UpdatedBy,
		"created_at":             formatTime(info.CreatedAt),
		"updated_at":             formatTime(info.UpdatedAt),
	}

	// Fields that should never be overridden from API (preserve state)
	onlyCreationFields := []string{
		"key",      // Sensitive - only available during creation
		"token",    // Sensitive - only available during creation
		"token_id", // Sensitive - only available during creation
		"key_type", // Only available during creation
	}

	// Set fields from API if they have values, otherwise preserve state
	for field, apiValue := range apiFields {
		// Skip sensitive fields - always preserve from state
		if contains(onlyCreationFields, field) {
			continue
		}

		// If API has a value, use it; otherwise preserve state
		if shouldUseAPIValue(apiValue) {
			if err := d.Set(field, apiValue); err != nil {
				log.Printf("[WARN] Error setting %s: %s", field, err)
				return fmt.Errorf("error setting %s: %s", field, err)
			}
		}
		// If API doesn't have a value, we implicitly preserve the state by not calling d.Set()
	}

	return nil
}

// shouldUseAPIValue determines if we should use the API value or preserve state
func shouldUseAPIValue(apiValue interface{}) bool {
	if apiValue == nil {
		return false
	}

	switch v := apiValue.(type) {
	case string:
		return v != ""
	case *string:
		return v != nil && *v != ""
	case []string:
		return len(v) > 0
	case *[]string:
		return v != nil && len(*v) > 0
	case map[string]interface{}:
		return len(v) > 0
	case *map[string]interface{}:
		return v != nil && len(*v) > 0
	case int:
		return true // Always use int values from API, including 0
	case *int:
		return v != nil
	case float64:
		return true // Always use float64 values from API, including 0.0
	case *float64:
		return v != nil
	case bool:
		return true // Always use bool values from API
	case *bool:
		return v != nil
	default:
		return apiValue != nil
	}
}

// contains checks if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
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
