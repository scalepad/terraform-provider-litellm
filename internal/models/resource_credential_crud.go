package models

import (
	"context"
	"fmt"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"net/http"
	"strings"
)

func createCredential(ctx context.Context, c *litellm.Client, credential *Credential) (*Credential, error) {
	_, err := c.SendRequest(ctx, http.MethodPost, "/credentials", credential)
	if err != nil {
		return nil, err
	}

	// For create operations, we might get a simple success response
	// Return the credential that was sent for creation
	return credential, nil
}

func getCredential(ctx context.Context, c *litellm.Client, credentialName string) (*Credential, error) {
	endpoint := fmt.Sprintf("/credentials/by_name/%s", credentialName)
	resp, err := c.SendRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		// Check if it's a not found error
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
			return nil, nil
		}
		return nil, err
	}

	return parseCredentialAPIResponse(resp)
}

func updateCredential(ctx context.Context, c *litellm.Client, credential *Credential) (*Credential, error) {
	endpoint := fmt.Sprintf("/credentials/%s", credential.CredentialName)

	// Create update data with only the fields that can be updated
	updateData := map[string]interface{}{
		"credential_name":   credential.CredentialName,
		"credential_info":   credential.CredentialInfo,
		"credential_values": credential.CredentialValues,
	}

	if credential.ModelID != "" {
		updateData["model_id"] = credential.ModelID
	}

	resp, err := c.SendRequest(ctx, http.MethodPatch, endpoint, updateData)
	if err != nil {
		return nil, err
	}

	// For update operations, we might get a simple success response
	// Try to parse the response, but if it fails, return the updated credential
	if updatedCredential, parseErr := parseCredentialAPIResponse(resp); parseErr == nil {
		return updatedCredential, nil
	}

	return credential, nil
}

func deleteCredential(ctx context.Context, c *litellm.Client, credentialName string) error {
	endpoint := fmt.Sprintf("/credentials/%s", credentialName)
	_, err := c.SendRequest(ctx, http.MethodDelete, endpoint, nil)

	// If it's a not found error, consider it successful (already deleted)
	if err != nil && (strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404")) {
		return nil
	}

	return err
}
