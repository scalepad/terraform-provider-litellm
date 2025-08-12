package team

// Team represents a team configuration in LiteLLM
type Team struct {
	TeamID                string                 `json:"team_id,omitempty"`
	TeamAlias             string                 `json:"team_alias,omitempty"`
	OrganizationID        string                 `json:"organization_id,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	TPMLimit              int                    `json:"tpm_limit,omitempty"`
	RPMLimit              int                    `json:"rpm_limit,omitempty"`
	MaxBudget             float64                `json:"max_budget,omitempty"`
	BudgetDuration        string                 `json:"budget_duration,omitempty"`
	TeamMemberBudget      float64                `json:"team_member_budget,omitempty"`
	Models                []string               `json:"models,omitempty"`
	Blocked               bool                   `json:"blocked,omitempty"`
	TeamMemberPermissions []string               `json:"team_member_permissions,omitempty"`
}

// TeamResponse represents a response from the API containing team information.
type TeamResponse struct {
	TeamID                string                 `json:"team_id,omitempty"`
	TeamAlias             string                 `json:"team_alias,omitempty"`
	OrganizationID        string                 `json:"organization_id,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	TPMLimit              int                    `json:"tpm_limit,omitempty"`
	RPMLimit              int                    `json:"rpm_limit,omitempty"`
	MaxBudget             float64                `json:"max_budget,omitempty"`
	BudgetDuration        string                 `json:"budget_duration,omitempty"`
	TeamMemberBudget      float64                `json:"team_member_budget,omitempty"`
	Models                []string               `json:"models"`
	Blocked               bool                   `json:"blocked,omitempty"`
	TeamMemberPermissions []string               `json:"team_member_permissions,omitempty"`
}

// TeamPermissionsResponse represents a response from the API containing team permissions information.
type TeamPermissionsResponse struct {
	TeamID                  string   `json:"team_id"`
	TeamMemberPermissions   []string `json:"team_member_permissions"`
	AllAvailablePermissions []string `json:"all_available_permissions"`
}
