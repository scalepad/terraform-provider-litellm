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

  blocked         = false
  tpm_limit       = 500000
  rpm_limit       = 5000
  max_budget      = 1000.0
  budget_duration = "monthly"
  
  # Team member permissions
  team_member_permissions = [
    "create_key",
    "delete_key",
    "view_spend",
    "edit_team"
  ]
}
```

## Argument Reference

The following arguments are supported:

- `team_alias` - (Required) A human-readable identifier for the team.

- `organization_id` - (Optional) The ID of the organization this team belongs to.

- `models` - (Optional) List of model names that this team can access.

- `metadata` - (Optional) A map of metadata key-value pairs associated with the team.

- `blocked` - (Optional) Whether the team is blocked from making requests. Default is `false`.

- `tpm_limit` - (Optional) Team-wide tokens per minute limit.

- `rpm_limit` - (Optional) Team-wide requests per minute limit.

- `max_budget` - (Optional) Maximum budget allocated to the team.

- `budget_duration` - (Optional) Duration for the budget cycle. Valid values are:

  - `daily`
  - `weekly`
  - `monthly`
  - `yearly`

- `team_member_permissions` - (Optional) List of permissions granted to team members. Available permissions can be retrieved from the API endpoint `/team/permissions_list`.

* `team_member_permissions` - (Optional) List of permissions granted to team members. This controls what actions team members can perform within the team context.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The unique identifier for the team.

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
