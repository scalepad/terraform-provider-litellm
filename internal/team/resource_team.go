package team

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
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
	log.Printf("[INFO] Creating LiteLLM team")

	client := m.(*litellm.Client)

	teamData := buildTeamDataForUtils(d)
	team := buildTeamForCreation(teamData)

	teamResp, err := createTeam(ctx, client, team)
	if err != nil {
		return diag.Errorf("error creating team: %v", err)
	}

	d.SetId(teamResp.TeamID)
	log.Printf("[INFO] Created team with ID: %s", teamResp.TeamID)

	return resourceTeamRead(ctx, d, m)
}

// resourceTeamRead reads the current state of a team from LiteLLM.
func resourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Reading LiteLLM team with ID: %s", d.Id())

	client := m.(*litellm.Client)

	teamResp, err := getTeam(ctx, client, d.Id())
	if err != nil {
		return diag.Errorf("error reading team: %v", err)
	}

	if teamResp == nil {
		log.Printf("[WARN] Team %s not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err := setTeamResourceData(d, teamResp); err != nil {
		return diag.Errorf("error setting team data: %v", err)
	}

	// Get and set permissions
	permResp, err := getTeamPermissions(ctx, client, d.Id())
	if err != nil {
		log.Printf("[WARN] Error getting team permissions: %v", err)
	} else if permResp != nil {
		d.Set("team_member_permissions", permResp.TeamMemberPermissions)
	}

	log.Printf("[INFO] Successfully read team with ID: %s", d.Id())
	return nil
}

// resourceTeamUpdate updates an existing team in LiteLLM.
func resourceTeamUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Updating LiteLLM team with ID: %s", d.Id())

	client := m.(*litellm.Client)

	teamData := buildTeamDataForUtils(d)
	team := buildTeamForCreation(teamData)

	_, err := updateTeam(ctx, client, team)
	if err != nil {
		return diag.Errorf("error updating team: %v", err)
	}

	// Update permissions if they have changed
	if d.HasChange("team_member_permissions") {
		if _, ok := d.GetOk("team_member_permissions"); ok {
			permData := make(map[string]interface{})
			utils.GetStringListValue(d, "team_member_permissions", permData)
			if permissions, exists := permData["team_member_permissions"].([]string); exists {
				if err := updateTeamPermissions(ctx, client, d.Id(), permissions); err != nil {
					return diag.Errorf("error updating team permissions: %v", err)
				}
			}
		}
	}

	log.Printf("[INFO] Successfully updated team with ID: %s", d.Id())
	return resourceTeamRead(ctx, d, m)
}

// resourceTeamDelete deletes a team from LiteLLM.
func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Deleting LiteLLM team with ID: %s", d.Id())

	client := m.(*litellm.Client)

	if err := deleteTeam(ctx, client, d.Id()); err != nil {
		return diag.Errorf("error deleting team: %v", err)
	}

	log.Printf("[INFO] Successfully deleted team with ID: %s", d.Id())
	return nil
}
