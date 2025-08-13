package team

import "time"

// TeamCreateRequest represents the request payload for creating a new team
type TeamCreateRequest struct {
	// Core configuration
	TeamAlias      *string `json:"team_alias,omitempty"`      // User defined team alias
	TeamID         *string `json:"team_id,omitempty"`         // The team id of the user. If none passed, we'll generate it.
	OrganizationID *string `json:"organization_id,omitempty"` // The organization id of the team. Default is None. Create via /organization/new.

	// Members and roles
	MembersWithRoles []MemberWithRole `json:"members_with_roles,omitempty"` // A list of users and their roles in the team
	Members          []string         `json:"members,omitempty"`            // Control team members via /team/member/add and /team/member/delete (deprecated)
	Admins           []string         `json:"admins,omitempty"`             // A list of user_id's for the admin role (deprecated)
	Users            []string         `json:"users,omitempty"`              // A list of user_id's for the user role (deprecated)

	// Permissions and access control
	TeamMemberPermissions []string               `json:"team_member_permissions,omitempty"` // A list of routes that non-admin team members can access
	Models                []string               `json:"models,omitempty"`                  // A list of models associated with the team
	ModelAliases          map[string]interface{} `json:"model_aliases,omitempty"`           // Model aliases for the team
	Guardrails            []string               `json:"guardrails,omitempty"`              // Guardrails for the team
	Prompts               []string               `json:"prompts,omitempty"`                 // List of prompts that the team is allowed to use
	Tags                  []string               `json:"tags,omitempty"`                    // Tags for tracking spend and/or doing tag-based routing

	// Budget and limits
	TPMLimit       *int     `json:"tpm_limit,omitempty"`       // The TPM (Tokens Per Minute) limit for this team
	RPMLimit       *int     `json:"rpm_limit,omitempty"`       // The RPM (Requests Per Minute) limit for this team
	MaxBudget      *float64 `json:"max_budget,omitempty"`      // The maximum budget allocated to the team
	BudgetDuration *string  `json:"budget_duration,omitempty"` // The duration of the budget for the team

	// Team member configuration
	TeamMemberBudget      *float64 `json:"team_member_budget,omitempty"`       // The maximum budget allocated to an individual team member
	TeamMemberKeyDuration *string  `json:"team_member_key_duration,omitempty"` // The duration for a team member's key

	// Additional configuration
	Metadata         map[string]interface{} `json:"metadata,omitempty"`          // Metadata for team, store information for team
	Blocked          bool                   `json:"blocked"`                     // Flag indicating if the team is blocked or not
	ObjectPermission *ObjectPermissionBase  `json:"object_permission,omitempty"` // team-specific object permission
}

// TeamUpdateRequest represents the request payload for updating a team
type TeamUpdateRequest struct {
	// Required field
	TeamID string `json:"team_id"` // The team id of the user. Required param.

	// Core configuration
	TeamAlias      *string `json:"team_alias,omitempty"`      // User defined team alias
	OrganizationID *string `json:"organization_id,omitempty"` // The organization id of the team

	// Permissions and access control
	TeamMemberPermissions []string               `json:"team_member_permissions,omitempty"` // A list of routes that non-admin team members can access
	Models                []string               `json:"models,omitempty"`                  // A list of models associated with the team
	ModelAliases          map[string]interface{} `json:"model_aliases,omitempty"`           // Model aliases for the team
	Guardrails            []string               `json:"guardrails,omitempty"`              // Guardrails for the team
	Prompts               []string               `json:"prompts,omitempty"`                 // List of prompts that the team is allowed to use
	Tags                  []string               `json:"tags,omitempty"`                    // Tags for tracking spend and/or doing tag-based routing

	// Budget and limits
	TPMLimit       *int     `json:"tpm_limit,omitempty"`       // The TPM (Tokens Per Minute) limit for this team
	RPMLimit       *int     `json:"rpm_limit,omitempty"`       // The RPM (Requests Per Minute) limit for this team
	MaxBudget      *float64 `json:"max_budget,omitempty"`      // The maximum budget allocated to the team
	BudgetDuration *string  `json:"budget_duration,omitempty"` // The duration of the budget for the team

	// Team member configuration
	TeamMemberBudget      *float64 `json:"team_member_budget,omitempty"`       // The maximum budget allocated to an individual team member
	TeamMemberKeyDuration *string  `json:"team_member_key_duration,omitempty"` // The duration for a team member's key

	// Additional configuration
	Metadata         map[string]interface{} `json:"metadata,omitempty"`          // Metadata for team, store information for team
	Blocked          bool                   `json:"blocked"`                     // Flag indicating if the team is blocked or not
	ObjectPermission *ObjectPermissionBase  `json:"object_permission,omitempty"` // team-specific object permission
}

