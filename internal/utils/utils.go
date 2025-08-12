package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// Type aliases for convenience
type ErrorResponse = litellm.ErrorResponse
type Client = litellm.Client

// expandStringList converts []interface{} to []string
func expandStringList(list []interface{}) []string {
	result := make([]string, len(list))
	for i, v := range list {
		if str, ok := v.(string); ok {
			result[i] = str
		}
	}
	return result
}

func isModelNotFoundError(errResp ErrorResponse) bool {
	if msg, ok := errResp.Error.Message.(string); ok {
		if strings.Contains(msg, "model not found") {
			return true
		}
	}

	if msgMap, ok := errResp.Error.Message.(map[string]interface{}); ok {
		if errStr, ok := msgMap["error"].(string); ok {
			if strings.Contains(errStr, "Model with id=") && strings.Contains(errStr, "not found in db") {
				return true
			}
		}
	}

	// Check Detail.Error field for LiteLLM proxy error format
	if errResp.Detail.Error != "" {
		if strings.Contains(errResp.Detail.Error, "not found on litellm proxy") {
			return true
		}
	}

	return false
}

// MakeRequest is a helper function to make HTTP requests
// Note: This function is deprecated and should use the client's SendRequest method instead
func MakeRequest(client *Client, method, endpoint string, body interface{}) (*http.Response, error) {
	// This is a placeholder - actual implementations should use client.SendRequest
	return nil, fmt.Errorf("MakeRequest is deprecated, use client.SendRequest instead")
}

// GetValueDefault extracts a value from ResourceData with type assertion using the modern GetOk method
// For boolean types, it always includes the value even if false
func GetValueDefault[T any](d *schema.ResourceData, key string, keyData map[string]interface{}) {
	var zero T
	// Check if T is bool type
	if _, isBool := any(zero).(bool); isBool {
		// For booleans, always get the value even if false
		keyData[key] = d.Get(key).(T)
	} else {
		// For other types, use GetOk to only include non-zero values
		if v, ok := d.GetOk(key); ok {
			keyData[key] = v.(T)
		}
	}
}

// GetStringListValue extracts a string list value from ResourceData using the modern GetOk method
func GetStringListValue(d *schema.ResourceData, key string, keyData map[string]interface{}) {
	if v, ok := d.GetOk(key); ok {
		keyData[key] = expandStringList(v.([]interface{}))
	}
}

// Helper functions to handle potential nil values from the API response with generics
// For boolean types, it returns the apiValue directly (no zero-value fallback)
func GetValueWithDefault[T comparable](apiValue, defaultValue T) T {
	var zero T
	// Check if T is bool type
	if _, isBool := any(zero).(bool); isBool {
		// For booleans, return the actual value (including false)
		return apiValue
	}
	// For other types, use zero-value check
	if apiValue != zero {
		return apiValue
	}
	return defaultValue
}

// SetIfNotZero sets a value in ResourceData only if the API value is not zero,
// otherwise keeps the existing value from ResourceData
// For boolean types, it always sets the value (including false)
func SetIfNotZero[T comparable](d *schema.ResourceData, key string, apiValue T) {
	var zero T
	// Check if T is bool type
	if _, isBool := any(zero).(bool); isBool {
		// For booleans, always set the value (including false)
		d.Set(key, apiValue)
	} else {
		// For other types, only set if not zero
		if apiValue != zero {
			d.Set(key, apiValue)
		}
	}
	// If apiValue is zero for non-bool types, we don't set anything, keeping the existing value
}

// handleMCPAPIResponse handles API responses specifically for MCP server operations
func handleMCPAPIResponse(resp *http.Response, result interface{}) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(bodyBytes, &errResp); err == nil {
			if isMCPServerNotFoundError(errResp) {
				return fmt.Errorf("mcp_server_not_found")
			}
		}
		return fmt.Errorf("API request failed: Status: %s, Response: %s",
			resp.Status, string(bodyBytes))
	}

	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	return nil
}

