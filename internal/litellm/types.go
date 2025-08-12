package litellm

// ProviderConfig holds the configuration for the LiteLLM provider.
type ProviderConfig struct {
	APIBase            string
	APIKey             string
	InsecureSkipVerify bool
}

// ErrorResponse represents an error response from the API.
type ErrorResponse struct {
	Error struct {
		Message interface{} `json:"message"`
	} `json:"error"`
	Detail struct {
		Error string `json:"error"`
	} `json:"detail"`
}

// LiteLLMParams represents the parameters for LiteLLM.
type LiteLLMParams struct {
	CustomLLMProvider              string                 `json:"custom_llm_provider"`
	TPM                            int                    `json:"tpm,omitempty"`
	RPM                            int                    `json:"rpm,omitempty"`
	ReasoningEffort                string                 `json:"reasoning_effort,omitempty"`
	Thinking                       map[string]interface{} `json:"thinking,omitempty"`
	MergeReasoningContentInChoices bool                   `json:"merge_reasoning_content_in_choices,omitempty"`
	APIKey                         string                 `json:"api_key,omitempty"`
	APIBase                        string                 `json:"api_base,omitempty"`
	APIVersion                     string                 `json:"api_version,omitempty"`
	Model                          string                 `json:"model"`
	InputCostPerToken              float64                `json:"input_cost_per_token,omitempty"`
	OutputCostPerToken             float64                `json:"output_cost_per_token,omitempty"`
	InputCostPerPixel              float64                `json:"input_cost_per_pixel,omitempty"`
	OutputCostPerPixel             float64                `json:"output_cost_per_pixel,omitempty"`
	InputCostPerSecond             float64                `json:"input_cost_per_second,omitempty"`
	OutputCostPerSecond            float64                `json:"output_cost_per_second,omitempty"`
	AWSAccessKeyID                 string                 `json:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey             string                 `json:"aws_secret_access_key,omitempty"`
	AWSRegionName                  string                 `json:"aws_region_name,omitempty"`
	VertexProject                  string                 `json:"vertex_project,omitempty"`
	VertexLocation                 string                 `json:"vertex_location,omitempty"`
	VertexCredentials              string                 `json:"vertex_credentials,omitempty"`
}
