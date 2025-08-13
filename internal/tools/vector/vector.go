package vector

import "time"

// VectorStoreGenerateRequest represents the request payload for creating a new vector store
type VectorStoreGenerateRequest struct {
	VectorStoreID          *string                `json:"vector_store_id,omitempty"`          // Unique identifier for the vector store
	CustomLLMProvider      *string                `json:"custom_llm_provider,omitempty"`      // Provider of the vector store
	VectorStoreName        *string                `json:"vector_store_name,omitempty"`        // Name of the vector store
	VectorStoreDescription *string                `json:"vector_store_description,omitempty"` // Description of the vector store
	VectorStoreMetadata    map[string]interface{} `json:"vector_store_metadata,omitempty"`    // Additional metadata for the vector store
	LiteLLMCredentialName  *string                `json:"litellm_credential_name,omitempty"`  // Name of the LiteLLM credential
	LiteLLMParams          map[string]interface{} `json:"litellm_params,omitempty"`           // Additional LiteLLM parameters
}

// VectorStoreGenerateResponse represents the response from creating a new vector store
type VectorStoreGenerateResponse struct {
	Status      string          `json:"status"`       // Status of the operation
	Message     string          `json:"message"`      // Message from the API
	VectorStore VectorStoreInfo `json:"vector_store"` // The created vector store information
}

// VectorStoreInfoResponse represents the response from the /vector_store/info endpoint
type VectorStoreInfoResponse struct {
	VectorStore VectorStoreInfo `json:"vector_store"` // The vector store information
}

// VectorStoreInfo represents the detailed information about a vector store
type VectorStoreInfo struct {
	VectorStoreID          string                 `json:"vector_store_id"`
	CustomLLMProvider      string                 `json:"custom_llm_provider"`
	VectorStoreName        string                 `json:"vector_store_name"`
	VectorStoreDescription string                 `json:"vector_store_description,omitempty"`
	VectorStoreMetadata    map[string]interface{} `json:"vector_store_metadata,omitempty"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
	LiteLLMCredentialName  string                 `json:"litellm_credential_name,omitempty"`
	LiteLLMParams          map[string]interface{} `json:"litellm_params,omitempty"`
}

// VectorStoreUpdateResponse represents the response from updating a vector store
type VectorStoreUpdateResponse struct {
	VectorStore VectorStoreInfo `json:"vector_store"` // The updated vector store information
}

// VectorStoreListResponse represents a response from the API containing a list of vector stores
type VectorStoreListResponse struct {
	Object      string            `json:"object"`
	Data        []VectorStoreInfo `json:"data"`
	TotalCount  int               `json:"total_count"`
	CurrentPage int               `json:"current_page"`
	TotalPages  int               `json:"total_pages"`
}

// VectorStoreDeleteRequest represents a request to delete a vector store
type VectorStoreDeleteRequest struct {
	VectorStoreID string `json:"vector_store_id"`
}

// VectorStoreInfoRequest represents a request to get vector store information
type VectorStoreInfoRequest struct {
	VectorStoreID string `json:"vector_store_id"`
}
