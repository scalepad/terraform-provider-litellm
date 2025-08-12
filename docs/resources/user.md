# litellm_user Resource

Manages a LiteLLM user with comprehensive configuration options including budget controls, rate limiting, model access, permissions, and advanced features.

## Example Usage

### Basic User

```hcl
resource "litellm_user" "basic" {
  user_email = "user@example.com"
  user_alias = "John Doe"
  user_role  = "internal_user"
}
```

### Comprehensive User Configuration

```hcl
resource "litellm_user" "advanced" {
  user_id            = "custom-user-123"  # Optional - auto-generates UUIDv7 if not provided
  user_email         = "admin@example.com"
  user_alias         = "Admin User"
  user_role          = "proxy_admin"

  # Budget and limits
  max_budget         = 1000.0
  soft_budget        = 800.0
  budget_duration    = "1mo"
  tpm_limit          = 50000
  rpm_limit          = 500
  max_parallel_requests = 10

  # Model access and aliases
  models = ["gpt-4", "claude-3.5-sonnet", "gpt-3.5-turbo"]
  aliases = {
    "my-gpt4" = "gpt-4"
    "my-claude" = "claude-3.5-sonnet"
  }

  # Model-specific limits
  model_max_budget = {
    "gpt-4" = "500"
    "claude-3.5-sonnet" = "300"
  }
  model_tpm_limit = {
    "gpt-4" = "10000"
    "claude-3.5-sonnet" = "15000"
  }
  model_rpm_limit = {
    "gpt-4" = "100"
    "claude-3.5-sonnet" = "150"
  }

  # Security and permissions
  permissions = {
    "admin_ui_access" = "true"
    "model_management" = "true"
  }
  object_permission = {
    "team_123" = "read"
    "project_456" = "write"
  }
  guardrails = ["content_filter", "pii_detection"]

  # User management
  send_invite_email = true
  auto_create_key   = true
  key_alias         = "admin-key"
  duration          = "30d"
  blocked           = false

  # Advanced features
  allowed_cache_controls = ["no-cache", "no-store"]
  prompts = ["system_prompt_1", "user_prompt_template"]
  organizations = ["org_123", "org_456"]
  sso_user_id = "sso_external_id_789"

  # Metadata
  metadata = {
    department = "engineering"
    team       = "ai-platform"
    cost_center = "R&D"
    environment = "production"
  }

  # Deprecated but supported
  config = {
    "legacy_setting" = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

### Core User Information

- `user_id` - (Optional) Unique identifier for the user. If not provided, a UUIDv7 will be automatically generated. This cannot be changed after creation.

- `user_email` - (Optional) Email address of the user. Used for notifications and user identification.

- `user_alias` - (Optional) Alias or display name for the user. Provides a human-readable name.

- `user_role` - (Optional) Role assigned to the user. Valid values:
  - `proxy_admin` - Full administrative access to the proxy
  - `proxy_admin_viewer` - Read-only administrative access to the proxy
  - `internal_user` - Internal user with standard access
  - `internal_user_viewer` - Internal user with read-only access

### Budget and Rate Limiting

- `max_budget` - (Optional) Maximum budget allowed for the user. Sets an upper limit on total spend.

- `soft_budget` - (Optional) Soft budget limit that triggers alerts but doesn't block requests.

- `budget_duration` - (Optional) Duration for the budget (e.g., '30s', '30m', '30h', '30d', '1mo'). Defines the time period for budget limits.

- `tpm_limit` - (Optional) Tokens per minute limit for the user. Rate limit based on token processing.

- `rpm_limit` - (Optional) Requests per minute limit for the user. Rate limit based on API calls.

- `max_parallel_requests` - (Optional) Maximum number of parallel requests allowed for the user.

### Model Access and Configuration

- `models` - (Optional) List of models the user has access to. Set to `['no-default-models']` to block all model access.

- `aliases` - (Optional) Model aliases for the user. Map of alias names to actual model names.

- `model_max_budget` - (Optional) Model-specific maximum budget limits. Map of model names to budget strings.

- `model_tpm_limit` - (Optional) Model-specific TPM limits. Map of model names to limit strings.

- `model_rpm_limit` - (Optional) Model-specific RPM limits. Map of model names to limit strings.

### Security and Permissions

- `permissions` - (Optional) User-specific permissions. Map of permission names to values.

- `object_permission` - (Optional) Object-specific permissions for the user. Map of object IDs to permission levels.

- `guardrails` - (Optional) List of active guardrails for the user (e.g., content filters, PII detection).

- `blocked` - (Optional) Whether the user is blocked. Defaults to `false`.

### User Management

- `send_invite_email` - (Optional) Whether to send an invite email to the user. Defaults to `false`.

- `auto_create_key` - (Optional) Whether to automatically create a key for the user. Defaults to `true`.

- `key_alias` - (Optional) Alias for the auto-created key.

- `duration` - (Optional) Duration for the auto-created key.

- `sso_user_id` - (Optional) SSO user identifier for external authentication systems.

### Advanced Features

- `allowed_cache_controls` - (Optional) List of allowed cache control values (e.g., `['no-cache', 'no-store']`).

- `prompts` - (Optional) List of allowed prompts for the user.

- `organizations` - (Optional) List of organization IDs the user belongs to.

### Metadata and Configuration

- `metadata` - (Optional) Additional metadata for the user. Can store custom information like department, team, or organizational data.

- `config` - (Optional) User-specific configuration. **Deprecated** - use other specific fields instead.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

### Computed Fields

- `spend` - Current spend amount for the user. Reflects total amount spent across all API usage.

- `key_count` - Number of keys associated with the user. Shows how many API keys are linked to this user.

- `created_at` - Timestamp when the user was created (RFC3339 format).

- `updated_at` - Timestamp when the user was last updated (RFC3339 format).

- `budget_reset_at` - Timestamp when the budget will be reset (RFC3339 format).

## State Management

The User resource efficiently manages its state by preserving all configured values in the Terraform state file. The provider uses utility functions to ensure that non-zero and non-empty values are correctly maintained, preventing unnecessary updates and ensuring consistency between your configuration and the actual resource state.

## Import

LiteLLM users can be imported using the `user_id` in two ways:

### Using terraform import command

```
$ terraform import litellm_user.example user123
```

### Using import block (Terraform 1.5+)

```hcl
import {
  to = litellm_user.example
  id = "user123"
}

resource "litellm_user" "example" {
  user_id    = "user123"
  user_email = "imported@example.com"
  # ... other configuration
}
```

This allows you to import existing users into your Terraform state, enabling management of users that were created outside of Terraform.

## API Endpoints

This resource interacts with the following LiteLLM API endpoints:

- **Create**: `POST /user/new` - Creates a new user with the specified configuration
- **Read**: `GET /user/info?user_id={user_id}` - Retrieves current user information including spend and key count
- **Update**: `POST /user/update` - Updates user configuration (supports most user fields)
- **Delete**: `POST /user/delete` - Removes the user from the system

Note that the delete operation uses a POST request with the user ID in the request body, as required by the LiteLLM API.

## Notes

- When `user_id` is not provided, the system automatically generates a UUIDv7 identifier
- The `auto_create_key` feature automatically creates an API key for the user when enabled
- Model-specific limits override global limits for those specific models
- Soft budgets trigger alerts but don't block requests, while hard budgets (`max_budget`) do block requests
- The `config` field is deprecated in favor of more specific configuration fields
