package vector

// VectorStore represents a LiteLLM vector store.
type VectorStore struct {
	VectorStoreID          string                 `json:"vector_store_id,omitempty"`
	CustomLLMProvider      string                 `json:"custom_llm_provider"`
	VectorStoreName        string                 `json:"vector_store_name"`
	VectorStoreDescription string                 `json:"vector_store_description,omitempty"`
	VectorStoreMetadata    map[string]interface{} `json:"vector_store_metadata,omitempty"`
	CreatedAt              string                 `json:"created_at,omitempty"`
	UpdatedAt              string                 `json:"updated_at,omitempty"`
	LiteLLMCredentialName  string                 `json:"litellm_credential_name,omitempty"`
	LiteLLMParams          map[string]interface{} `json:"litellm_params,omitempty"`
}

// VectorStoreRequest represents a request to create or update a vector store.
type VectorStoreRequest struct {
	VectorStoreID          string                 `json:"vector_store_id,omitempty"`
	CustomLLMProvider      string                 `json:"custom_llm_provider"`
	VectorStoreName        string                 `json:"vector_store_name"`
	VectorStoreDescription string                 `json:"vector_store_description,omitempty"`
	VectorStoreMetadata    map[string]interface{} `json:"vector_store_metadata,omitempty"`
	LiteLLMCredentialName  string                 `json:"litellm_credential_name,omitempty"`
	LiteLLMParams          map[string]interface{} `json:"litellm_params,omitempty"`
}

// VectorStoreResponse represents a response from the API containing vector store information.
type VectorStoreResponse struct {
	VectorStoreID          string                 `json:"vector_store_id"`
	CustomLLMProvider      string                 `json:"custom_llm_provider"`
	VectorStoreName        string                 `json:"vector_store_name"`
	VectorStoreDescription string                 `json:"vector_store_description,omitempty"`
	VectorStoreMetadata    map[string]interface{} `json:"vector_store_metadata,omitempty"`
	CreatedAt              string                 `json:"created_at,omitempty"`
	UpdatedAt              string                 `json:"updated_at,omitempty"`
	LiteLLMCredentialName  string                 `json:"litellm_credential_name,omitempty"`
	LiteLLMParams          map[string]interface{} `json:"litellm_params,omitempty"`
}

// VectorStoreListResponse represents a response from the API containing a list of vector stores.
type VectorStoreListResponse struct {
	Object      string                `json:"object"`
	Data        []VectorStoreResponse `json:"data"`
	TotalCount  int                   `json:"total_count"`
	CurrentPage int                   `json:"current_page"`
	TotalPages  int                   `json:"total_pages"`
}

// VectorStoreDeleteRequest represents a request to delete a vector store.
type VectorStoreDeleteRequest struct {
	VectorStoreID string `json:"vector_store_id"`
}

// VectorStoreInfoRequest represents a request to get vector store information.
type VectorStoreInfoRequest struct {
	VectorStoreID string `json:"vector_store_id"`
}
