package key

type Key struct {
	Key                  string                 `json:"key,omitempty"`
	Models               []string               `json:"models"`
	Spend                float64                `json:"spend,omitempty"`
	MaxBudget            float64                `json:"max_budget,omitempty"`
	UserID               string                 `json:"user_id,omitempty"`
	TeamID               string                 `json:"team_id,omitempty"`
	MaxParallelRequests  int                    `json:"max_parallel_requests,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	TPMLimit             int                    `json:"tpm_limit,omitempty"`
	RPMLimit             int                    `json:"rpm_limit,omitempty"`
	BudgetDuration       string                 `json:"budget_duration,omitempty"`
	AllowedCacheControls []string               `json:"allowed_cache_controls,omitempty"`
	SoftBudget           float64                `json:"soft_budget,omitempty"`
	KeyAlias             string                 `json:"key_alias,omitempty"`
	Duration             string                 `json:"duration,omitempty"`
	Aliases              map[string]interface{} `json:"aliases,omitempty"`
	Config               map[string]interface{} `json:"config,omitempty"`
	Permissions          map[string]interface{} `json:"permissions,omitempty"`
	ModelMaxBudget       map[string]interface{} `json:"model_max_budget,omitempty"`
	ModelRPMLimit        map[string]interface{} `json:"model_rpm_limit,omitempty"`
	ModelTPMLimit        map[string]interface{} `json:"model_tpm_limit,omitempty"`
	Guardrails           []string               `json:"guardrails,omitempty"`
	Blocked              bool                   `json:"blocked"`
	Tags                 []string               `json:"tags,omitempty"`
	SendInviteEmail      bool                   `json:"send_invite_email,omitempty"`
	KeyType              string                 `json:"key_type,omitempty"`
}
