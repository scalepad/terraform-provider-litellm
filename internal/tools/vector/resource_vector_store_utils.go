package vector

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func buildVectorStoreData(d *schema.ResourceData) map[string]interface{} {
	vectorStoreData := make(map[string]interface{})

	// String fields
	utils.GetValueDefault[string](d, "vector_store_name", vectorStoreData)
	utils.GetValueDefault[string](d, "custom_llm_provider", vectorStoreData)
	utils.GetValueDefault[string](d, "vector_store_description", vectorStoreData)
	utils.GetValueDefault[string](d, "litellm_credential_name", vectorStoreData)

	// Map fields
	utils.GetValueDefault[map[string]interface{}](d, "vector_store_metadata", vectorStoreData)
	utils.GetValueDefault[map[string]interface{}](d, "litellm_params", vectorStoreData)

	return vectorStoreData
}

func setVectorStoreResourceData(d *schema.ResourceData, vectorStore *VectorStore) error {
	fields := map[string]interface{}{
		"vector_store_id":          vectorStore.VectorStoreID,
		"vector_store_name":        vectorStore.VectorStoreName,
		"custom_llm_provider":      vectorStore.CustomLLMProvider,
		"vector_store_description": vectorStore.VectorStoreDescription,
		"vector_store_metadata":    vectorStore.VectorStoreMetadata,
		"litellm_credential_name":  vectorStore.LiteLLMCredentialName,
		"litellm_params":           vectorStore.LiteLLMParams,
		"created_at":               vectorStore.CreatedAt,
		"updated_at":               vectorStore.UpdatedAt,
	}

	for field, value := range fields {
		if err := d.Set(field, value); err != nil {
			log.Printf("[WARN] Error setting %s: %s", field, err)
			return fmt.Errorf("error setting %s: %s", field, err)
		}
	}

	return nil
}

func mapToVectorStore(data map[string]interface{}) *VectorStore {
	vectorStore := &VectorStore{}
	for k, v := range data {
		switch k {
		case "vector_store_id":
			if s, ok := v.(string); ok {
				vectorStore.VectorStoreID = s
			}
		case "vector_store_name":
			if s, ok := v.(string); ok {
				vectorStore.VectorStoreName = s
			}
		case "custom_llm_provider":
			if s, ok := v.(string); ok {
				vectorStore.CustomLLMProvider = s
			}
		case "vector_store_description":
			if s, ok := v.(string); ok {
				vectorStore.VectorStoreDescription = s
			}
		case "vector_store_metadata":
			if m, ok := v.(map[string]interface{}); ok {
				vectorStore.VectorStoreMetadata = m
			}
		case "litellm_credential_name":
			if s, ok := v.(string); ok {
				vectorStore.LiteLLMCredentialName = s
			}
		case "litellm_params":
			if m, ok := v.(map[string]interface{}); ok {
				vectorStore.LiteLLMParams = m
			}
		case "created_at":
			if s, ok := v.(string); ok {
				vectorStore.CreatedAt = s
			}
		case "updated_at":
			if s, ok := v.(string); ok {
				vectorStore.UpdatedAt = s
			}
		}
	}
	return vectorStore
}

func buildVectorStoreForCreation(data map[string]interface{}) *VectorStore {
	return mapToVectorStore(data)
}
