package member

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"github.com/scalepad/terraform-provider-litellm/internal/team"
)

// createTeamMember creates a new team member using the typed request/response pattern
func createTeamMember(ctx context.Context, c *litellm.Client, request *TeamMemberCreateRequest) (*TeamMemberResponse, error) {
	maxRetries := 3
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * time.Second) // Progressive backoff: 1s, 2s
		}

		response, err := litellm.SendRequestTypedRateLimited[TeamMemberCreateRequest, TeamMemberCreateResponse](
			ctx, c, http.MethodPost, "/team/member_add", request,
		)
		if err != nil {
			lastErr = err
			continue
		}

		// Find the created user in the updated_users list
		if len(request.Member) > 0 {
			requestedUserID := request.Member[0].UserID
			for _, updatedUser := range response.UpdatedUsers {
				if updatedUser.UserID == requestedUserID {
					return &TeamMemberResponse{
						TeamID:          request.TeamID,
						UserID:          updatedUser.UserID,
						UserEmail:       updatedUser.UserEmail,
						Role:            request.Member[0].Role,
						MaxBudgetInTeam: updatedUser.MaxBudget,
						Status:          "active",
					}, nil
				}
			}
		}

		// If not found in response, verify by checking team info
		teamInfo, err := team.GetTeam(ctx, c, request.TeamID)
		if err != nil {
			lastErr = err
			continue
		}

		if teamInfo != nil && len(request.Member) > 0 {
			requestedUserID := request.Member[0].UserID
			for _, memberWithRole := range teamInfo.TeamInfo.MembersWithRoles {
				if memberWithRole.UserID == requestedUserID {
					var maxBudget float64 = request.MaxBudgetInTeam
					for _, membership := range teamInfo.TeamMemberships {
						if membership.UserID == memberWithRole.UserID && membership.LitellmBudgetTable.MaxBudget != nil {
							maxBudget = *membership.LitellmBudgetTable.MaxBudget
							break
						}
					}
					return &TeamMemberResponse{
						TeamID:          request.TeamID,
						UserID:          memberWithRole.UserID,
						UserEmail:       request.Member[0].UserEmail,
						Role:            memberWithRole.Role,
						MaxBudgetInTeam: maxBudget,
						Status:          "active",
					}, nil
				}
			}
		}
	}

	if lastErr != nil {
		return nil, fmt.Errorf("failed to create team member after %d attempts: %w", maxRetries, lastErr)
	}
	return nil, fmt.Errorf("team member was not found after %d attempts", maxRetries)
}

// updateTeamMember updates an existing team member
func updateTeamMember(ctx context.Context, c *litellm.Client, request *TeamMemberUpdateRequest) (*TeamMemberResponse, error) {
	maxRetries := 3
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * time.Second) // Progressive backoff: 1s, 2s
		}

		_, err := litellm.SendRequestTyped[TeamMemberUpdateRequest, TeamMemberUpdateResponse](
			ctx, c, http.MethodPost, "/team/member_update", request,
		)
		if err != nil {
			lastErr = err
			continue
		}

		// Verify the update by checking team info
		teamInfo, err := team.GetTeam(ctx, c, request.TeamID)
		if err != nil {
			lastErr = err
			continue
		}

		if teamInfo != nil {
			for _, memberWithRole := range teamInfo.TeamInfo.MembersWithRoles {
				if memberWithRole.UserID == request.UserID {
					var maxBudget float64
					if request.MaxBudgetInTeam != nil {
						maxBudget = *request.MaxBudgetInTeam
					}
					for _, membership := range teamInfo.TeamMemberships {
						if membership.UserID == memberWithRole.UserID && membership.LitellmBudgetTable.MaxBudget != nil {
							maxBudget = *membership.LitellmBudgetTable.MaxBudget
							break
						}
					}

					var userEmail string
					if request.UserEmail != nil {
						userEmail = *request.UserEmail
					}

					return &TeamMemberResponse{
						TeamID:          request.TeamID,
						UserID:          memberWithRole.UserID,
						UserEmail:       userEmail,
						Role:            memberWithRole.Role,
						MaxBudgetInTeam: maxBudget,
						Status:          "active",
					}, nil
				}
			}
		}
	}

	if lastErr != nil {
		return nil, fmt.Errorf("failed to update team member after %d attempts: %w", maxRetries, lastErr)
	}
	return nil, fmt.Errorf("team member was not found after update after %d attempts", maxRetries)
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
