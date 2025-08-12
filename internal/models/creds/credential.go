package creds

// Credential represents a credential configuration in LiteLLM
type Credential struct {
	CredentialName   string                 `json:"credential_name,omitempty"`
	ModelID          string                 `json:"model_id,omitempty"`
	CredentialInfo   map[string]interface{} `json:"credential_info,omitempty"`
	CredentialValues map[string]interface{} `json:"credential_values,omitempty"`
}

// CredentialRequest represents a request to create or update a credential.
type CredentialRequest struct {
	CredentialName   string                 `json:"credential_name"`
	CredentialInfo   map[string]interface{} `json:"credential_info,omitempty"`
	CredentialValues map[string]interface{} `json:"credential_values,omitempty"`
	ModelID          string                 `json:"model_id,omitempty"`
}

// CredentialResponse represents a response from the API containing credential information.
type CredentialResponse struct {
	CredentialName   string                 `json:"credential_name"`
	CredentialInfo   map[string]interface{} `json:"credential_info,omitempty"`
	CredentialValues map[string]interface{} `json:"credential_values,omitempty"`
}
