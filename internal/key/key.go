package key

import "time"

// KeyGenerateRequest represents the request payload for creating a new key
type KeyGenerateRequest struct {
	// Core configuration
	Duration *string `json:"duration,omitempty"`  // Duration for which this key is valid
	KeyAlias *string `json:"key_alias,omitempty"` // User defined key alias
	Key      *string `json:"key,omitempty"`       // User defined key value
	TeamID   *string `json:"team_id,omitempty"`   // The team id of the key
	UserID   *string `json:"user_id,omitempty"`   // The user id of the key
	BudgetID *string `json:"budget_id,omitempty"` // The budget id associated with the key
	KeyType  *string `json:"key_type,omitempty"`  // Type of key that determines default allowed routes

	// Model and access control
	Models               []string               `json:"models,omitempty"`                 // Models a user is allowed to call
	Aliases              map[string]interface{} `json:"aliases,omitempty"`                // Any alias mappings
	Permissions          map[string]interface{} `json:"permissions,omitempty"`            // Key-specific permissions
	AllowedCacheControls []string               `json:"allowed_cache_controls,omitempty"` // List of allowed cache control values
	Guardrails           []string               `json:"guardrails,omitempty"`             // List of active guardrails for the key
	Prompts              []string               `json:"prompts,omitempty"`                // List of prompts that the key is allowed to use
	Tags                 []string               `json:"tags,omitempty"`                   // Tags for tracking spend and/or doing tag-based routing

	// Budget and limits
	Spend               *float64               `json:"spend,omitempty"`                 // Amount spent by key
	MaxBudget           *float64               `json:"max_budget,omitempty"`            // Specify max budget for a given key
	SoftBudget          *float64               `json:"soft_budget,omitempty"`           // Specify soft budget for a given key
	BudgetDuration      *string                `json:"budget_duration,omitempty"`       // Budget is reset at the end of specified duration
	MaxParallelRequests *int                   `json:"max_parallel_requests,omitempty"` // Rate limit based on parallel requests
	RPMLimit            *int                   `json:"rpm_limit,omitempty"`             // Requests per minute limit
	TPMLimit            *int                   `json:"tpm_limit,omitempty"`             // Tokens per minute limit
	ModelMaxBudget      map[string]interface{} `json:"model_max_budget,omitempty"`      // Model-specific budgets
	ModelRPMLimit       map[string]interface{} `json:"model_rpm_limit,omitempty"`       // Key-specific model rpm limit
	ModelTPMLimit       map[string]interface{} `json:"model_tpm_limit,omitempty"`       // Key-specific model tpm limit

	// Additional configuration
	Metadata        map[string]interface{} `json:"metadata,omitempty"`        // Metadata for key
	SendInviteEmail bool                   `json:"send_invite_email"`         // Whether to send an invite email
	Blocked         bool                   `json:"blocked"`                   // Whether the key is blocked
	EnforcedParams  map[string]interface{} `json:"enforced_params,omitempty"` // List of enforced params for the key
}

