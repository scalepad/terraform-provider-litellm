package key

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceKeySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:      schema.TypeString,
			Computed:  true,
			Sensitive: true,
		},
		"models": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"spend": {
			Type:     schema.TypeFloat,
			Computed: true,
		},
		"max_budget": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"user_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"team_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"max_parallel_requests": {
			Type:     schema.TypeInt,
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
		"budget_duration": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: validation.StringMatch(
				regexp.MustCompile(`^(\d+[smhd])$`),
				"Budget duration must be in format: number followed by 's' (seconds), 'm' (minutes), 'h' (hours), or 'd' (days). Examples: '30s', '30m', '30h', '30d'",
			),
			Description: "Budget is reset at the end of specified duration. If not set, budget is never reset. You can set duration as seconds ('30s'), minutes ('30m'), hours ('30h'), days ('30d').",
		},
		"allowed_cache_controls": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"soft_budget": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"key_alias": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"duration": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: validation.StringMatch(
				regexp.MustCompile(`^(\d+[smhd])$`),
				"Duration must be in format: number followed by 's' (seconds), 'm' (minutes), 'h' (hours), or 'd' (days). Examples: '30s', '30m', '30h', '30d'",
			),
			Description: "Duration for which this key is valid. You can set duration as seconds ('30s'), minutes ('30m'), hours ('30h'), days ('30d').",
		},
		"aliases": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"config": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"permissions": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"model_max_budget": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeFloat},
		},
		"model_rpm_limit": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"model_tpm_limit": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"guardrails": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"blocked": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"send_invite_email": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"key_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			ValidateFunc: validation.StringInSlice([]string{
				"llm_api",
				"management",
				"read_only",
				"default",
			}, false),
			Description: "Type of key that determines default allowed routes. Options: 'llm_api' (can call LLM API routes), 'management' (can call management routes), 'read_only' (can only call info/read routes), 'default' (uses default allowed routes). Defaults to 'default'.",
		},
	}
}
