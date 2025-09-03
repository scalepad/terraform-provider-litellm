package serviceaccount

import "time"

// ServiceAccountGenerateRequest represents the request payload for creating a new service account
type ServiceAccountGenerateRequest struct {
	// Core configuration
	TeamID   string                 `json:"team_id"`             // The team id the service account belongs to
	KeyAlias string                 `json:"key_alias,omitempty"` // User defined key alias
	Models   []string               `json:"models,omitempty"`    // Models the service account can access
	KeyType  string                 `json:"key_type,omitempty"`  // Type of key that determines default allowed routes
	Metadata map[string]interface{} `json:"metadata,omitempty"`  // Metadata for the service account containing service_account_id
}

// ServiceAccountGenerateResponse represents the response from creating a new service account
// This matches the actual API response structure
type ServiceAccountGenerateResponse struct {
	// Core key fields - these are returned by the API
	Key      string  `json:"key"`                 // The generated API key (sensitive)
	KeyAlias *string `json:"key_alias,omitempty"` // User defined key alias
	KeyName  string  `json:"key_name"`            // The truncated key name for display
	Token    string  `json:"token"`               // The token identifier
	TokenID  string  `json:"token_id"`            // The unique token ID

	// Configuration fields
	Models               []string                `json:"models"`                          // List of models this key can access
	Spend                float64                 `json:"spend"`                           // Current spend amount
	MaxBudget            *float64                `json:"max_budget,omitempty"`            // Maximum budget
	UserID               *string                 `json:"user_id,omitempty"`               // User ID (null for service accounts)
	TeamID               string                  `json:"team_id"`                         // The team ID associated with this key
	MaxParallelRequests  *int                    `json:"max_parallel_requests,omitempty"` // Maximum parallel requests
	Metadata             map[string]interface{}  `json:"metadata"`                        // Additional metadata containing service_account_id
	TPMLimit             *int                    `json:"tpm_limit,omitempty"`             // Tokens per minute limit
	RPMLimit             *int                    `json:"rpm_limit,omitempty"`             // Requests per minute limit
	BudgetDuration       *string                 `json:"budget_duration,omitempty"`       // Budget reset duration
	AllowedCacheControls []string                `json:"allowed_cache_controls"`          // Allowed cache control values
	Config               map[string]interface{}  `json:"config"`                          // Additional configuration
	Permissions          map[string]interface{}  `json:"permissions"`                     // Permissions configuration
	ModelMaxBudget       map[string]interface{}  `json:"model_max_budget"`                // Per-model budget limits
	ModelRPMLimit        *map[string]interface{} `json:"model_rpm_limit,omitempty"`       // Per-model RPM limits
	ModelTPMLimit        *map[string]interface{} `json:"model_tpm_limit,omitempty"`       // Per-model TPM limits
	Guardrails           *[]string               `json:"guardrails,omitempty"`            // List of guardrails
	Prompts              *[]string               `json:"prompts,omitempty"`               // List of prompts
	Blocked              *bool                   `json:"blocked,omitempty"`               // Whether key is blocked
	Aliases              map[string]interface{}  `json:"aliases"`                         // Model aliases
	ObjectPermission     interface{}             `json:"object_permission"`               // Object-level permissions
	BudgetID             *string                 `json:"budget_id,omitempty"`             // The budget ID associated with this key
	Tags                 *[]string               `json:"tags,omitempty"`                  // Tags
	EnforcedParams       *map[string]interface{} `json:"enforced_params,omitempty"`       // Enforced parameters
	AllowedRoutes        []string                `json:"allowed_routes"`                  // Allowed API routes
	Expires              *time.Time              `json:"expires,omitempty"`               // Expiration timestamp
	LitellmBudgetTable   interface{}             `json:"litellm_budget_table"`            // Reference to budget table
	CreatedBy            string                  `json:"created_by"`                      // User who created the key
	UpdatedBy            string                  `json:"updated_by"`                      // User who last updated the key
	CreatedAt            time.Time               `json:"created_at"`                      // Creation timestamp
	UpdatedAt            time.Time               `json:"updated_at"`                      // Last update timestamp
}

// ServiceAccountInfoResponse represents the response from the /key/info endpoint for service accounts
type ServiceAccountInfoResponse struct {
	Key  string             `json:"key"`  // The key identifier
	Info ServiceAccountInfo `json:"info"` // The key information
}

// ServiceAccountUpdateRequest represents the request payload for updating a service account
type ServiceAccountUpdateRequest struct {
	// Core configuration
	KeyAlias            string                 `json:"key_alias,omitempty"`             // User defined key alias
	Models              []string               `json:"models,omitempty"`                // Models the service account can access
	MaxBudget           *float64               `json:"max_budget,omitempty"`            // Maximum budget
	BudgetDuration      string                 `json:"budget_duration,omitempty"`       // Budget reset duration
	TPMLimit            *int                   `json:"tpm_limit,omitempty"`             // Tokens per minute limit
	RPMLimit            *int                   `json:"rpm_limit,omitempty"`             // Requests per minute limit
	MaxParallelRequests *int                   `json:"max_parallel_requests,omitempty"` // Maximum parallel requests
	Guardrails          []string               `json:"guardrails,omitempty"`            // List of guardrails
	Prompts             []string               `json:"prompts,omitempty"`               // List of prompts
	TeamID              string                 `json:"team_id"`                         // The team id the service account belongs to
	Metadata            map[string]interface{} `json:"metadata,omitempty"`              // Metadata for the service account containing service_account_id
	Token               string                 `json:"token"`                           // The token identifier for updates
	DisabledCallbacks   []string               `json:"disabled_callbacks,omitempty"`    // Disabled callbacks
	Key                 string                 `json:"key,omitempty"`                   // The key value for updates
	ObjectPermission    map[string]interface{} `json:"object_permission,omitempty"`     // Object-level permissions
}

// ServiceAccountInfo represents the detailed information about a service account
type ServiceAccountInfo struct {
	KeyName   string                 `json:"key_name"`
	KeyAlias  *string                `json:"key_alias"`
	Spend     float64                `json:"spend"`
	Models    []string               `json:"models"`
	TeamID    string                 `json:"team_id"`
	Metadata  map[string]interface{} `json:"metadata"`
	MaxBudget *float64               `json:"max_budget"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