// KeyGenerateResponse represents the response from creating a new key
type KeyGenerateResponse struct {
	// Core key fields - these are returned by the API
	Key      string  `json:"key"`                 // The generated API key (sensitive)
	KeyAlias *string `json:"key_alias,omitempty"` // User defined key alias
	KeyName  string  `json:"key_name"`            // The truncated key name for display
	Token    string  `json:"token"`               // The token identifier
	TokenID  string  `json:"token_id"`            // The unique token ID
	BudgetID *string `json:"budget_id,omitempty"` // The budget ID associated with this key

	// Configuration fields
	Models               []string               `json:"models"`                          // List of models this key can access
	Duration             *string                `json:"duration,omitempty"`              // Duration for which this key is valid
	UserID               string                 `json:"user_id"`                         // The user ID associated with this key
	TeamID               *string                `json:"team_id,omitempty"`               // The team ID associated with this key
	MaxParallelRequests  *int                   `json:"max_parallel_requests,omitempty"` // Maximum parallel requests
	Metadata             map[string]interface{} `json:"metadata"`                        // Additional metadata
	TPMLimit             *int                   `json:"tpm_limit,omitempty"`             // Tokens per minute limit
	RPMLimit             *int                   `json:"rpm_limit,omitempty"`             // Requests per minute limit
	BudgetDuration       *string                `json:"budget_duration,omitempty"`       // Budget reset duration
	AllowedCacheControls []string               `json:"allowed_cache_controls"`          // Allowed cache control values
	AllowedRoutes        []string               `json:"allowed_routes"`                  // Allowed API routes
	KeyType              *string                `json:"key_type,omitempty"`              // Type of key

	// Budget and spending fields
	Spend      float64  `json:"spend"`                 // Current spend amount
	MaxBudget  *float64 `json:"max_budget,omitempty"`  // Maximum budget
	SoftBudget *float64 `json:"soft_budget,omitempty"` // Soft budget limit

	// Advanced configuration
	Aliases          map[string]interface{}  `json:"aliases"`                   // Model aliases
	Config           map[string]interface{}  `json:"config"`                    // Additional configuration
	Permissions      map[string]interface{}  `json:"permissions"`               // Permissions configuration
	ObjectPermission interface{}             `json:"object_permission"`         // Object-level permissions
	ModelMaxBudget   map[string]interface{}  `json:"model_max_budget"`          // Per-model budget limits
	ModelRPMLimit    *map[string]interface{} `json:"model_rpm_limit,omitempty"` // Per-model RPM limits
	ModelTPMLimit    *map[string]interface{} `json:"model_tpm_limit,omitempty"` // Per-model TPM limits
	EnforcedParams   *map[string]interface{} `json:"enforced_params,omitempty"` // Enforced parameters

	// Security and control fields
	Guardrails      *[]string `json:"guardrails,omitempty"`        // List of guardrails
	Prompts         *[]string `json:"prompts,omitempty"`           // List of prompts
	Blocked         *bool     `json:"blocked,omitempty"`           // Whether key is blocked
	Tags            *[]string `json:"tags,omitempty"`              // Tags
	SendInviteEmail *bool     `json:"send_invite_email,omitempty"` // Whether invite email was sent

	// Timestamps and audit fields
	Expires   *time.Time `json:"expires,omitempty"` // Expiration timestamp
	CreatedBy string     `json:"created_by"`        // User who created the key
	UpdatedBy string     `json:"updated_by"`        // User who last updated the key
	CreatedAt time.Time  `json:"created_at"`        // Creation timestamp
	UpdatedAt time.Time  `json:"updated_at"`        // Last update timestamp

	// Budget table reference
	LitellmBudgetTable interface{} `json:"litellm_budget_table"` // Reference to budget table
}

// KeyInfoResponse represents the response from the /key/info endpoint
type KeyInfoResponse struct {
	Key  string  `json:"key"`  // The key identifier
	Info KeyInfo `json:"info"` // The key information
}

