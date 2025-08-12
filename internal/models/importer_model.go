package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ModelImporter provides import functionality for LiteLLM model resources
func ModelImporter() *schema.ResourceImporter {
	return &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	}
}
