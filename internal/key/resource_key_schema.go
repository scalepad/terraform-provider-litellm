package key

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceKeySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Core key fields - sensitive and computed
		"key": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "The generated API key. This is only available during creation and is sensitive.",
		},
		"key_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The truncated key name for display purposes.",
		},
		"token": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "The token identifier for the key.",
		},
		"token_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique token ID.",
		},
		"budget_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The budget ID associated with this key.",
		},

		// Configuration fields - user configurable
		"key_alias": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An alias for the key for easier identification.",
		},
		"models": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of models this key can access.",
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
		"user_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The user ID associated with this key.",
		},
		"team_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The team ID associated with this key.",
		},
		"max_parallel_requests": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Maximum number of parallel requests allowed for this key.",
		},
		"metadata": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Additional metadata for the key.",
		},
		"tpm_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Tokens per minute limit for this key.",
		},
		"rpm_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Requests per minute limit for this key.",
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
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of allowed cache control values.",
		},
		"key_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			ForceNew: true,
			ValidateFunc: validation.StringInSlice([]string{
				"llm_api",
				"management",
				"read_only",
				"default",
			}, false),
			Description: "Type of key that determines default allowed routes. Options: 'llm_api' (can call LLM API routes), 'management' (can call management routes), 'read_only' (can only call info/read routes), 'default' (uses default allowed routes). Defaults to 'default'.",
		},

		// Budget and spending fields - some computed, some configurable
		"spend": {
			Type:        schema.TypeFloat,
			Computed:    true,
			Description: "Current spend amount for this key.",
		},
		"max_budget": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Maximum budget allowed for this key.",
		},
		"soft_budget": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Soft budget limit that triggers warnings.",
		},

		// Advanced configuration
		"aliases": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Model aliases for this key.",
		},
		"permissions": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Permissions configuration for the key.",
		},
		"model_max_budget": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeFloat},
			Description: "Per-model budget limits.",
		},
		"model_rpm_limit": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeInt},
			Description: "Per-model RPM limits.",
		},
		"model_tpm_limit": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeInt},
			Description: "Per-model TPM limits.",
		},
		"enforced_params": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Parameters that are enforced for all requests using this key.",
		},

		// Security and control fields
		"guardrails": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of guardrails applied to this key.",
		},
		"prompts": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of prompts associated with this key.",
		},
		"blocked": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether the key is blocked from making requests.",
		},
		"tags": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Tags associated with this key.",
		},
		"send_invite_email": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to send an invite email when creating the key.",
		},

		// Timestamps and audit fields - all computed
		"expires": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Expiration timestamp for the key.",
		},
		"created_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User who created the key.",
		},
		"updated_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User who last updated the key.",
		},
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the key was created.",
		},
		"updated_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the key was last updated.",
		},
	}
}
