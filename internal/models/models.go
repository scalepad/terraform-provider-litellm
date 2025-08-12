package models

import "github.com/scalepad/terraform-provider-litellm/internal/litellm"

// ModelResponse represents a response from the API containing model information.
type ModelResponse struct {
	ModelName     string                 `json:"model_name"`
	LiteLLMParams litellm.LiteLLMParams  `json:"litellm_params"`
	ModelInfo     ModelInfo              `json:"model_info"`
	Additional    map[string]interface{} `json:"additional"`
}

// ModelRequest represents a request to create or update a model.
type ModelRequest struct {
	ModelName     string                 `json:"model_name"`
	LiteLLMParams map[string]interface{} `json:"litellm_params"`
	ModelInfo     ModelInfo              `json:"model_info"`
	Additional    map[string]interface{} `json:"additional"`
}
type ModelInfo struct {
	ID        string `json:"id"`
	DBModel   bool   `json:"db_model"`
	BaseModel string `json:"base_model"`
	Tier      string `json:"tier"`
	Mode      string `json:"mode"`
	TeamID    string `json:"team_id,omitempty"`
}
