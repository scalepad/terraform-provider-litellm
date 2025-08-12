package team

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

func createTeam(ctx context.Context, c *litellm.Client, team *Team) (*Team, error) {
	// Generate a UUID for new teams
	if team.TeamID == "" {
		team.TeamID = uuid.New().String()
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/new", team)
	if err != nil {
		return nil, err
	}

	// For create operations, return the team with the generated ID
	return team, nil
}

func getTeam(ctx context.Context, c *litellm.Client, teamID string) (*TeamResponse, error) {
	endpoint := fmt.Sprintf("/team/info?team_id=%s", teamID)
	resp, err := c.SendRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		// Check if it's a not found error
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
			return nil, nil
		}
		return nil, err
	}

	return parseTeamAPIResponse(resp)
}

func updateTeam(ctx context.Context, c *litellm.Client, team *Team) (*Team, error) {
	_, err := c.SendRequest(ctx, http.MethodPost, "/team/update", team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func deleteTeam(ctx context.Context, c *litellm.Client, teamID string) error {
	deleteReq := map[string]interface{}{
		"team_ids": []string{teamID},
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/delete", deleteReq)

	// If it's a not found error, consider it successful (already deleted)
	if err != nil && (strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404")) {
		return nil
	}

	return err
}

func getTeamPermissions(ctx context.Context, c *litellm.Client, teamID string) (*TeamPermissionsResponse, error) {
	endpoint := fmt.Sprintf("/team/permissions_list?team_id=%s", teamID)
	resp, err := c.SendRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Parse the permissions response
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	permResp := &TeamPermissionsResponse{}

	if v, ok := resp["team_id"].(string); ok {
		permResp.TeamID = v
	}

	if permissions, ok := resp["team_member_permissions"].([]interface{}); ok {
		permResp.TeamMemberPermissions = make([]string, len(permissions))
		for i, perm := range permissions {
			if s, ok := perm.(string); ok {
				permResp.TeamMemberPermissions[i] = s
			}
		}
	}

	if allPerms, ok := resp["all_available_permissions"].([]interface{}); ok {
		permResp.AllAvailablePermissions = make([]string, len(allPerms))
		for i, perm := range allPerms {
			if s, ok := perm.(string); ok {
				permResp.AllAvailablePermissions[i] = s
			}
		}
	}

	return permResp, nil
}

func updateTeamPermissions(ctx context.Context, c *litellm.Client, teamID string, permissions []string) error {
	permData := map[string]interface{}{
		"team_id":                 teamID,
		"team_member_permissions": permissions,
	}

	_, err := c.SendRequest(ctx, http.MethodPost, "/team/permissions_update", permData)
	return err
}
