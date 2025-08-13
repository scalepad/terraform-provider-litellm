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

// listKeys queries the /key/list endpoint to find keys by alias
func listKeys(ctx context.Context, c *litellm.Client, keyAlias string) (*KeyListResponse, error) {
	// Build query parameters
	queryParams := fmt.Sprintf("?page=1&size=10&key_alias=%s&return_full_object=true&include_team_keys=false&sort_order=desc", keyAlias)

	response, err := litellm.SendRequestTyped[interface{}, KeyListResponse](
		ctx, c, http.MethodGet, "/key/list"+queryParams, nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list keys: %w", err)
	}

	return response, nil
}

// findKeyByAlias finds a key by its alias using the list endpoint
func findKeyByAlias(ctx context.Context, c *litellm.Client, keyAlias string) (*KeyListItem, error) {
	response, err := listKeys(ctx, c, keyAlias)
	if err != nil {
		return nil, err
	}

	if len(response.Keys) == 0 {
		return nil, fmt.Errorf("no key found with alias: %s", keyAlias)
	}

	// Return the first matching key
	return &response.Keys[0], nil
}
