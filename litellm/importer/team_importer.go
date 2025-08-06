package importer

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TeamImporter provides import functionality for LiteLLM team resources
func TeamImporter() *schema.ResourceImporter {
	return &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	}
}
