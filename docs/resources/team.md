# litellm_team Resource

Manages a team configuration in LiteLLM. Teams allow you to group users and manage their access to models and usage limits.

## Example Usage

```hcl
resource "litellm_team" "engineering" {
  team_alias      = "engineering-team"
  organization_id = "org_123456"
  models          = ["gpt-4-proxy", "claude-2"]

  team_member_permissions = [
    "/key/generate",
    "/key/update",
    "/key/info",
    "/key/list"
  ]

  metadata = {
    department = "Engineering"
    project    = "AI Research"
  }

  blocked             = false
  tpm_limit           = 500000
  rpm_limit           = 5000
  max_budget          = 1000.0
  budget_duration     = "30d"
  team_member_budget  = 100.0
}
```

## Argument Reference

The following arguments are supported:

- `team_alias` - (Required) A human-readable identifier for the team.

- `organization_id` - (Optional) The ID of the organization this team belongs to.

- `models` - (Optional) List of model names that this team can access. If empty, assumes all models are allowed.

- `metadata` - (Optional) A map of metadata key-value pairs associated with the team. Store information for team tracking and organization.

- `blocked` - (Optional) Whether the team is blocked from making requests. Default is `false`. When blocked, all calls from keys with this team_id will be stopped.

- `tpm_limit` - (Optional) The TPM (Tokens Per Minute) limit for this team. All keys with this team_id will have at max this TPM limit.

- `rpm_limit` - (Optional) The RPM (Requests Per Minute) limit for this team. All keys associated with this team_id will have at max this RPM limit.

- `max_budget` - (Optional) The maximum budget allocated to the team. All keys for this team_id will have at max this max_budget.

- `budget_duration` - (Optional) The duration of the budget for the team. Budget is reset at the end of specified duration. If not set, budget is never reset. You can set duration as seconds ('30s'), minutes ('30m'), hours ('30h'), days ('30d'). Format must be: number followed by 's' (seconds), 'm' (minutes), 'h' (hours), or 'd' (days). Examples: '30s', '30m', '30h', '30d'.

- `team_member_permissions` - (Optional) A list of routes that non-admin team members can access. Example: `["/key/generate", "/key/update", "/key/delete"]`. Available permissions can be retrieved from the API endpoint `/team/permissions_list`.

- `team_member_budget` - (Optional) The maximum budget allocated to an individual team member. Budget automatically given to a new team member.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The unique identifier for the team (team_id).

## Import

Teams can be imported using either the traditional `terraform import` command or the newer import block syntax (recommended for Terraform 1.5+).

### Using terraform import command

```shell
terraform import litellm_team.engineering <team-id>
```

### Using import blocks (Terraform 1.5+)

Import blocks provide a declarative way to import existing resources. Add the following to your Terraform configuration:

```hcl
import {
  to = litellm_team.engineering
  id = "<team-id>"
}

resource "litellm_team" "engineering" {
  # Configuration will be populated after import
  team_alias = "engineering-team"
  # Add other required configuration here
}
```

After adding the import block, run:

```shell
terraform plan
```

Terraform will show you the configuration that needs to be added to match the imported resource. You can then run `terraform apply` to complete the import.

**Note:** The team ID is generated when the team is created and is different from the `team_alias`. Import blocks are the recommended approach for Terraform 1.5+ as they provide better integration with Terraform's planning and state management.

## Note on Team Members

Team members are managed through the separate `litellm_team_member` resource. This allows for more granular control over team membership and permissions. See the `litellm_team_member` resource documentation for details on managing team members.

## API Compatibility

This resource supports all the features available in the LiteLLM team API, including:

- Team creation with auto-generated UUIDs
- Budget and rate limiting controls
- Model access restrictions
- Team member permission management
- Organization-level team grouping
- Metadata for team tracking and organization

For more information about the underlying API, refer to the LiteLLM documentation.
