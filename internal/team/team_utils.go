package team

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func buildTeamData(d *schema.ResourceData, teamID string) map[string]interface{} {
	teamData := map[string]interface{}{
		"team_id":    teamID,
		"team_alias": d.Get("team_alias").(string),
	}

	// Add optional fields if they exist
	for _, key := range []string{"organization_id", "metadata", "tpm_limit", "rpm_limit", "max_budget", "budget_duration", "models", "blocked", "team_member_permissions", "team_member_budget"} {
		if v, ok := d.GetOk(key); ok {
			teamData[key] = v
		}
	}

	return teamData
}

func buildTeamDataForUtils(d *schema.ResourceData) map[string]interface{} {
	teamData := make(map[string]interface{})

	// String fields
	utils.GetValueDefault[string](d, "team_alias", teamData)
	utils.GetValueDefault[string](d, "organization_id", teamData)
	utils.GetValueDefault[string](d, "budget_duration", teamData)

	// Int fields
	utils.GetValueDefault[int](d, "tpm_limit", teamData)
	utils.GetValueDefault[int](d, "rpm_limit", teamData)

	// Float64 fields
	utils.GetValueDefault[float64](d, "max_budget", teamData)
	utils.GetValueDefault[float64](d, "team_member_budget", teamData)

	// Bool fields
	utils.GetValueDefault[bool](d, "blocked", teamData)

	// Map fields
	utils.GetValueDefault[map[string]interface{}](d, "metadata", teamData)

	// String list fields
	utils.GetStringListValue(d, "models", teamData)
	utils.GetStringListValue(d, "team_member_permissions", teamData)

	return teamData
}

func setTeamResourceData(d *schema.ResourceData, team *TeamResponse) error {
	fields := map[string]interface{}{
		"team_alias":              team.TeamAlias,
		"organization_id":         team.OrganizationID,
		"tpm_limit":               team.TPMLimit,
		"rpm_limit":               team.RPMLimit,
		"max_budget":              team.MaxBudget,
		"budget_duration":         team.BudgetDuration,
		"team_member_budget":      team.TeamMemberBudget,
		"blocked":                 team.Blocked,
		"team_member_permissions": team.TeamMemberPermissions,
	}

	for field, value := range fields {
		// Use SetIfNotZero to preserve existing values when API doesn't return them
		utils.SetIfNotZero(d, field, value)
	}

	// Handle metadata separately as it's a map
	if team.Metadata != nil {
		d.Set("metadata", team.Metadata)
	}

	// Handle models separately as it's a list
	if team.Models != nil {
		d.Set("models", team.Models)
	} else {
		d.Set("models", d.Get("models"))
	}

	return nil
}

func parseTeamAPIResponse(resp map[string]interface{}) (*TeamResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	teamResp := &TeamResponse{}

	// Parse basic fields
	if v, ok := resp["team_id"].(string); ok {
		teamResp.TeamID = v
	}
	if v, ok := resp["team_alias"].(string); ok {
		teamResp.TeamAlias = v
	}
	if v, ok := resp["organization_id"].(string); ok {
		teamResp.OrganizationID = v
	}
	if v, ok := resp["budget_duration"].(string); ok {
		teamResp.BudgetDuration = v
	}

	// Parse numeric fields
	if v, ok := resp["tpm_limit"].(float64); ok {
		teamResp.TPMLimit = int(v)
	}
	if v, ok := resp["rpm_limit"].(float64); ok {
		teamResp.RPMLimit = int(v)
	}
	if v, ok := resp["max_budget"].(float64); ok {
		teamResp.MaxBudget = v
	}
	if v, ok := resp["team_member_budget"].(float64); ok {
		teamResp.TeamMemberBudget = v
	}

	// Parse boolean fields
	if v, ok := resp["blocked"].(bool); ok {
		teamResp.Blocked = v
	}

	// Parse map fields
	if v, ok := resp["metadata"].(map[string]interface{}); ok {
		teamResp.Metadata = v
	}

	// Parse array fields
	if models, ok := resp["models"].([]interface{}); ok {
		teamResp.Models = make([]string, len(models))
		for i, model := range models {
			if s, ok := model.(string); ok {
				teamResp.Models[i] = s
			}
		}
	}

	if permissions, ok := resp["team_member_permissions"].([]interface{}); ok {
		teamResp.TeamMemberPermissions = make([]string, len(permissions))
		for i, perm := range permissions {
			if s, ok := perm.(string); ok {
				teamResp.TeamMemberPermissions[i] = s
			}
		}
	}

	return teamResp, nil
}

func buildTeamForCreation(data map[string]interface{}) *Team {
	team := &Team{}

	if v, ok := data["team_alias"].(string); ok {
		team.TeamAlias = v
	}
	if v, ok := data["organization_id"].(string); ok {
		team.OrganizationID = v
	}
	if v, ok := data["budget_duration"].(string); ok {
		team.BudgetDuration = v
	}
	if v, ok := data["tpm_limit"].(int); ok {
		team.TPMLimit = v
	}
	if v, ok := data["rpm_limit"].(int); ok {
		team.RPMLimit = v
	}
	if v, ok := data["max_budget"].(float64); ok {
		team.MaxBudget = v
	}
	if v, ok := data["team_member_budget"].(float64); ok {
		team.TeamMemberBudget = v
	}
	if v, ok := data["blocked"].(bool); ok {
		team.Blocked = v
	}
	if v, ok := data["metadata"].(map[string]interface{}); ok {
		team.Metadata = v
	}
	if v, ok := data["models"].([]string); ok {
		team.Models = v
	}
	if v, ok := data["team_member_permissions"].([]string); ok {
		team.TeamMemberPermissions = v
	}

	return team
}
