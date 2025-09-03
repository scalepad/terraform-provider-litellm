# litellm_service_account Resource

Manages a LiteLLM service account key for production projects.

## Overview

Service account keys are designed for production use cases where you need API keys that are not tied to individual users. They provide better security and management for automated systems, microservices, and production applications.

### Why use service account keys?

- **User Independence**: Keys are not deleted when users are removed from the system
- **Team-Level Limits**: Apply team limits rather than individual user limits
- **Better Separation**: Clear distinction between user keys and service keys
- **Production Ready**: Ideal for automated systems and production deployments

## Example Usage

```hcl
resource "litellm_service_account" "production_api" {
  team_id          = "01991103-7931-701f-8a87-d56fb671593e"
  key_alias        = "production-api-service"
  models           = ["all-team-models"]
  max_budget       = 1000.0
  tpm_limit        = 50000
  rpm_limit        = 500
  budget_duration  = "30d"
  metadata         = {
    environment = "production"
    service     = "api-gateway"
    owner       = "platform-team"
  }
  service_account_id = "prod-api-gateway-001"  # Optional - auto-generated if not provided
}

# Service account with minimal configuration
resource "litellm_service_account" "simple_service" {
  team_id   = "01991103-7931-701f-8a87-d56fb671593e"
  key_alias = "simple-service"
  models    = ["gpt-4", "claude-3.5-sonnet"]
}
```

## Argument Reference

The following arguments are supported:

### Required

- `team_id` - (Required) The team ID the service account belongs to. This associates the service account with a specific team and its permissions.

### Optional

- `key_alias` - (Optional) A human-readable alias for the service account. This helps identify the service account in logs and management interfaces.

- `models` - (Optional) List of models that this service account can access. Use `["all-team-models"]` to allow access to all models available to the team.

- `key_type` - (Optional) Type of key that determines default allowed routes. Options: `"llm_api"` (can call LLM API routes), `"management"` (can call management routes), `"read_only"` (can only call info/read routes), `"default"` (uses default allowed routes). Defaults to `"default"`.

- `metadata` - (Optional) Additional metadata for the service account. This can include custom information like environment, service name, owner, etc.

- `service_account_id` - (Optional) Unique identifier for the service account. If not provided, a UUIDv7 will be automatically generated.

### Budget and Rate Limiting

- `max_budget` - (Optional) Maximum budget for this service account. Sets an upper limit on total spend.

- `soft_budget` - (Optional) Soft budget limit that triggers warnings before reaching the hard limit.

- `budget_duration` - (Optional) Budget is reset at the end of specified duration. If not set, budget is never reset. You can set duration as seconds ("30s"), minutes ("30m"), hours ("30h"), days ("30d").

- `tpm_limit` - (Optional) Tokens per minute limit for rate limiting.

- `rpm_limit` - (Optional) Requests per minute limit for rate limiting.

- `max_parallel_requests` - (Optional) Maximum number of parallel requests allowed.

### Advanced Configuration

- `aliases` - (Optional) Model aliases for custom model naming.

- `permissions` - (Optional) Key-specific permissions configuration.

- `model_max_budget` - (Optional) Per-model budget limits as a map.

- `model_rpm_limit` - (Optional) Per-model RPM limits as a map.

- `model_tpm_limit` - (Optional) Per-model TPM limits as a map.

- `enforced_params` - (Optional) Parameters that are enforced for all requests.

- `guardrails` - (Optional) List of guardrails applied to this service account.

- `prompts` - (Optional) List of prompts associated with this service account.

- `allowed_cache_controls` - (Optional) List of allowed cache control directives.

- `tags` - (Optional) List of tags for organization and filtering.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `key` - The generated API key value (sensitive).

- `key_name` - The truncated key name for display.

- `token` - The token identifier (sensitive).

- `token_id` - The unique token ID.

- `spend` - Current spend amount for this service account.

- `created_at` - Timestamp when the service account was created.

- `updated_at` - Timestamp when the service account was last updated.

- `created_by` - User who created the service account.

- `updated_by` - User who last updated the service account.

## Service Account ID

The `service_account_id` is a unique identifier that distinguishes service accounts from regular user keys. If you don't provide one, the system will automatically generate a UUIDv7. This ID is stored in the key's metadata and can be used for:

- Tracking and auditing service account usage
- Managing service accounts programmatically
- Integration with external systems

## Import

Service accounts can be imported using the token ID:

```bash
terraform import litellm_service_account.example <token_id>
```

## Best Practices

1. **Use descriptive aliases**: Choose meaningful names that indicate the service account's purpose
2. **Set appropriate limits**: Configure budget and rate limits based on your service requirements
3. **Use metadata**: Include relevant metadata for tracking and management
4. **Team association**: Always associate service accounts with appropriate teams for proper access control
5. **Monitor usage**: Regularly review spend and usage patterns

## Differences from Regular Keys

Service account keys differ from regular `litellm_key` resources in the following ways:

- **No user association**: Not tied to individual users
- **Team-level limits**: Subject to team limits rather than user limits
- **Persistence**: Not affected by user lifecycle changes
- **Automatic ID generation**: Service account ID is automatically managed
- **Production focus**: Designed specifically for production and automated use cases
