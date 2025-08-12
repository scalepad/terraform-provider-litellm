package member

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

// buildTeamMemberData builds team member data using utils methods
func buildTeamMemberData(d *schema.ResourceData) map[string]interface{} {
	memberData := make(map[string]interface{})

	// String fields
	utils.GetValueDefault[string](d, "team_id", memberData)
	utils.GetValueDefault[string](d, "user_id", memberData)
	utils.GetValueDefault[string](d, "user_email", memberData)
	utils.GetValueDefault[string](d, "role", memberData)

	// Float64 fields
	utils.GetValueDefault[float64](d, "max_budget_in_team", memberData)

	return memberData
}

// setTeamMemberResourceData sets team member resource data using utils methods
func setTeamMemberResourceData(d *schema.ResourceData, member *TeamMemberResponse) error {
	fields := map[string]interface{}{
		"team_id":            member.TeamID,
		"user_id":            member.UserID,
		"user_email":         member.UserEmail,
		"role":               member.Role,
		"max_budget_in_team": member.MaxBudgetInTeam,
	}

	for field, value := range fields {
		utils.SetIfNotZero(d, field, value)
	}

	return nil
}

// buildTeamMemberForCreation converts map data to TeamMember struct
func buildTeamMemberForCreation(data map[string]interface{}) *TeamMember {
	member := &TeamMember{}

	if v, ok := data["team_id"].(string); ok {
		member.TeamID = v
	}
	if v, ok := data["user_id"].(string); ok {
		member.UserID = v
	}
	if v, ok := data["user_email"].(string); ok {
		member.UserEmail = v
	}
	if v, ok := data["role"].(string); ok {
		member.Role = v
	}
	if v, ok := data["max_budget_in_team"].(float64); ok {
		member.MaxBudgetInTeam = v
	}

	return member
}

// parseTeamMemberAPIResponse parses API response into TeamMemberResponse
func parseTeamMemberAPIResponse(resp map[string]interface{}) (*TeamMemberResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	memberResp := &TeamMemberResponse{}

	if v, ok := resp["team_id"].(string); ok {
		memberResp.TeamID = v
	}
	if v, ok := resp["user_id"].(string); ok {
		memberResp.UserID = v
	}
	if v, ok := resp["user_email"].(string); ok {
		memberResp.UserEmail = v
	}
	if v, ok := resp["role"].(string); ok {
		memberResp.Role = v
	}
	if v, ok := resp["max_budget_in_team"].(float64); ok {
		memberResp.MaxBudgetInTeam = v
	}
	if v, ok := resp["status"].(string); ok {
		memberResp.Status = v
	}

	return memberResp, nil
}

// buildTeamMemberAddData builds data for bulk team member operations
func buildTeamMemberAddData(d *schema.ResourceData) map[string]interface{} {
	memberData := make(map[string]interface{})

	// String fields
	utils.GetValueDefault[string](d, "team_id", memberData)

	// Float64 fields
	utils.GetValueDefault[float64](d, "max_budget_in_team", memberData)

	// Handle member set
	if v, ok := d.GetOk("member"); ok {
		members := v.(*schema.Set)
		membersList := make([]map[string]interface{}, 0, members.Len())

		for _, member := range members.List() {
			m := member.(map[string]interface{})
			memberItem := make(map[string]interface{})

			if role, ok := m["role"].(string); ok && role != "" {
				memberItem["role"] = role
			}
			if userID, ok := m["user_id"].(string); ok && userID != "" {
				memberItem["user_id"] = userID
			}
			if userEmail, ok := m["user_email"].(string); ok && userEmail != "" {
				memberItem["user_email"] = userEmail
			}

			membersList = append(membersList, memberItem)
		}

		memberData["member"] = membersList
	}

	return memberData
}

// buildTeamMemberAddForCreation converts map data to TeamMemberAdd struct
func buildTeamMemberAddForCreation(data map[string]interface{}) *TeamMemberAdd {
	memberAdd := &TeamMemberAdd{}

	if v, ok := data["team_id"].(string); ok {
		memberAdd.TeamID = v
	}
	if v, ok := data["max_budget_in_team"].(float64); ok {
		memberAdd.MaxBudgetInTeam = v
	}

	if members, ok := data["member"].([]map[string]interface{}); ok {
		memberAdd.Members = make([]TeamMemberAddMember, len(members))
		for i, member := range members {
			if userID, ok := member["user_id"].(string); ok {
				memberAdd.Members[i].UserID = userID
			}
			if userEmail, ok := member["user_email"].(string); ok {
				memberAdd.Members[i].UserEmail = userEmail
			}
			if role, ok := member["role"].(string); ok {
				memberAdd.Members[i].Role = role
			}
		}
	}

	return memberAdd
}

// getMemberKey returns a unique key for a member based on user_id or user_email
func getMemberKey(member map[string]interface{}) string {
	if userID, ok := member["user_id"].(string); ok && userID != "" {
		return "id:" + userID
	}
	if userEmail, ok := member["user_email"].(string); ok && userEmail != "" {
		return "email:" + userEmail
	}
	return ""
}

// memberAttributesChanged checks if member attributes have changed between old and new
func memberAttributesChanged(oldMember, newMember map[string]interface{}) bool {
	// Compare role
	oldRole, _ := oldMember["role"].(string)
	newRole, _ := newMember["role"].(string)
	return oldRole != newRole
}
