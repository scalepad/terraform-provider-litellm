package team

import (
	"context"
	"regexp"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// ResourceTeam defines the schema for the LiteLLM team resource.
func ResourceTeam() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTeamCreate,
		ReadContext:   resourceTeamRead,
		UpdateContext: resourceTeamUpdate,
		DeleteContext: resourceTeamDelete,
		Importer:      TeamImporter(),

		Schema: map[string]*schema.Schema{
			"team_alias": {
				Type:     schema.TypeString,
				Required: true,
			},
			"organization_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tpm_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rpm_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_budget": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"budget_duration": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^(\d+[smhd])$`),
					"Budget duration must be in format: number followed by 's' (seconds), 'm' (minutes), 'h' (hours), or 'd' (days). Examples: '30s', '30m', '30h', '30d'",
				),
				Description: "Budget is reset at the end of specified duration. If not set, budget is never reset. You can set duration as seconds ('30s'), minutes ('30m'), hours ('30h'), days ('30d').",
			},
			"models": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"blocked": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"team_member_permissions": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of permissions granted to team members",
			},
			"team_member_budget": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "Budget automatically given to a new team member",
			},
		},
	}
}

// resourceTeamCreate creates a new team in LiteLLM.
func resourceTeamCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Creating LiteLLM team")

	client := m.(*litellm.Client)

	request := buildTeamCreateRequest(d)

	// Generate UUIDv7 if team_id is not provided
	if request.TeamID == nil {
		teamUUID, err := uuid.NewV7()
		if err != nil {
			return diag.Errorf("failed to generate team ID: %v", err)
		}
		teamID := teamUUID.String()
		request.TeamID = &teamID
	}

	teamResp, err := createTeam(ctx, client, request)
	if err != nil {
		return diag.Errorf("error creating team: %v", err)
	}

	d.SetId(teamResp.TeamID)
	tflog.Info(ctx, "Created team", map[string]interface{}{"team_id": teamResp.TeamID})

	if _, ok := d.GetOk("team_member_permissions"); ok {
		tflog.Info(ctx, "Updating team member permissions", map[string]interface{}{"team_id": d.Id()})

		v := d.Get("team_member_permissions")
		permissionsList := v.([]interface{})
		permissions := make([]string, len(permissionsList))
		for i, perm := range permissionsList {
			permissions[i] = perm.(string)
		}

		if err := updateTeamPermissions(ctx, client, d.Id(), permissions); err != nil {
			return diag.Errorf("error updating team permissions: %v", err)
		}
	}

	return resourceTeamRead(ctx, d, m)
}

// resourceTeamRead reads the current state of a team from LiteLLM.
func resourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Reading LiteLLM team", map[string]interface{}{"team_id": d.Id()})

	client := m.(*litellm.Client)

	teamResp, err := getTeam(ctx, client, d.Id())
	if err != nil {
		return diag.Errorf("error reading team: %v", err)
	}

	if teamResp == nil {
		tflog.Warn(ctx, "Team not found, removing from state", map[string]interface{}{"team_id": d.Id()})
		d.SetId("")
		return nil
	}

	if err := setTeamResourceData(d, teamResp); err != nil {
		return diag.Errorf("error setting team data: %v", err)
	}

	// Get and set permissions
	permResp, err := getTeamPermissions(ctx, client, d.Id())
	if err != nil {
		tflog.Warn(ctx, "Error getting team permissions", map[string]interface{}{"error": err.Error()})
	} else if permResp != nil {
		d.Set("team_member_permissions", permResp.TeamMemberPermissions)
	}

	tflog.Info(ctx, "Successfully read team", map[string]interface{}{"team_id": d.Id()})
	return nil
}

// resourceTeamUpdate updates an existing team in LiteLLM.
func resourceTeamUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Updating LiteLLM team", map[string]interface{}{"team_id": d.Id()})

	client := m.(*litellm.Client)

	request := buildTeamUpdateRequest(d, d.Id())

	_, err := updateTeam(ctx, client, request)
	if err != nil {
		return diag.Errorf("error updating team: %v", err)
	}
	if d.HasChange("team_member_permissions") {
		tflog.Info(ctx, "Updating team member permissions", map[string]interface{}{"team_id": d.Id()})

		v := d.Get("team_member_permissions")
		permissionsList := v.([]interface{})
		permissions := make([]string, len(permissionsList))
		for i, perm := range permissionsList {
			permissions[i] = perm.(string)
		}

		if err := updateTeamPermissions(ctx, client, d.Id(), permissions); err != nil {
			return diag.Errorf("error updating team permissions: %v", err)
		}
	}
	tflog.Info(ctx, "Successfully updated team", map[string]interface{}{"team_id": d.Id()})
	return resourceTeamRead(ctx, d, m)
}

// resourceTeamDelete deletes a team from LiteLLM.
func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Deleting LiteLLM team", map[string]interface{}{"team_id": d.Id()})

	client := m.(*litellm.Client)

	if err := deleteTeam(ctx, client, d.Id()); err != nil {
		return diag.Errorf("error deleting team: %v", err)
	}

	tflog.Info(ctx, "Successfully deleted team", map[string]interface{}{"team_id": d.Id()})
	return nil
}
