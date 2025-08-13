package team

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

// buildTeamCreateRequest builds a TeamCreateRequest from Terraform resource data
func buildTeamCreateRequest(d *schema.ResourceData) *TeamCreateRequest {
	request := &TeamCreateRequest{}

	// String fields
	if v, ok := d.GetOk("team_alias"); ok {
		alias := v.(string)
		request.TeamAlias = &alias
	}
	if v, ok := d.GetOk("organization_id"); ok {
		orgID := v.(string)
		request.OrganizationID = &orgID
	}
	if v, ok := d.GetOk("budget_duration"); ok {
		duration := v.(string)
		request.BudgetDuration = &duration
	}

	// Int fields
	if v, ok := d.GetOk("tpm_limit"); ok {
		limit := v.(int)
		request.TPMLimit = &limit
	}
	if v, ok := d.GetOk("rpm_limit"); ok {
		limit := v.(int)
		request.RPMLimit = &limit
	}

	// Float64 fields
	if v, ok := d.GetOk("max_budget"); ok {
		budget := v.(float64)
		request.MaxBudget = &budget
	}
	if v, ok := d.GetOk("team_member_budget"); ok {
		budget := v.(float64)
		request.TeamMemberBudget = &budget
	}

	// Bool fields
	if v, ok := d.GetOk("blocked"); ok {
		request.Blocked = v.(bool)
	}

	// Map fields
	if v, ok := d.GetOk("metadata"); ok {
		request.Metadata = v.(map[string]interface{})
	}

	// String list fields
	if v, ok := d.GetOk("models"); ok {
		models := make([]string, 0)
		for _, model := range v.([]interface{}) {
			if s, ok := model.(string); ok {
				models = append(models, s)
			}
		}
		request.Models = models
	}

	if v, ok := d.GetOk("team_member_permissions"); ok {
		permissions := make([]string, 0)
		for _, perm := range v.([]interface{}) {
			if s, ok := perm.(string); ok {
				permissions = append(permissions, s)
			}
		}
		request.TeamMemberPermissions = permissions
	}

	return request
}

// buildTeamUpdateRequest builds a TeamUpdateRequest from Terraform resource data
// Only includes fields that have changed
func buildTeamUpdateRequest(d *schema.ResourceData, teamID string) *TeamUpdateRequest {
	request := &TeamUpdateRequest{
		TeamID: teamID,
	}

	// String fields - only set if changed
	if d.HasChange("team_alias") {
		if v, ok := d.GetOk("team_alias"); ok {
			alias := v.(string)
			request.TeamAlias = &alias
		}
	}
	if d.HasChange("organization_id") {
		if v, ok := d.GetOk("organization_id"); ok {
			orgID := v.(string)
			request.OrganizationID = &orgID
		}
	}
	if d.HasChange("budget_duration") {
		if v, ok := d.GetOk("budget_duration"); ok {
			duration := v.(string)
			request.BudgetDuration = &duration
		}
	}

	// Int fields - only set if changed
	if d.HasChange("tpm_limit") {
		if v, ok := d.GetOk("tpm_limit"); ok {
			limit := v.(int)
			request.TPMLimit = &limit
		}
	}
	if d.HasChange("rpm_limit") {
		if v, ok := d.GetOk("rpm_limit"); ok {
			limit := v.(int)
			request.RPMLimit = &limit
		}
	}

	// Float64 fields - only set if changed
	if d.HasChange("max_budget") {
		if v, ok := d.GetOk("max_budget"); ok {
			budget := v.(float64)
			request.MaxBudget = &budget
		}
	}
	if d.HasChange("team_member_budget") {
		if v, ok := d.GetOk("team_member_budget"); ok {
			budget := v.(float64)
			request.TeamMemberBudget = &budget
		}
	}

	// Bool fields - only set if changed
	if d.HasChange("blocked") {
		request.Blocked = d.Get("blocked").(bool)
	}

	// Map fields - only set if changed
	if d.HasChange("metadata") {
		if v, ok := d.GetOk("metadata"); ok {
			request.Metadata = v.(map[string]interface{})
		}
	}

	// String list fields - only set if changed
	if d.HasChange("models") {
		if v, ok := d.GetOk("models"); ok {
			models := make([]string, 0)
			for _, model := range v.([]interface{}) {
				if s, ok := model.(string); ok {
					models = append(models, s)
				}
			}
			request.Models = models
		}
	}

	return request
}

// setTeamResourceData sets Terraform resource data from TeamInfoResponse
func setTeamResourceData(d *schema.ResourceData, teamResp *TeamInfoResponse) error {
	teamInfo := teamResp.TeamInfo

	fields := map[string]interface{}{
		"team_alias":              teamInfo.TeamAlias,
		"organization_id":         teamInfo.OrganizationID,
		"blocked":                 teamInfo.Blocked,
		"team_member_permissions": teamInfo.TeamMemberPermissions,
	}

	// Handle pointer fields
	if teamInfo.TPMLimit != nil {
		fields["tpm_limit"] = *teamInfo.TPMLimit
	}
	if teamInfo.RPMLimit != nil {
		fields["rpm_limit"] = *teamInfo.RPMLimit
	}
	if teamInfo.MaxBudget != nil {
		fields["max_budget"] = *teamInfo.MaxBudget
	}
	if teamInfo.BudgetDuration != nil {
		fields["budget_duration"] = *teamInfo.BudgetDuration
	}

	for field, value := range fields {
		// Use SetIfNotZero to preserve existing values when API doesn't return them
		utils.SetIfNotZero(d, field, value)
	}

	// Handle metadata separately as it's a map
	if teamInfo.Metadata != nil {
		d.Set("metadata", teamInfo.Metadata)
	}

	// Handle models separately as it's a list
	if teamInfo.Models != nil {
		d.Set("models", teamInfo.Models)
	} else {
		d.Set("models", d.Get("models"))
	}

	return nil
}
