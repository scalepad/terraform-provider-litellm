package member

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TeamMemberImporter provides import functionality for LiteLLM team member resources
func TeamMemberImporter() *schema.ResourceImporter {
	return &schema.ResourceImporter{
		StateContext: teamMemberImportState,
	}
}

func teamMemberImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Expected format: "team_id:user_id"
	importID := d.Id()

	// Split the import ID on ":"
	parts := strings.Split(importID, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import ID format. Expected 'team_id:user_id', got: %s", importID)
	}

	teamID := parts[0]
	userID := parts[1]

	// Validate that both parts are non-empty
	if teamID == "" || userID == "" {
		return nil, fmt.Errorf("invalid import ID format. Both team_id and user_id must be non-empty. Got: %s", importID)
	}

	// Set the resource ID to the full composite ID
	d.SetId(importID)

	// Set the individual fields that can be derived from the import ID
	d.Set("team_id", teamID)
	d.Set("user_id", userID)

	return []*schema.ResourceData{d}, nil
}
