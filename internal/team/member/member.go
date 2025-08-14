package member

// TeamMember represents a team member for input operations
type TeamMember struct {
	TeamID          string  `json:"team_id"`
	UserID          string  `json:"user_id"`
	UserEmail       string  `json:"user_email"`
	Role            string  `json:"role"`
	MaxBudgetInTeam float64 `json:"max_budget_in_team"`
}

// TeamMemberResponse represents the API response for team member operations
type TeamMemberResponse struct {
	TeamID          string  `json:"team_id"`
	UserID          string  `json:"user_id"`
	UserEmail       string  `json:"user_email"`
	Role            string  `json:"role"`
	MaxBudgetInTeam float64 `json:"max_budget_in_team"`
	Status          string  `json:"status,omitempty"`
}

// TeamMemberAdd represents multiple team members for bulk operations
type TeamMemberAdd struct {
	TeamID          string                `json:"team_id"`
	Members         []TeamMemberAddMember `json:"member"`
	MaxBudgetInTeam float64               `json:"max_budget_in_team"`
}

// TeamMemberAddMember represents a single member in bulk operations
type TeamMemberAddMember struct {
	UserID    string `json:"user_id,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
	Role      string `json:"role"`
}

// TeamMemberCreateRequest represents the request for creating a team member
type TeamMemberCreateRequest struct {
	TeamID          string                   `json:"team_id"`
	Member          []TeamMemberCreateMember `json:"member"`
	MaxBudgetInTeam float64                  `json:"max_budget_in_team"`
}

// TeamMemberCreateMember represents a single member in the create request
type TeamMemberCreateMember struct {
	UserID    string `json:"user_id,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
	Role      string `json:"role"`
}

// TeamMemberCreateResponse represents the response from creating a team member
type TeamMemberCreateResponse struct {
	TeamAlias              string                 `json:"team_alias"`
	TeamID                 string                 `json:"team_id"`
	OrganizationID         *string                `json:"organization_id"`
	Admins                 []string               `json:"admins"`
	Members                []string               `json:"members"`
	MembersWithRoles       []MemberWithRole       `json:"members_with_roles"`
	TeamMemberPermissions  []string               `json:"team_member_permissions"`
	Metadata               map[string]interface{} `json:"metadata"`
	TPMLimit               *int                   `json:"tpm_limit"`
	RPMLimit               *int                   `json:"rpm_limit"`
	MaxBudget              *float64               `json:"max_budget"`
	BudgetDuration         *string                `json:"budget_duration"`
	Models                 []string               `json:"models"`
	Blocked                bool                   `json:"blocked"`
	Spend                  float64                `json:"spend"`
	MaxParallelRequests    *int                   `json:"max_parallel_requests"`
	BudgetResetAt          *string                `json:"budget_reset_at"`
	ModelID                *string                `json:"model_id"`
	LitellmModelTable      interface{}            `json:"litellm_model_table"`
	ObjectPermission       interface{}            `json:"object_permission"`
	UpdatedAt              string                 `json:"updated_at"`
	CreatedAt              string                 `json:"created_at"`
	ObjectPermissionID     *string                `json:"object_permission_id"`
	UpdatedUsers           []UpdatedUser          `json:"updated_users"`
	UpdatedTeamMemberships []interface{}          `json:"updated_team_memberships"`
}

// MemberWithRole represents a team member with their role
type MemberWithRole struct {
	UserID    string  `json:"user_id"`
	UserEmail *string `json:"user_email"`
	Role      string  `json:"role"`
}

// UpdatedUser represents a user in the updated_users list
type UpdatedUser struct {
	UserID                  string                 `json:"user_id"`
	MaxBudget               float64                `json:"max_budget"`
	Spend                   float64                `json:"spend"`
	ModelMaxBudget          map[string]interface{} `json:"model_max_budget"`
	ModelSpend              map[string]interface{} `json:"model_spend"`
	UserEmail               string                 `json:"user_email"`
	UserAlias               string                 `json:"user_alias"`
	Models                  []string               `json:"models"`
	TPMLimit                int                    `json:"tpm_limit"`
	RPMLimit                int                    `json:"rpm_limit"`
	UserRole                string                 `json:"user_role"`
	OrganizationMemberships interface{}            `json:"organization_memberships"`
	Teams                   []string               `json:"teams"`
	SSOUserID               *string                `json:"sso_user_id"`
	BudgetDuration          string                 `json:"budget_duration"`
	BudgetResetAt           string                 `json:"budget_reset_at"`
	Metadata                map[string]interface{} `json:"metadata"`
	CreatedAt               string                 `json:"created_at"`
	UpdatedAt               string                 `json:"updated_at"`
	ObjectPermission        interface{}            `json:"object_permission"`
}

// TeamMemberUpdateRequest represents the request for updating a team member
type TeamMemberUpdateRequest struct {
	UserID          string   `json:"user_id"`                      // Required for identification
	TeamID          string   `json:"team_id"`                      // Required for identification
	UserEmail       *string  `json:"user_email,omitempty"`         // Optional - only set if changed
	MaxBudgetInTeam *float64 `json:"max_budget_in_team,omitempty"` // Optional - only set if changed
	Role            *string  `json:"role,omitempty"`               // Optional - only set if changed
}

// TeamMemberUpdateResponse represents the response from updating a team member
type TeamMemberUpdateResponse struct {
	// Assuming the update response is similar to the create response
	TeamMemberCreateResponse
}
