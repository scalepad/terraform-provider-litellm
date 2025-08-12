package users

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// resourceUserSchema returns the schema for the user resource
func resourceUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "Unique identifier for the user. If not set, a unique id will be generated.",
		},
		"user_email": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Email address of the user",
		},
		"user_alias": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Alias or display name for the user",
		},
		"user_role": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				"proxy_admin",
				"proxy_admin_viewer",
				"internal_user",
				"internal_user_viewer",
			}, false),
			Description: "Role assigned to the user. Valid values: proxy_admin, proxy_admin_viewer, internal_user, internal_user_viewer",
		},
		"max_budget": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Maximum budget allowed for the user",
		},
		"budget_duration": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Duration for the budget (e.g., '30s', '30m', '30h', '30d', '1mo')",
		},
		"models": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of models the user has access to. Set to ['no-default-models'] to block all model access.",
		},
		"tpm_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Tokens per minute limit for the user",
		},
		"rpm_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Requests per minute limit for the user",
		},
		"metadata": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Additional metadata for the user",
		},
		"send_invite_email": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to send an invite email to the user",
		},
		"auto_create_key": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether to automatically create a key for the user",
		},
		"aliases": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Model aliases for the user",
		},
		"config": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "User-specific configuration (deprecated)",
		},
		"allowed_cache_controls": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of allowed cache control values (e.g., ['no-cache', 'no-store'])",
		},
		"blocked": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether the user is blocked",
		},
		"guardrails": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of active guardrails for the user",
		},
		"permissions": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "User-specific permissions",
		},
		"max_parallel_requests": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Maximum number of parallel requests allowed for the user",
		},
		"soft_budget": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Soft budget limit that triggers alerts but doesn't block requests",
		},
		"model_max_budget": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Model-specific maximum budget for the user",
		},
		"model_rpm_limit": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Model-specific RPM limit for the user",
		},
		"model_tpm_limit": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Model-specific TPM limit for the user",
		},
		"duration": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Duration for the auto-created key",
		},
		"key_alias": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Alias for the auto-created key",
		},
		"sso_user_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SSO user identifier",
		},
		"object_permission": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Object-specific permissions for the user",
		},
		"prompts": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of allowed prompts for the user",
		},
		"organizations": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of organization IDs the user belongs to",
		},
		// Computed fields
		"spend": {
			Type:        schema.TypeFloat,
			Computed:    true,
			Description: "Current spend amount for the user",
		},
		"key_count": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of keys associated with the user",
		},
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the user was created",
		},
		"updated_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the user was last updated",
		},
		"budget_reset_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the budget will be reset",
		},
	}
}
