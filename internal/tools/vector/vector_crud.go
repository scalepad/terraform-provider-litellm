package vector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

func createVectorStore(ctx context.Context, c *litellm.Client, request *VectorStoreGenerateRequest) (*VectorStoreGenerateResponse, error) {
	response, err := litellm.SendRequestTyped[VectorStoreGenerateRequest, VectorStoreGenerateResponse](
		ctx, c, http.MethodPost, "/vector_store/new", request,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create vector store: %w", err)
	}

	return response, nil
}

func getVectorStore(ctx context.Context, c *litellm.Client, vectorStoreID string) (*VectorStoreInfoResponse, error) {
	request := &VectorStoreInfoRequest{
		VectorStoreID: vectorStoreID,
	}

	response, err := litellm.SendRequestTyped[VectorStoreInfoRequest, VectorStoreInfoResponse](
		ctx, c, http.MethodPost, "/vector_store/info", request,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get vector store: %w", err)
	}

	return response, nil
}

func updateVectorStore(ctx context.Context, c *litellm.Client, vectorStoreID string, request *VectorStoreGenerateRequest) (*VectorStoreUpdateResponse, error) {
	// Add the vector store ID to the request for updates
	updateRequest := *request
	updateRequest.VectorStoreID = &vectorStoreID

	response, err := litellm.SendRequestTyped[VectorStoreGenerateRequest, VectorStoreUpdateResponse](
		ctx, c, http.MethodPost, "/vector_store/update", &updateRequest,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update vector store: %w", err)
	}

	return response, nil
}

func deleteVectorStore(ctx context.Context, c *litellm.Client, vectorStoreID string) error {
	deleteRequest := struct {
		VectorStoreID string `json:"vector_store_id"`
	}{
		VectorStoreID: vectorStoreID,
	}

	_, err := litellm.SendRequestTyped[struct {
		VectorStoreID string `json:"vector_store_id"`
	}, interface{}](
		ctx, c, http.MethodPost, "/vector_store/delete", &deleteRequest,
	)
	if err != nil {
		return fmt.Errorf("failed to delete vector store: %w", err)
	}

	return nil
}
