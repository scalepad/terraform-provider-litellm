package vector

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func TestBuildVectorStoreData(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "all fields populated",
			input: map[string]interface{}{
				"vector_store_name":        "test-vector-store",
				"custom_llm_provider":      "openai",
				"vector_store_description": "Test vector store description",
				"litellm_credential_name":  "test-credential",
				"vector_store_metadata":    map[string]interface{}{"env": "test"},
				"litellm_params":           map[string]interface{}{"model": "text-embedding-ada-002"},
			},
			expected: map[string]interface{}{
				"vector_store_name":        "test-vector-store",
				"custom_llm_provider":      "openai",
				"vector_store_description": "Test vector store description",
				"litellm_credential_name":  "test-credential",
				"vector_store_metadata":    map[string]interface{}{"env": "test"},
				"litellm_params":           map[string]interface{}{"model": "text-embedding-ada-002"},
			},
		},
		{
			name: "minimal fields",
			input: map[string]interface{}{
				"vector_store_name":   "minimal-store",
				"custom_llm_provider": "openai",
			},
			expected: map[string]interface{}{
				"vector_store_name":   "minimal-store",
				"custom_llm_provider": "openai",
			},
		},
		{
			name:     "empty input",
			input:    map[string]interface{}{},
			expected: map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the actual ResourceLiteLLMVectorStore schema
			resource := ResourceLiteLLMVectorStore()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildVectorStoreData(d)

			// Check that all expected keys are present
			for key, expectedValue := range tt.expected {
				actualValue, exists := result[key]
				if !exists {
					t.Errorf("buildVectorStoreData() missing key %s", key)
					continue
				}
				if !utils.CompareValues(actualValue, expectedValue) {
					t.Errorf("buildVectorStoreData() key %s = %v, want %v", key, actualValue, expectedValue)
				}
			}

			// Check that no unexpected keys are present
			for key := range result {
				if _, expected := tt.expected[key]; !expected {
					t.Errorf("buildVectorStoreData() unexpected key %s with value %v", key, result[key])
				}
			}
		})
	}
}

func TestSetVectorStoreResourceData(t *testing.T) {
	tests := []struct {
		name        string
		vectorStore *VectorStore
		expectError bool
	}{
		{
			name: "complete vector store data",
			vectorStore: &VectorStore{
				VectorStoreID:          "vs-123",
				VectorStoreName:        "test-store",
				CustomLLMProvider:      "openai",
				VectorStoreDescription: "Test description",
				VectorStoreMetadata:    map[string]interface{}{"env": "test"},
				LiteLLMCredentialName:  "test-cred",
				LiteLLMParams:          map[string]interface{}{"model": "text-embedding-ada-002"},
				CreatedAt:              "2023-01-01T00:00:00Z",
				UpdatedAt:              "2023-01-02T00:00:00Z",
			},
			expectError: false,
		},
		{
			name: "minimal vector store data",
			vectorStore: &VectorStore{
				VectorStoreID:     "vs-minimal",
				VectorStoreName:   "minimal-store",
				CustomLLMProvider: "openai",
			},
			expectError: false,
		},
		{
			name: "vector store with nil maps",
			vectorStore: &VectorStore{
				VectorStoreID:       "vs-nil",
				VectorStoreName:     "nil-store",
				CustomLLMProvider:   "openai",
				VectorStoreMetadata: nil,
				LiteLLMParams:       nil,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := ResourceLiteLLMVectorStore()
			d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

			err := setVectorStoreResourceData(d, tt.vectorStore)

			if tt.expectError && err == nil {
				t.Errorf("setVectorStoreResourceData() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("setVectorStoreResourceData() unexpected error: %v", err)
			}

			if err == nil {
				// Verify some key fields were set correctly
				if d.Get("vector_store_id") != tt.vectorStore.VectorStoreID {
					t.Errorf("Expected vector_store_id %s, got %s", tt.vectorStore.VectorStoreID, d.Get("vector_store_id"))
				}
				if d.Get("vector_store_name") != tt.vectorStore.VectorStoreName {
					t.Errorf("Expected vector_store_name %s, got %s", tt.vectorStore.VectorStoreName, d.Get("vector_store_name"))
				}
				if d.Get("custom_llm_provider") != tt.vectorStore.CustomLLMProvider {
					t.Errorf("Expected custom_llm_provider %s, got %s", tt.vectorStore.CustomLLMProvider, d.Get("custom_llm_provider"))
				}
			}
		})
	}
}

func TestMapToVectorStore(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *VectorStore
	}{
		{
			name: "complete data map",
			input: map[string]interface{}{
				"vector_store_id":          "vs-123",
				"vector_store_name":        "test-store",
				"custom_llm_provider":      "openai",
				"vector_store_description": "Test description",
				"vector_store_metadata":    map[string]interface{}{"env": "test"},
				"litellm_credential_name":  "test-cred",
				"litellm_params":           map[string]interface{}{"model": "text-embedding-ada-002"},
				"created_at":               "2023-01-01T00:00:00Z",
				"updated_at":               "2023-01-02T00:00:00Z",
			},
			expected: &VectorStore{
				VectorStoreID:          "vs-123",
				VectorStoreName:        "test-store",
				CustomLLMProvider:      "openai",
				VectorStoreDescription: "Test description",
				VectorStoreMetadata:    map[string]interface{}{"env": "test"},
				LiteLLMCredentialName:  "test-cred",
				LiteLLMParams:          map[string]interface{}{"model": "text-embedding-ada-002"},
				CreatedAt:              "2023-01-01T00:00:00Z",
				UpdatedAt:              "2023-01-02T00:00:00Z",
			},
		},
		{
			name:     "empty map",
			input:    map[string]interface{}{},
			expected: &VectorStore{},
		},
		{
			name: "partial data",
			input: map[string]interface{}{
				"vector_store_id":     "vs-partial",
				"vector_store_name":   "partial-store",
				"custom_llm_provider": "openai",
			},
			expected: &VectorStore{
				VectorStoreID:     "vs-partial",
				VectorStoreName:   "partial-store",
				CustomLLMProvider: "openai",
			},
		},
		{
			name: "unknown fields ignored",
			input: map[string]interface{}{
				"vector_store_id":     "vs-unknown",
				"unknown_field":       "should_be_ignored",
				"vector_store_name":   "unknown-store",
				"custom_llm_provider": "openai",
			},
			expected: &VectorStore{
				VectorStoreID:     "vs-unknown",
				VectorStoreName:   "unknown-store",
				CustomLLMProvider: "openai",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapToVectorStore(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mapToVectorStore() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestBuildVectorStoreForCreation(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *VectorStore
	}{
		{
			name: "creation data",
			input: map[string]interface{}{
				"vector_store_name":        "creation-store",
				"custom_llm_provider":      "openai",
				"vector_store_description": "Creation test",
				"vector_store_metadata":    map[string]interface{}{"env": "test"},
			},
			expected: &VectorStore{
				VectorStoreName:        "creation-store",
				CustomLLMProvider:      "openai",
				VectorStoreDescription: "Creation test",
				VectorStoreMetadata:    map[string]interface{}{"env": "test"},
			},
		},
		{
			name:     "empty creation data",
			input:    map[string]interface{}{},
			expected: &VectorStore{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildVectorStoreForCreation(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildVectorStoreForCreation() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}
