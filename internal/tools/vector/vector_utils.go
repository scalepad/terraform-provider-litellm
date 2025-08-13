package vector

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

// buildVectorStoreGenerateRequest creates a VectorStoreGenerateRequest directly from ResourceData
func buildVectorStoreGenerateRequest(d *schema.ResourceData) *VectorStoreGenerateRequest {
	request := &VectorStoreGenerateRequest{}

	// String fields
	if v, ok := d.GetOk("vector_store_id"); ok {
		request.VectorStoreID = utils.StringPtr(v.(string))
	} else {
		// Generate UUIDv7 if no ID is provided
		uuidv7, err := uuid.NewV7()
		if err != nil {
			log.Printf("[WARN] Failed to generate UUIDv7, falling back to UUIDv4: %v", err)
			uuidv7 = uuid.New()
		}
		request.VectorStoreID = utils.StringPtr(uuidv7.String())
	}

	if v, ok := d.GetOk("custom_llm_provider"); ok {
		request.CustomLLMProvider = utils.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("vector_store_name"); ok {
		request.VectorStoreName = utils.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("vector_store_description"); ok {
		request.VectorStoreDescription = utils.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("litellm_credential_name"); ok {
		request.LiteLLMCredentialName = utils.StringPtr(v.(string))
	}

	// Map fields
	if v, ok := d.GetOk("vector_store_metadata"); ok {
		request.VectorStoreMetadata = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("litellm_params"); ok {
		request.LiteLLMParams = v.(map[string]interface{})
	}

	return request
}

// buildVectorStoreUpdateRequest creates a VectorStoreGenerateRequest with only changed fields
// This is used for updates to avoid sending unchanged fields to the API
func buildVectorStoreUpdateRequest(d *schema.ResourceData) *VectorStoreGenerateRequest {
	request := &VectorStoreGenerateRequest{}

	// String fields - only include if changed
	if d.HasChange("custom_llm_provider") {
		if v, ok := d.GetOk("custom_llm_provider"); ok {
			request.CustomLLMProvider = utils.StringPtr(v.(string))
		}
	}
	if d.HasChange("vector_store_name") {
		if v, ok := d.GetOk("vector_store_name"); ok {
			request.VectorStoreName = utils.StringPtr(v.(string))
		}
	}
	if d.HasChange("vector_store_description") {
		if v, ok := d.GetOk("vector_store_description"); ok {
			request.VectorStoreDescription = utils.StringPtr(v.(string))
		}
	}
	if d.HasChange("litellm_credential_name") {
		if v, ok := d.GetOk("litellm_credential_name"); ok {
			request.LiteLLMCredentialName = utils.StringPtr(v.(string))
		}
	}

	// Map fields - only include if changed
	if d.HasChange("vector_store_metadata") {
		if v, ok := d.GetOk("vector_store_metadata"); ok {
			request.VectorStoreMetadata = v.(map[string]interface{})
		}
	}
	if d.HasChange("litellm_params") {
		if v, ok := d.GetOk("litellm_params"); ok {
			request.LiteLLMParams = v.(map[string]interface{})
		}
	}

	return request
}

// setVectorStoreResourceDataFromGenerate sets resource data from a VectorStoreGenerateResponse
// This is used during creation when we have the full response
func setVectorStoreResourceDataFromGenerate(d *schema.ResourceData, response *VectorStoreGenerateResponse) error {
	// Use the VectorStore field from the response
	vectorStore := response.VectorStore

	// Map of all possible fields from API response
	// Only set fields that have values from the API response
	apiFields := map[string]interface{}{
		"vector_store_id":          vectorStore.VectorStoreID,
		"custom_llm_provider":      vectorStore.CustomLLMProvider,
		"vector_store_name":        vectorStore.VectorStoreName,
		"vector_store_description": vectorStore.VectorStoreDescription,
		"vector_store_metadata":    vectorStore.VectorStoreMetadata,
		"litellm_credential_name":  vectorStore.LiteLLMCredentialName,
		"litellm_params":           convertLiteLLMParamsToStrings(vectorStore.LiteLLMParams),
		"created_at":               formatTime(vectorStore.CreatedAt),
		"updated_at":               formatTime(vectorStore.UpdatedAt),
	}

	// Set fields from API if they have values
	for field, apiValue := range apiFields {
		// If API has a value, use it; otherwise don't set it
		if utils.ShouldUseAPIValue(apiValue) {
			if err := d.Set(field, apiValue); err != nil {
				log.Printf("[WARN] Error setting %s: %s", field, err)
				return fmt.Errorf("error setting %s: %s", field, err)
			}
		}
		// If API doesn't have a value, we don't set it (preserves defaults or existing state)
	}

	return nil
}

// setVectorStoreResourceDataFromInfo sets resource data from a VectorStoreInfoResponse
// Logic: If API has the field, use it; otherwise preserve what's in the state
func setVectorStoreResourceDataFromInfo(d *schema.ResourceData, response *VectorStoreInfoResponse) error {
	info := response.VectorStore

	// Map of all possible fields from API response that exist in the schema
	// If the API field is not nil/empty, we use it; otherwise we preserve state
	apiFields := map[string]interface{}{
		"vector_store_id":          info.VectorStoreID,
		"custom_llm_provider":      info.CustomLLMProvider,
		"vector_store_name":        info.VectorStoreName,
		"vector_store_description": info.VectorStoreDescription,
		"vector_store_metadata":    info.VectorStoreMetadata,
		"litellm_credential_name":  info.LiteLLMCredentialName,
		"litellm_params":           convertLiteLLMParamsToStrings(info.LiteLLMParams),
		"created_at":               formatTime(info.CreatedAt),
		"updated_at":               formatTime(info.UpdatedAt),
	}

	// Set fields from API if they have values, otherwise preserve state
	for field, apiValue := range apiFields {
		// If API has a value, use it; otherwise preserve state
		if utils.ShouldUseAPIValue(apiValue) {
			if err := d.Set(field, apiValue); err != nil {
				log.Printf("[WARN] Error setting %s: %s", field, err)
				return fmt.Errorf("error setting %s: %s", field, err)
			}
		}
		// If API doesn't have a value, we implicitly preserve the state by not calling d.Set()
	}

	return nil
}

// Helper functions for time formatting
func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

// convertLiteLLMParamsToStrings converts all values in litellm_params to strings
// This is needed because Terraform schema expects string values in TypeMap
func convertLiteLLMParamsToStrings(params map[string]interface{}) map[string]interface{} {
	if params == nil {
		return nil
	}

	converted := make(map[string]interface{})
	for key, value := range params {
		switch v := value.(type) {
		case bool:
			converted[key] = fmt.Sprintf("%t", v)
		case int:
			converted[key] = fmt.Sprintf("%d", v)
		case int64:
			converted[key] = fmt.Sprintf("%d", v)
		case float64:
			converted[key] = fmt.Sprintf("%g", v)
		case string:
			converted[key] = v
		default:
			// For any other type, convert to string using fmt.Sprintf
			converted[key] = fmt.Sprintf("%v", v)
		}
	}

	return converted
}
