package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"net/http"
	"strings"
)

func createModel(ctx context.Context, c *litellm.Client, model *Model) (*Model, error) {
	// Generate a UUID for new models
	if model.ModelInfo.ID == "" {
		model.ModelInfo.ID = uuid.New().String()
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/model/new", model)
	if err != nil {
		return nil, err
	}

	// For create operations, return the model with the generated ID
	return model, nil
}

func getModel(ctx context.Context, c *litellm.Client, modelID string) (*ModelResponse, error) {
	endpoint := fmt.Sprintf("/model/info?litellm_model_id=%s", modelID)
	resp, err := c.SendRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		// Check if it's a not found error
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
			return nil, nil
		}
		return nil, err
	}

	return parseModelAPIResponse(resp)
}

func updateModel(ctx context.Context, c *litellm.Client, model *Model) (*Model, error) {
	_, err := c.SendRequest(ctx, http.MethodPost, "/model/update", model)
	if err != nil {
		// If model not found during update, try to create it instead
		if strings.Contains(err.Error(), "not found") {
			return createModel(ctx, c, model)
		}
		return nil, err
	}

	return model, nil
}

func deleteModel(ctx context.Context, c *litellm.Client, modelID string) error {
	deleteReq := map[string]interface{}{
		"id": modelID,
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/model/delete", deleteReq)

	// If it's a not found error, consider it successful (already deleted)
	if err != nil && (strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404")) {
		return nil
	}

	return err
}
