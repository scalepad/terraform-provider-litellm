package member

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// ResourceTeamMemberAdd defines the schema for the LiteLLM team member add resource.
func ResourceTeamMemberAdd() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTeamMemberAddCreate,
		ReadContext:   resourceTeamMemberAddRead,
		UpdateContext: resourceTeamMemberAddUpdate,
		DeleteContext: resourceTeamMemberAddDelete,

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"member": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:     schema.TypeString,
							Optional: true,
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
					},
				},
			},
			"max_budget_in_team": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

// resourceTeamMemberAddCreate creates multiple team members in bulk.
func resourceTeamMemberAddCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Creating LiteLLM team members")

	client := m.(*litellm.Client)

	memberData := buildTeamMemberAddData(d)
	memberAdd := buildTeamMemberAddForCreation(memberData)

	if err := createTeamMembersBulk(ctx, client, memberAdd); err != nil {
		return diag.Errorf("error adding team members: %v", err)
	}

	// Set ID as team_id since this resource manages all members for a team
	d.SetId(memberAdd.TeamID)

	log.Printf("[INFO] Team members created for team ID: %s", d.Id())
	return nil
}

// resourceTeamMemberAddRead reads the current state of team members.
func resourceTeamMemberAddRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Reading LiteLLM team members for team ID: %s", d.Id())

	// The API doesn't provide a way to read specific team members
	// We'll maintain the state as is
	log.Printf("[INFO] Successfully read team members for team ID: %s", d.Id())
	return nil
}

// resourceTeamMemberAddUpdate updates team members in bulk.
func resourceTeamMemberAddUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Updating LiteLLM team members for team ID: %s", d.Id())

	client := m.(*litellm.Client)
	teamID := d.Get("team_id").(string)
	maxBudget := d.Get("max_budget_in_team").(float64)

	o, n := d.GetChange("member")
	oldMembers := o.(*schema.Set)
	newMembers := n.(*schema.Set)

	// Create maps for easier lookup by user identifier
	oldMemberMap := make(map[string]map[string]interface{})
	newMemberMap := make(map[string]map[string]interface{})

	// Build old member map using user_id or user_email as key
	for _, member := range oldMembers.List() {
		m := member.(map[string]interface{})
		key := getMemberKey(m)
		if key != "" {
			oldMemberMap[key] = m
		}
	}

	// Build new member map using user_id or user_email as key
	for _, member := range newMembers.List() {
		m := member.(map[string]interface{})
		key := getMemberKey(m)
		if key != "" {
			newMemberMap[key] = m
		}
	}

	// Track which members have been updated to avoid duplicates
	updatedMembers := make(map[string]bool)

	// Check if max_budget_in_team has changed
	if d.HasChange("max_budget_in_team") {
		log.Printf("[DEBUG] max_budget_in_team changed, updating all existing members with new budget: %f", maxBudget)

		// Update ALL existing members with the new budget
		for key, newMember := range newMemberMap {
			if _, exists := oldMemberMap[key]; exists {
				userID, _ := newMember["user_id"].(string)
				userEmail, _ := newMember["user_email"].(string)
				role, _ := newMember["role"].(string)

				if err := updateTeamMemberBudget(ctx, client, teamID, userID, userEmail, role, maxBudget); err != nil {
					return diag.Errorf("error updating team member budget: %v", err)
				}

				// Mark this member as updated
				updatedMembers[key] = true
			}
		}
	}

	// Find members to delete (in old but not in new)
	for key, oldMember := range oldMemberMap {
		if _, exists := newMemberMap[key]; !exists {
			userID, _ := oldMember["user_id"].(string)
			userEmail, _ := oldMember["user_email"].(string)

			if err := deleteTeamMember(ctx, client, teamID, userID, userEmail); err != nil {
				return diag.Errorf("error deleting team member: %v", err)
			}
		}
	}

	// Find members to update (exist in both but with different attributes)
	// Skip members that were already updated due to budget change
	for key, newMember := range newMemberMap {
		if oldMember, exists := oldMemberMap[key]; exists {
			// Skip if already updated due to budget change
			if updatedMembers[key] {
				continue
			}

			// Check if member attributes have changed
			if memberAttributesChanged(oldMember, newMember) {
				userID, _ := newMember["user_id"].(string)
				userEmail, _ := newMember["user_email"].(string)
				role, _ := newMember["role"].(string)

				if err := updateTeamMemberRole(ctx, client, teamID, userID, userEmail, role, maxBudget); err != nil {
					return diag.Errorf("error updating team member: %v", err)
				}
			}
		}
	}

	// Find members to add (in new but not in old)
	var membersToAdd []TeamMemberAddMember
	for key, newMember := range newMemberMap {
		if _, exists := oldMemberMap[key]; !exists {
			member := TeamMemberAddMember{}
			if userID, ok := newMember["user_id"].(string); ok && userID != "" {
				member.UserID = userID
			}
			if userEmail, ok := newMember["user_email"].(string); ok && userEmail != "" {
				member.UserEmail = userEmail
			}
			if role, ok := newMember["role"].(string); ok {
				member.Role = role
			}
			membersToAdd = append(membersToAdd, member)
		}
	}

	if len(membersToAdd) > 0 {
		memberAdd := &TeamMemberAdd{
			TeamID:          teamID,
			Members:         membersToAdd,
			MaxBudgetInTeam: maxBudget,
		}

		if err := createTeamMembersBulk(ctx, client, memberAdd); err != nil {
			return diag.Errorf("error adding team members: %v", err)
		}
	}

	log.Printf("[INFO] Successfully updated team members for team ID: %s", d.Id())
	return nil
}

// resourceTeamMemberAddDelete deletes all team members for a team.
func resourceTeamMemberAddDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] Deleting LiteLLM team members for team ID: %s", d.Id())

	client := m.(*litellm.Client)
	teamID := d.Get("team_id").(string)
	members := d.Get("member").(*schema.Set)

	// Delete each member
	for _, member := range members.List() {
		m := member.(map[string]interface{})
		userID, _ := m["user_id"].(string)
		userEmail, _ := m["user_email"].(string)

		if err := deleteTeamMember(ctx, client, teamID, userID, userEmail); err != nil {
			return diag.Errorf("error deleting team member: %v", err)
		}
	}

	log.Printf("[INFO] Successfully deleted team members for team ID: %s", d.Id())
	return nil
}
