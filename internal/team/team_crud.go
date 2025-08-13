package team

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// createTeam creates a new team using the typed request/response pattern
func createTeam(ctx context.Context, c *litellm.Client, request *TeamCreateRequest) (*TeamCreateResponse, error) {
	response, err := litellm.SendRequestTyped[TeamCreateRequest, TeamCreateResponse](
		ctx, c, http.MethodPost, "/team/new", request,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create team: %w", err)
	}

	return response, nil
}

// getTeam retrieves team information by team ID
func getTeam(ctx context.Context, c *litellm.Client, teamID string) (*TeamInfoResponse, error) {
	response, err := litellm.SendRequestTyped[interface{}, TeamInfoResponse](
		ctx, c, http.MethodGet, fmt.Sprintf("/team/info?team_id=%s", url.QueryEscape(teamID)), nil,
	)
	if err != nil {
		// Check if it's a not found error
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get team: %w", err)
	}

	return response, nil
}

// updateTeam updates an existing team using the typed request pattern
func updateTeam(ctx context.Context, c *litellm.Client, request *TeamUpdateRequest) (*TeamCreateResponse, error) {
	response, err := litellm.SendRequestTyped[TeamUpdateRequest, TeamCreateResponse](
		ctx, c, http.MethodPost, "/team/update", request,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update team: %w", err)
	}

	return response, nil
}

// deleteTeam deletes a team by team ID
func deleteTeam(ctx context.Context, c *litellm.Client, teamID string) error {
	deleteRequest := &TeamDeleteRequest{
		TeamIDs: []string{teamID},
	}

	_, err := litellm.SendRequestTyped[TeamDeleteRequest, interface{}](
		ctx, c, http.MethodPost, "/team/delete", deleteRequest,
	)
	if err != nil {
		// If it's a not found error, consider it successful (already deleted)
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
			return nil
		}
		return fmt.Errorf("failed to delete team: %w", err)
	}

	return nil
}

// getTeamPermissions retrieves team permissions by team ID
func getTeamPermissions(ctx context.Context, c *litellm.Client, teamID string) (*TeamPermissionsResponse, error) {
	response, err := litellm.SendRequestTyped[interface{}, TeamPermissionsResponse](
		ctx, c, http.MethodGet, fmt.Sprintf("/team/permissions_list?team_id=%s", url.QueryEscape(teamID)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get team permissions: %w", err)
	}

	return response, nil
}

// updateTeamPermissions updates team permissions
func updateTeamPermissions(ctx context.Context, c *litellm.Client, teamID string, permissions []string) error {
	permData := map[string]interface{}{
		"team_id":                 teamID,
		"team_member_permissions": permissions,
	}

	_, err := litellm.SendRequestTyped[map[string]interface{}, interface{}](
		ctx, c, http.MethodPost, "/team/permissions_update", &permData,
	)
	if err != nil {
		return fmt.Errorf("failed to update team permissions: %w", err)
	}

	return nil
}