// TeamCreateResponse represents the response from creating a new team
type TeamCreateResponse struct {
	TeamID string `json:"team_id"` // Unique team id - used for tracking spend across multiple keys for same team id
}

// TeamInfoResponse represents the response from the /team/info endpoint
type TeamInfoResponse struct {
	TeamID          string           `json:"team_id"`
	TeamInfo        TeamInfo         `json:"team_info"`
	Keys            []interface{}    `json:"keys"`
	TeamMemberships []TeamMembership `json:"team_memberships"`
}

// TeamInfo represents the detailed information about a team
type TeamInfo struct {
	TeamAlias             string                 `json:"team_alias"`
	TeamID                string                 `json:"team_id"`
	OrganizationID        *string                `json:"organization_id"`
	Admins                []string               `json:"admins"`
	Members               []string               `json:"members"`
	MembersWithRoles      []MemberWithRole       `json:"members_with_roles"`
	TeamMemberPermissions []string               `json:"team_member_permissions"`
	Metadata              map[string]interface{} `json:"metadata"`
	TPMLimit              *int                   `json:"tpm_limit"`
	RPMLimit              *int                   `json:"rpm_limit"`
	MaxBudget             *float64               `json:"max_budget"`
	BudgetDuration        *string                `json:"budget_duration"`
	Models                []string               `json:"models"`
	Blocked               bool                   `json:"blocked"`
	Spend                 float64                `json:"spend"`
	MaxParallelRequests   *int                   `json:"max_parallel_requests"`
	BudgetResetAt         *time.Time             `json:"budget_reset_at"`
	ModelID               *string                `json:"model_id"`
	LitellmModelTable     interface{}            `json:"litellm_model_table"`
	ObjectPermission      *ObjectPermissionBase  `json:"object_permission"`
	UpdatedAt             time.Time              `json:"updated_at"`
	CreatedAt             time.Time              `json:"created_at"`
	ObjectPermissionID    *string                `json:"object_permission_id"`
	TeamMemberBudgetTable interface{}            `json:"team_member_budget_table"`
}

// MemberWithRole represents a team member with their role
type MemberWithRole struct {
	UserID    string  `json:"user_id"`
	UserEmail *string `json:"user_email"`
	Role      string  `json:"role"` // "admin" or "user"
}

// TeamMembership represents a team membership with budget information
type TeamMembership struct {
	UserID             string             `json:"user_id"`
	TeamID             string             `json:"team_id"`
	BudgetID           string             `json:"budget_id"`
	Spend              float64            `json:"spend"`
	LitellmBudgetTable LitellmBudgetTable `json:"litellm_budget_table"`
}

// LitellmBudgetTable represents budget table information
type LitellmBudgetTable struct {
	BudgetID            string      `json:"budget_id"`
	SoftBudget          *float64    `json:"soft_budget"`
	MaxBudget           *float64    `json:"max_budget"`
	MaxParallelRequests *int        `json:"max_parallel_requests"`
	TPMLimit            *int        `json:"tpm_limit"`
	RPMLimit            *int        `json:"rpm_limit"`
	ModelMaxBudget      interface{} `json:"model_max_budget"`
	BudgetDuration      *string     `json:"budget_duration"`
}

// ObjectPermissionBase represents object-level permissions
type ObjectPermissionBase struct {
	VectorStores []string `json:"vector_stores,omitempty"`
	// Add other object permission fields as needed
}

// TeamPermissionsResponse represents a response from the API containing team permissions information
type TeamPermissionsResponse struct {
	TeamID                  string   `json:"team_id"`
	TeamMemberPermissions   []string `json:"team_member_permissions"`
	AllAvailablePermissions []string `json:"all_available_permissions"`
}

// TeamDeleteRequest represents the request for deleting teams
type TeamDeleteRequest struct {
	TeamIDs []string `json:"team_ids"`
}
