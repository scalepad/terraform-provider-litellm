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
}
```

## Argument Reference

The following arguments are supported:

* `team_alias` - (Required) A human-readable identifier for the team.

* `organization_id` - (Optional) The ID of the organization this team belongs to.

* `models` - (Optional) List of model names that this team can access.

* `metadata` - (Optional) A map of metadata key-value pairs associated with the team.

* `blocked` - (Optional) Whether the team is blocked from making requests. Default is `false`.

* `tpm_limit` - (Optional) Team-wide tokens per minute limit.

* `rpm_limit` - (Optional) Team-wide requests per minute limit.

* `max_budget` - (Optional) Maximum budget allocated to the team.

* `budget_duration` - (Optional) Duration for the budget cycle. Valid values are:
  * `daily`
  * `weekly`
  * `monthly`
  * `yearly`

* `team_member_permissions` - (Optional) List of permissions granted to team members. Available permissions can be retrieved from the API endpoint `/team/permissions_list`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The unique identifier for the team.

## Import

Teams can be imported using the team ID:

```shell
terraform import litellm_team.engineering <team-id>
```

Note: The team ID is generated when the team is created and is different from the `team_alias`.

## Note on Team Members

Team members are managed through the separate `litellm_team_member` resource. This allows for more granular control over team membership and permissions. See the `litellm_team_member` resource documentation for details on managing team members.
