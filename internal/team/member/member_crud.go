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

		// First try to find the created user in the updated_team_memberships list for budget info
		requestedUserID := request.Member.UserID

		// Look for budget information in updated_team_memberships
		for _, membership := range response.UpdatedTeamMemberships {
			if membership.UserID == requestedUserID {
				var maxBudget float64
				if membership.LitellmBudgetTable.MaxBudget != nil {
					maxBudget = *membership.LitellmBudgetTable.MaxBudget
				}

				// Get user email and role from updated_users or request
				var userEmail string
				var role string = request.Member.Role

				for _, updatedUser := range response.UpdatedUsers {
					if updatedUser.UserID == requestedUserID {
						userEmail = updatedUser.UserEmail
						break
					}
				}

				if userEmail == "" && request.Member.UserEmail != nil {
					userEmail = *request.Member.UserEmail
				}

				return &TeamMemberResponse{
					TeamID:          request.TeamID,
					UserID:          membership.UserID,
					UserEmail:       userEmail,
					Role:            role,
					MaxBudgetInTeam: maxBudget,
					Status:          "active",
				}, nil
			}
		}

		// Fallback: Find the created user in the updated_users list
		for _, updatedUser := range response.UpdatedUsers {
			if updatedUser.UserID == requestedUserID {
				return &TeamMemberResponse{
					TeamID:          request.TeamID,
					UserID:          updatedUser.UserID,
					UserEmail:       updatedUser.UserEmail,
					Role:            request.Member.Role,
					MaxBudgetInTeam: updatedUser.MaxBudget,
					Status:          "active",
				}, nil
			}
		}

		// If not found in response, verify by checking team info
		teamInfo, err := team.GetTeam(ctx, c, request.TeamID)
		if err != nil {
			lastErr = err
			continue
		}

		if teamInfo != nil {
			requestedUserID := request.Member.UserID
			for _, memberWithRole := range teamInfo.TeamInfo.MembersWithRoles {
				if memberWithRole.UserID == requestedUserID {
					var maxBudget float64 = request.MaxBudgetInTeam
					for _, membership := range teamInfo.TeamMemberships {
						if membership.UserID == memberWithRole.UserID && membership.LitellmBudgetTable.MaxBudget != nil {
							maxBudget = *membership.LitellmBudgetTable.MaxBudget
							break
						}
					}
					var userEmail string
					if request.Member.UserEmail != nil {
						userEmail = *request.Member.UserEmail
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

		response, err := litellm.SendRequestTyped[TeamMemberUpdateRequest, TeamMemberUpdateResponse](
			ctx, c, http.MethodPost, "/team/member_update", request,
		)
		if err != nil {
			lastErr = err
			continue
		}

		// Use the update response data directly
		var maxBudget float64
		if response.MaxBudgetInTeam != nil {
			maxBudget = *response.MaxBudgetInTeam
		}

		var userEmail string
		if response.UserEmail != nil {
			userEmail = *response.UserEmail
		}

		// For role, we need to get it from team info since it's not in the update response
		var role string
		teamInfo, err := team.GetTeam(ctx, c, request.TeamID)
		if err != nil {
			lastErr = err
			continue
		}

		if teamInfo != nil {
			for _, memberWithRole := range teamInfo.TeamInfo.MembersWithRoles {
				if memberWithRole.UserID == request.UserID {
					role = memberWithRole.Role
					break
				}
			}
		}

		return &TeamMemberResponse{
			TeamID:          response.TeamID,
			UserID:          response.UserID,
			UserEmail:       userEmail,
			Role:            role,
			MaxBudgetInTeam: maxBudget,
			Status:          "active",
		}, nil
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
	_, err := litellm.SendRequestTypedRateLimited[TeamMemberAdd, TeamMemberCreateResponse](
		ctx, c, http.MethodPost, "/team/member_add", memberAdd,
	)
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
