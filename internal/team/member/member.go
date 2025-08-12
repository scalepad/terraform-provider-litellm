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
