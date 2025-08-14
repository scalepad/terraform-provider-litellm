package member

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"github.com/scalepad/terraform-provider-litellm/internal/team"
)

// ResourceTeamMember defines the schema for the LiteLLM team member resource.
func ResourceTeamMember() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTeamMemberCreate,
		ReadContext:   resourceTeamMemberRead,
		UpdateContext: resourceTeamMemberUpdate,
		DeleteContext: resourceTeamMemberDelete,
		Importer:      TeamMemberImporter(),

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"admin",
					"user",
				}, false),
			},
			"max_budget_in_team": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

// resourceTeamMemberCreate creates a new team member in LiteLLM.
func resourceTeamMemberCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Creating LiteLLM team member")

	client := m.(*litellm.Client)

	// Build the create request
	member := TeamMemberCreateMember{
		UserID: d.Get("user_id").(string),
		Role:   d.Get("role").(string),
	}

	// Only set user_email if provided
	if userEmail, ok := d.GetOk("user_email"); ok {
		email := userEmail.(string)
		member.UserEmail = &email
	}

	request := &TeamMemberCreateRequest{
		TeamID:          d.Get("team_id").(string),
		Member:          member,
		MaxBudgetInTeam: d.Get("max_budget_in_team").(float64),
	}

	memberResp, err := createTeamMember(ctx, client, request)
	if err != nil {
		return diag.Errorf("error creating team member: %v", err)
	}

	// Set a composite ID since there's no specific member ID returned
	d.SetId(fmt.Sprintf("%s:%s", d.Get("team_id").(string), d.Get("user_id").(string)))

	tflog.Info(ctx, "Team member created", map[string]interface{}{
		"id": d.Id(),
	})

	// Set the response data
	if err := setTeamMemberResourceData(d, memberResp); err != nil {
		return diag.Errorf("error setting team member data: %v", err)
	}

	return nil
}

// resourceTeamMemberRead reads the current state of a team member from LiteLLM.
func resourceTeamMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Reading LiteLLM team member", map[string]interface{}{
		"id": d.Id(),
	})

	client := m.(*litellm.Client)

	// Parse the composite ID to get team_id and user_id
	parts := strings.Split(d.Id(), ":")
	if len(parts) != 2 {
		return diag.Errorf("invalid team member ID format: %s", d.Id())
	}

	teamID := parts[0]
	userID := parts[1]

	// Get team information which includes team memberships
	teamInfo, err := team.GetTeam(ctx, client, teamID)
	if err != nil {
		return diag.Errorf("error getting team info: %v", err)
	}

	if teamInfo == nil {
		tflog.Warn(ctx, "Team not found, removing team member from state", map[string]interface{}{
			"team_id": teamID,
		})
		d.SetId("")
		return nil
	}

	// Find the specific team member in the team memberships
	var foundMember *team.TeamMembership
	var memberRole string
	var memberEmail string

	// First check team memberships for budget information
	for _, membership := range teamInfo.TeamMemberships {
		if membership.UserID == userID {
			foundMember = &membership
			break
		}
	}

	// Then check members_with_roles for role and email information
	for _, member := range teamInfo.TeamInfo.MembersWithRoles {
		if member.UserID == userID {
			memberRole = member.Role
			if member.UserEmail != nil {
				memberEmail = *member.UserEmail
			}
			break
		}
	}

	// If member is not found in either list, remove from state
	if foundMember == nil || memberRole == "" {
		tflog.Warn(ctx, "Team member not found in team, removing from state", map[string]interface{}{
			"user_id": userID,
			"team_id": teamID,
		})
		d.SetId("")
		return nil
	}

	// Update the resource data with the found information
	d.Set("team_id", teamID)
	d.Set("user_id", userID)

	if memberEmail != "" {
		d.Set("user_email", memberEmail)
	}

	if memberRole != "" {
		d.Set("role", memberRole)
	}

	// Set budget information if available from team membership
	if foundMember != nil {
		if foundMember.LitellmBudgetTable.MaxBudget != nil {
			d.Set("max_budget_in_team", *foundMember.LitellmBudgetTable.MaxBudget)
		}
	}

	tflog.Info(ctx, "Successfully read team member", map[string]interface{}{
		"id": d.Id(),
	})
	return nil
}

// resourceTeamMemberUpdate updates an existing team member in LiteLLM.
func resourceTeamMemberUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Updating LiteLLM team member", map[string]interface{}{
		"id": d.Id(),
	})

	client := m.(*litellm.Client)

	// Build the update request with only changed fields
	request := &TeamMemberUpdateRequest{
		UserID: d.Get("user_id").(string), // Always required for identification
		TeamID: d.Get("team_id").(string), // Always required for identification
	}

	// Only set fields that have changed
	if d.HasChange("user_email") {
		email := d.Get("user_email").(string)
		request.UserEmail = &email
	}

	if d.HasChange("role") {
		role := d.Get("role").(string)
		request.Role = &role
	}

	if d.HasChange("max_budget_in_team") {
		budget := d.Get("max_budget_in_team").(float64)
		request.MaxBudgetInTeam = &budget
	}

	memberResp, err := updateTeamMember(ctx, client, request)
	if err != nil {
		return diag.Errorf("error updating team member: %v", err)
	}

	// Set the response data
	if err := setTeamMemberResourceData(d, memberResp); err != nil {
		return diag.Errorf("error setting team member data: %v", err)
	}

	tflog.Info(ctx, "Successfully updated team member", map[string]interface{}{
		"id": d.Id(),
	})
	return nil
}

// resourceTeamMemberDelete deletes a team member from LiteLLM.
func resourceTeamMemberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Deleting LiteLLM team member", map[string]interface{}{
		"id": d.Id(),
	})

	client := m.(*litellm.Client)

	teamID := d.Get("team_id").(string)
	userID := d.Get("user_id").(string)
	userEmail := d.Get("user_email").(string)

	if err := deleteTeamMember(ctx, client, teamID, userID, userEmail); err != nil {
		return diag.Errorf("error deleting team member: %v", err)
	}

	tflog.Info(ctx, "Successfully deleted team member", map[string]interface{}{
		"id": d.Id(),
	})
	return nil
}
