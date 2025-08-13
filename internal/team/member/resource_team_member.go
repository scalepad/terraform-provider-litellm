package member

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
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
				Required: true,
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
	log.Printf("[INFO] Creating LiteLLM team member")

	client := m.(*litellm.Client)

	memberData := buildTeamMemberData(d)
	member := buildTeamMemberForCreation(memberData)

	memberResp, err := createTeamMember(ctx, client, member)
	if err != nil {
		return diag.Errorf("error creating team member: %v", err)
	}

	// Set a composite ID since there's no specific member ID returned
	d.SetId(fmt.Sprintf("%s:%s", d.Get("team_id").(string), d.Get("user_id").(string)))

	log.Printf("[INFO] Team member created with ID: %s", d.Id())

	// Set the response data
	if err := setTeamMemberResourceData(d, memberResp); err != nil {
		return diag.Errorf("error setting team member data: %v", err)
	}

	return nil
}

// resourceTeamMemberRead reads the current state of a team member from LiteLLM.
func resourceTeamMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Reading LiteLLM team member with ID: %s", d.Id())

	// The API doesn't provide a way to read specific team members
	// We'll maintain the state as is since the data is already in Terraform state
	// This is common for resources where the API doesn't support individual reads

	log.Printf("[INFO] Successfully read team member with ID: %s", d.Id())
	return nil
}

// resourceTeamMemberUpdate updates an existing team member in LiteLLM.
func resourceTeamMemberUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Updating LiteLLM team member with ID: %s", d.Id())

	client := m.(*litellm.Client)

	memberData := buildTeamMemberData(d)
	member := buildTeamMemberForCreation(memberData)

	memberResp, err := updateTeamMember(ctx, client, member)
	if err != nil {
		return diag.Errorf("error updating team member: %v", err)
	}

	// Set the response data
	if err := setTeamMemberResourceData(d, memberResp); err != nil {
		return diag.Errorf("error setting team member data: %v", err)
	}

	log.Printf("[INFO] Successfully updated team member with ID: %s", d.Id())
	return nil
}

// resourceTeamMemberDelete deletes a team member from LiteLLM.
func resourceTeamMemberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Deleting LiteLLM team member with ID: %s", d.Id())

	client := m.(*litellm.Client)

	teamID := d.Get("team_id").(string)
	userID := d.Get("user_id").(string)
	userEmail := d.Get("user_email").(string)

	if err := deleteTeamMember(ctx, client, teamID, userID, userEmail); err != nil {
		return diag.Errorf("error deleting team member: %v", err)
	}

	log.Printf("[INFO] Successfully deleted team member with ID: %s", d.Id())
	return nil
}
