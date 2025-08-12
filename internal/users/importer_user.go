package users

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// UserImporter returns the importer for the user resource
func UserImporter() *schema.ResourceImporter {
	return &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	}
}