// KeyInfo represents the detailed information about a key
type KeyInfo struct {
	KeyName              string                 `json:"key_name"`
	KeyAlias             *string                `json:"key_alias"`
	SoftBudgetCooldown   bool                   `json:"soft_budget_cooldown"`
	Spend                float64                `json:"spend"`
	Expires              *time.Time             `json:"expires"`
	Models               []string               `json:"models"`
	Aliases              map[string]interface{} `json:"aliases"`
	Config               map[string]interface{} `json:"config"`
	UserID               string                 `json:"user_id"`
	TeamID               *string                `json:"team_id"`
	Permissions          map[string]interface{} `json:"permissions"`
	MaxParallelRequests  *int                   `json:"max_parallel_requests"`
	Metadata             map[string]interface{} `json:"metadata"`
	Blocked              *bool                  `json:"blocked"`
	TPMLimit             *int                   `json:"tpm_limit"`
	RPMLimit             *int                   `json:"rpm_limit"`
	MaxBudget            *float64               `json:"max_budget"`
	BudgetDuration       *string                `json:"budget_duration"`
	BudgetResetAt        *time.Time             `json:"budget_reset_at"`
	AllowedCacheControls []string               `json:"allowed_cache_controls"`
	AllowedRoutes        []string               `json:"allowed_routes"`
	ModelSpend           map[string]interface{} `json:"model_spend"`
	ModelMaxBudget       map[string]interface{} `json:"model_max_budget"`
	BudgetID             *string                `json:"budget_id"`
	OrganizationID       *string                `json:"organization_id"`
	ObjectPermissionID   *string                `json:"object_permission_id"`
	CreatedAt            time.Time              `json:"created_at"`
	CreatedBy            string                 `json:"created_by"`
	UpdatedAt            time.Time              `json:"updated_at"`
	UpdatedBy            string                 `json:"updated_by"`
	LitellmBudgetTable   interface{}            `json:"litellm_budget_table"`
	LitellmOrgTable      interface{}            `json:"litellm_organization_table"`
	ObjectPermission     interface{}            `json:"object_permission"`
}

// Key represents the internal key structure used by the provider
type Key struct {
	// Core key fields
	Key      string `json:"key,omitempty"`
	KeyAlias string `json:"key_alias,omitempty"`
	KeyName  string `json:"key_name,omitempty"`
	Token    string `json:"token,omitempty"`
	TokenID  string `json:"token_id,omitempty"`
	BudgetID string `json:"budget_id,omitempty"`

	// Configuration fields
	Models               []string               `json:"models"`
	Duration             string                 `json:"duration,omitempty"`
	UserID               string                 `json:"user_id,omitempty"`
	TeamID               string                 `json:"team_id,omitempty"`
	MaxParallelRequests  int                    `json:"max_parallel_requests,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	TPMLimit             int                    `json:"tpm_limit,omitempty"`
	RPMLimit             int                    `json:"rpm_limit,omitempty"`
	BudgetDuration       string                 `json:"budget_duration,omitempty"`
	AllowedCacheControls []string               `json:"allowed_cache_controls,omitempty"`
	AllowedRoutes        []string               `json:"allowed_routes,omitempty"`
	KeyType              string                 `json:"key_type,omitempty"`

	// Budget and spending fields
	Spend      float64 `json:"spend,omitempty"`
	MaxBudget  float64 `json:"max_budget,omitempty"`
	SoftBudget float64 `json:"soft_budget,omitempty"`

	// Advanced configuration
	Aliases          map[string]interface{} `json:"aliases,omitempty"`
	Config           map[string]interface{} `json:"config,omitempty"`
	Permissions      map[string]interface{} `json:"permissions,omitempty"`
	ObjectPermission interface{}            `json:"object_permission,omitempty"`
	ModelMaxBudget   map[string]interface{} `json:"model_max_budget,omitempty"`
	ModelRPMLimit    map[string]interface{} `json:"model_rpm_limit,omitempty"`
	ModelTPMLimit    map[string]interface{} `json:"model_tpm_limit,omitempty"`
	EnforcedParams   map[string]interface{} `json:"enforced_params,omitempty"`

	// Security and control fields
	Guardrails      []string `json:"guardrails,omitempty"`
	Prompts         []string `json:"prompts,omitempty"`
	Blocked         bool     `json:"blocked"`
	Tags            []string `json:"tags,omitempty"`
	SendInviteEmail bool     `json:"send_invite_email,omitempty"`

	// Timestamps and audit fields
	Expires   *time.Time `json:"expires,omitempty"`
	CreatedBy string     `json:"created_by,omitempty"`
	UpdatedBy string     `json:"updated_by,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Budget table reference
	LitellmBudgetTable interface{} `json:"litellm_budget_table,omitempty"`
}
