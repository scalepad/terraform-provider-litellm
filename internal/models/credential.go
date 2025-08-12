package models

// Credential represents a credential configuration in LiteLLM
type Credential struct {
	CredentialName   string                 `json:"credential_name,omitempty"`
	ModelID          string                 `json:"model_id,omitempty"`
	CredentialInfo   map[string]interface{} `json:"credential_info,omitempty"`
	CredentialValues map[string]interface{} `json:"credential_values,omitempty"`
}
