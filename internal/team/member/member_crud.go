package member

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// createTeamMember creates a new team member
func createTeamMember(ctx context.Context, c *litellm.Client, member *TeamMember) (*TeamMemberResponse, error) {
	memberData := map[string]interface{}{
		"member": []map[string]interface{}{
			{
				"role":       member.Role,
				"user_id":    member.UserID,
				"user_email": member.UserEmail,
			},
		},
		"team_id":            member.TeamID,
		"max_budget_in_team": member.MaxBudgetInTeam,
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/member_add", memberData)
	if err != nil {
		return nil, err
	}

	// For team member creation, we typically don't get detailed response data
	// Return a response based on the input data
	return &TeamMemberResponse{
		TeamID:          member.TeamID,
		UserID:          member.UserID,
		UserEmail:       member.UserEmail,
		Role:            member.Role,
		MaxBudgetInTeam: member.MaxBudgetInTeam,
		Status:          "active",
	}, nil
}

// updateTeamMember updates an existing team member
func updateTeamMember(ctx context.Context, c *litellm.Client, member *TeamMember) (*TeamMemberResponse, error) {
	updateData := map[string]interface{}{
		"user_id":            member.UserID,
		"user_email":         member.UserEmail,
		"team_id":            member.TeamID,
		"max_budget_in_team": member.MaxBudgetInTeam,
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/member_update", updateData)
	if err != nil {
		return nil, err
	}

	// Return updated member data
	return &TeamMemberResponse{
		TeamID:          member.TeamID,
		UserID:          member.UserID,
		UserEmail:       member.UserEmail,
		Role:            member.Role,
		MaxBudgetInTeam: member.MaxBudgetInTeam,
		Status:          "active",
	}, nil
}

// deleteTeamMember deletes a team member
func deleteTeamMember(ctx context.Context, c *litellm.Client, teamID, userID, userEmail string) error {
	deleteData := map[string]interface{}{
		"user_id":    userID,
		"user_email": userEmail,
		"team_id":    teamID,
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/member_delete", deleteData)

	// If it's a not found error, consider it successful (already deleted)
	if err != nil && (strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404")) {
		return nil
	}

	return err
}

// createTeamMembersBulk creates multiple team members in bulk
func createTeamMembersBulk(ctx context.Context, c *litellm.Client, memberAdd *TeamMemberAdd) error {
	memberData := map[string]interface{}{
		"member":             memberAdd.Members,
		"team_id":            memberAdd.TeamID,
		"max_budget_in_team": memberAdd.MaxBudgetInTeam,
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/member_add", memberData)
	return err
}

// updateTeamMemberBudget updates the budget for a team member
func updateTeamMemberBudget(ctx context.Context, c *litellm.Client, teamID, userID, userEmail, role string, maxBudget float64) error {
	updateData := map[string]interface{}{
		"team_id":            teamID,
		"role":               role,
		"max_budget_in_team": maxBudget,
	}

	if userID != "" {
		updateData["user_id"] = userID
	}
	if userEmail != "" {
		updateData["user_email"] = userEmail
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/member_update", updateData)
	return err
}

// updateTeamMemberRole updates the role for a team member
func updateTeamMemberRole(ctx context.Context, c *litellm.Client, teamID, userID, userEmail, role string, maxBudget float64) error {
	updateData := map[string]interface{}{
		"team_id":            teamID,
		"role":               role,
		"max_budget_in_team": maxBudget,
	}

	if userID != "" {
		updateData["user_id"] = userID
	}
	if userEmail != "" {
		updateData["user_email"] = userEmail
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/member_update", updateData)
	return err
}

// getTeamMember retrieves a team member (note: API doesn't provide direct endpoint)
func getTeamMember(ctx context.Context, c *litellm.Client, teamID, userID string) (*TeamMemberResponse, error) {
	// The LiteLLM API doesn't provide a direct endpoint to get a single team member
	// This would typically require getting the entire team and filtering
	// For now, we return nil to indicate the member should be read from state
	return nil, fmt.Errorf("team member read not supported by API - maintaining state")
}