// isMCPServerNotFoundError checks if the error response indicates an MCP server not found
func isMCPServerNotFoundError(errResp ErrorResponse) bool {
	if msg, ok := errResp.Error.Message.(string); ok {
		if strings.Contains(msg, "mcp server not found") || strings.Contains(msg, "server not found") {
			return true
		}
	}

	if msgMap, ok := errResp.Error.Message.(map[string]interface{}); ok {
		if errStr, ok := msgMap["error"].(string); ok {
			if strings.Contains(errStr, "MCP server with id=") && strings.Contains(errStr, "not found") {
				return true
			}
		}
	}

	// Check Detail.Error field for LiteLLM proxy error format
	if errResp.Detail.Error != "" {
		if strings.Contains(errResp.Detail.Error, "not found") {
			return true
		}
	}

	return false
}

// isCredentialNotFoundError checks if the error response indicates a credential not found
func isCredentialNotFoundError(errResp ErrorResponse) bool {
	if msg, ok := errResp.Error.Message.(string); ok {
		if strings.Contains(msg, "credential not found") {
			return true
		}
	}

	if msgMap, ok := errResp.Error.Message.(map[string]interface{}); ok {
		if errStr, ok := msgMap["error"].(string); ok {
			if strings.Contains(errStr, "Credential with name=") && strings.Contains(errStr, "not found") {
				return true
			}
		}
	}

	// Check Detail.Error field for LiteLLM proxy error format
	if errResp.Detail.Error != "" {
		if strings.Contains(errResp.Detail.Error, "credential not found") {
			return true
		}
	}

	return false
}

// handleCredentialAPIResponse handles API responses specifically for credential operations
func handleCredentialAPIResponse(resp *http.Response, result interface{}) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("credential_not_found")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errResp ErrorResponse
		if err := json.Unmarshal(bodyBytes, &errResp); err == nil {
			if isCredentialNotFoundError(errResp) {
				return fmt.Errorf("credential_not_found")
			}
		}
		return fmt.Errorf("API request failed: Status: %s, Response: %s",
			resp.Status, string(bodyBytes))
	}

	// For credential operations, we might get a simple string response or a credential object
	if result != nil {
		if err := json.Unmarshal(bodyBytes, result); err != nil {
			// If parsing fails, it might be a simple string response which is fine for create/update/delete
			return nil
		}
	}

	return nil
}

// isVectorStoreNotFoundError checks if the error response indicates a vector store not found
func isVectorStoreNotFoundError(errResp ErrorResponse) bool {
	if msg, ok := errResp.Error.Message.(string); ok {
		if strings.Contains(msg, "vector store not found") {
			return true
		}
	}

	if msgMap, ok := errResp.Error.Message.(map[string]interface{}); ok {
		if errStr, ok := msgMap["error"].(string); ok {
			if strings.Contains(errStr, "Vector store with id=") && strings.Contains(errStr, "not found") {
				return true
			}
		}
	}

	// Check Detail.Error field for LiteLLM proxy error format
	if errResp.Detail.Error != "" {
		if strings.Contains(errResp.Detail.Error, "vector store not found") {
			return true
		}
	}

	return false
}

// handleVectorStoreAPIResponse handles API responses specifically for vector store operations
func handleVectorStoreAPIResponse(resp *http.Response, result interface{}) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("vector_store_not_found")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errResp ErrorResponse
		if err := json.Unmarshal(bodyBytes, &errResp); err == nil {
			if isVectorStoreNotFoundError(errResp) {
				return fmt.Errorf("vector_store_not_found")
			}
		}
		return fmt.Errorf("API request failed: Status: %s, Response: %s",
			resp.Status, string(bodyBytes))
	}

	if result != nil {
		if err := json.Unmarshal(bodyBytes, result); err != nil {
			return fmt.Errorf("failed to parse response: %v", err)
		}
	}

	return nil
}
