package serviceaccount

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceServiceAccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Core service account fields - sensitive and computed
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
			Description: "The token identifier for the service account.",
		},
		"token_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique token ID.",
		},

		// Required configuration fields
		"team_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The team ID the service account belongs to. This is required.",
		},
		"key_alias": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An alias for the service account for easier identification.",
		},
		"models": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of models this service account can access.",
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
		"metadata": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Additional metadata for the service account.",
		},
		"service_account_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service account ID. If not provided, a UUIDv7 will be generated automatically.",
		},

		// Budget and spending fields
		"spend": {
			Type:        schema.TypeFloat,
			Computed:    true,
			Description: "Current spend amount for this service account.",
		},
		"max_budget": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Maximum budget allowed for this service account.",
		},
		"soft_budget": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Soft budget limit that triggers warnings.",
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

		// Rate limiting fields
		"tpm_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Tokens per minute limit for this service account.",
		},
		"rpm_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Requests per minute limit for this service account.",
		},
		"max_parallel_requests": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Maximum number of parallel requests allowed for this service account.",
		},

		// Security and control fields
		"blocked": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether the service account is blocked from making requests.",
		},
		"guardrails": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of guardrails applied to this service account.",
		},
		"prompts": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of prompts associated with this service account.",
		},
		"allowed_cache_controls": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of allowed cache control values.",
		},

		// Advanced configuration
		"aliases": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Model aliases for this service account.",
		},
		"permissions": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Permissions configuration for the service account.",
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
			Description: "Parameters that are enforced for all requests using this service account.",
		},
		"tags": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Tags associated with this service account.",
		},

		// Timestamps and audit fields - all computed
		"expires": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Expiration timestamp for the service account.",
		},
		"created_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User who created the service account.",
		},
		"updated_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User who last updated the service account.",
		},
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the service account was created.",
		},
		"updated_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the service account was last updated.",
		},
	}
}
