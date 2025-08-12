package vector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

func createVectorStore(ctx context.Context, c *litellm.Client, vectorStore *VectorStore) (*VectorStore, error) {
	request := &VectorStoreRequest{
		CustomLLMProvider:      vectorStore.CustomLLMProvider,
		VectorStoreName:        vectorStore.VectorStoreName,
		VectorStoreDescription: vectorStore.VectorStoreDescription,
		VectorStoreMetadata:    vectorStore.VectorStoreMetadata,
		LiteLLMCredentialName:  vectorStore.LiteLLMCredentialName,
		LiteLLMParams:          vectorStore.LiteLLMParams,
	}

	resp, err := c.SendRequest(ctx, http.MethodPost, "/vector_store/new", request)
	if err != nil {
		return nil, err
	}

	return parseVectorStoreResponse(resp)
}

func getVectorStore(ctx context.Context, c *litellm.Client, vectorStoreID string) (*VectorStore, error) {
	request := &VectorStoreInfoRequest{
		VectorStoreID: vectorStoreID,
	}

	resp, err := c.SendRequest(ctx, http.MethodPost, "/vector_store/info", request)
	if err != nil {
		return nil, err
	}

	return parseVectorStoreResponse(resp)
}

func updateVectorStore(ctx context.Context, c *litellm.Client, vectorStore *VectorStore) (*VectorStore, error) {
	request := &VectorStoreRequest{
		VectorStoreID:          vectorStore.VectorStoreID,
		CustomLLMProvider:      vectorStore.CustomLLMProvider,
		VectorStoreName:        vectorStore.VectorStoreName,
		VectorStoreDescription: vectorStore.VectorStoreDescription,
		VectorStoreMetadata:    vectorStore.VectorStoreMetadata,
		LiteLLMCredentialName:  vectorStore.LiteLLMCredentialName,
		LiteLLMParams:          vectorStore.LiteLLMParams,
	}

	resp, err := c.SendRequest(ctx, http.MethodPost, "/vector_store/update", request)
	if err != nil {
		return nil, err
	}

	return parseVectorStoreResponse(resp)
}

func deleteVectorStore(ctx context.Context, c *litellm.Client, vectorStoreID string) error {
	request := &VectorStoreDeleteRequest{
		VectorStoreID: vectorStoreID,
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/vector_store/delete", request)
	return err
}

func parseVectorStoreResponse(resp map[string]interface{}) (*VectorStore, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	vectorStore := &VectorStore{}

	for k, v := range resp {
		if v == nil {
			continue
		}

		switch k {
		case "vector_store_id":
			if s, ok := v.(string); ok {
				vectorStore.VectorStoreID = s
			}
		case "custom_llm_provider":
			if s, ok := v.(string); ok {
				vectorStore.CustomLLMProvider = s
			}
		case "vector_store_name":
			if s, ok := v.(string); ok {
				vectorStore.VectorStoreName = s
			}
		case "vector_store_description":
			if s, ok := v.(string); ok {
				vectorStore.VectorStoreDescription = s
			}
		case "vector_store_metadata":
			if m, ok := v.(map[string]interface{}); ok {
				vectorStore.VectorStoreMetadata = m
			}
		case "created_at":
			if s, ok := v.(string); ok {
				vectorStore.CreatedAt = s
			}
		case "updated_at":
			if s, ok := v.(string); ok {
				vectorStore.UpdatedAt = s
			}
		case "litellm_credential_name":
			if s, ok := v.(string); ok {
				vectorStore.LiteLLMCredentialName = s
			}
		case "litellm_params":
			if m, ok := v.(map[string]interface{}); ok {
				vectorStore.LiteLLMParams = m
			}
		}
	}

	return vectorStore, nil
}
