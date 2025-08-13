package users

// UserInfo represents the user_info object from the API response
type UserInfo struct {
	UserID               string                 `json:"user_id"`
	UserEmail            string                 `json:"user_email,omitempty"`
	UserAlias            string                 `json:"user_alias,omitempty"`
	UserRole             string                 `json:"user_role,omitempty"`
	MaxBudget            float64                `json:"max_budget,omitempty"`
	BudgetDuration       string                 `json:"budget_duration,omitempty"`
	Models               []string               `json:"models,omitempty"`
	TPMLimit             int                    `json:"tpm_limit,omitempty"`
	RPMLimit             int                    `json:"rpm_limit,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	Spend                float64                `json:"spend,omitempty"`
	SSOUserID            string                 `json:"sso_user_id,omitempty"`
	OrganizationID       string                 `json:"organization_id,omitempty"`
	ObjectPermissionID   string                 `json:"object_permission_id,omitempty"`
	Password             string                 `json:"password,omitempty"`
	MaxParallelRequests  int                    `json:"max_parallel_requests,omitempty"`
	BudgetResetAt        string                 `json:"budget_reset_at,omitempty"`
	AllowedCacheControls []string               `json:"allowed_cache_controls,omitempty"`
	ModelSpend           map[string]interface{} `json:"model_spend,omitempty"`
	ModelMaxBudget       map[string]interface{} `json:"model_max_budget,omitempty"`
	CreatedAt            string                 `json:"created_at,omitempty"`
	UpdatedAt            string                 `json:"updated_at,omitempty"`
}

// User represents a LiteLLM user with all supported fields
type User struct {
	UserID               string                 `json:"user_id"`
	UserEmail            string                 `json:"user_email,omitempty"`
	UserAlias            string                 `json:"user_alias,omitempty"`
	UserRole             string                 `json:"user_role,omitempty"`
	MaxBudget            float64                `json:"max_budget,omitempty"`
	BudgetDuration       string                 `json:"budget_duration,omitempty"`
	Models               []string               `json:"models,omitempty"`
	TPMLimit             int                    `json:"tpm_limit,omitempty"`
	RPMLimit             int                    `json:"rpm_limit,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	Spend                float64                `json:"spend,omitempty"`
	KeyCount             int                    `json:"key_count,omitempty"`
	SendInviteEmail      bool                   `json:"send_invite_email"`
	AutoCreateKey        bool                   `json:"auto_create_key"`
	Aliases              map[string]interface{} `json:"aliases,omitempty"`
	Config               map[string]interface{} `json:"config,omitempty"`
	AllowedCacheControls []string               `json:"allowed_cache_controls,omitempty"`
	Blocked              bool                   `json:"blocked"`
	MaxParallelRequests  int                    `json:"max_parallel_requests,omitempty"`
	SoftBudget           float64                `json:"soft_budget,omitempty"`
	ModelMaxBudget       map[string]interface{} `json:"model_max_budget,omitempty"`
	ModelRPMLimit        map[string]interface{} `json:"model_rpm_limit,omitempty"`
	ModelTPMLimit        map[string]interface{} `json:"model_tpm_limit,omitempty"`
	Duration             string                 `json:"duration,omitempty"`
	KeyAlias             string                 `json:"key_alias,omitempty"`
	SSOUserID            string                 `json:"sso_user_id,omitempty"`
	Prompts              []string               `json:"prompts,omitempty"`
	Organizations        []string               `json:"organizations,omitempty"`
	CreatedAt            string                 `json:"created_at,omitempty"`
	UpdatedAt            string                 `json:"updated_at,omitempty"`
	BudgetResetAt        string                 `json:"budget_reset_at,omitempty"`
}

// UserCreateRequest represents the request body for creating a user
type UserCreateRequest struct {
	UserID               string                 `json:"user_id,omitempty"`
	UserEmail            string                 `json:"user_email,omitempty"`
	UserAlias            string                 `json:"user_alias,omitempty"`
	UserRole             string                 `json:"user_role,omitempty"`
	MaxBudget            float64                `json:"max_budget,omitempty"`
	BudgetDuration       string                 `json:"budget_duration,omitempty"`
	Models               []string               `json:"models,omitempty"`
	TPMLimit             int                    `json:"tpm_limit,omitempty"`
	RPMLimit             int                    `json:"rpm_limit,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	SendInviteEmail      bool                   `json:"send_invite_email"`
	AutoCreateKey        bool                   `json:"auto_create_key"`
	Aliases              map[string]interface{} `json:"aliases,omitempty"`
	Config               map[string]interface{} `json:"config,omitempty"`
	AllowedCacheControls []string               `json:"allowed_cache_controls,omitempty"`
	Blocked              bool                   `json:"blocked"`
	MaxParallelRequests  int                    `json:"max_parallel_requests,omitempty"`
	SoftBudget           float64                `json:"soft_budget,omitempty"`
	ModelMaxBudget       map[string]interface{} `json:"model_max_budget,omitempty"`
	ModelRPMLimit        map[string]interface{} `json:"model_rpm_limit,omitempty"`
	ModelTPMLimit        map[string]interface{} `json:"model_tpm_limit,omitempty"`
	Duration             string                 `json:"duration,omitempty"`
	KeyAlias             string                 `json:"key_alias,omitempty"`
	SSOUserID            string                 `json:"sso_user_id,omitempty"`
	Prompts              []string               `json:"prompts,omitempty"`
	Organizations        []string               `json:"organizations,omitempty"`
}

// UserUpdateRequest represents the request body for updating a user
type UserUpdateRequest struct {
	UserID               string                 `json:"user_id"`
	UserEmail            string                 `json:"user_email,omitempty"`
	UserAlias            string                 `json:"user_alias,omitempty"`
	UserRole             string                 `json:"user_role,omitempty"`
	MaxBudget            float64                `json:"max_budget,omitempty"`
	BudgetDuration       string                 `json:"budget_duration,omitempty"`
	Models               []string               `json:"models,omitempty"`
	TPMLimit             int                    `json:"tpm_limit,omitempty"`
	RPMLimit             int                    `json:"rpm_limit,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	Aliases              map[string]interface{} `json:"aliases,omitempty"`
	Config               map[string]interface{} `json:"config,omitempty"`
	AllowedCacheControls []string               `json:"allowed_cache_controls,omitempty"`
	Blocked              bool                   `json:"blocked"`
	MaxParallelRequests  int                    `json:"max_parallel_requests,omitempty"`
	SoftBudget           float64                `json:"soft_budget,omitempty"`
	ModelMaxBudget       map[string]interface{} `json:"model_max_budget,omitempty"`
	ModelRPMLimit        map[string]interface{} `json:"model_rpm_limit,omitempty"`
	ModelTPMLimit        map[string]interface{} `json:"model_tpm_limit,omitempty"`
	SSOUserID            string                 `json:"sso_user_id,omitempty"`
	Prompts              []string               `json:"prompts,omitempty"`
	Organizations        []string               `json:"organizations,omitempty"`
}

// UserDeleteRequest represents the request body for deleting users
type UserDeleteRequest struct {
	UserIDs []string `json:"user_ids"`
}

// UserResponse represents the API response when retrieving user info
type UserResponse struct {
	UserID   string        `json:"user_id"`
	UserInfo UserInfo      `json:"user_info"`
	Keys     []interface{} `json:"keys"`
	Teams    []interface{} `json:"teams"`
}
