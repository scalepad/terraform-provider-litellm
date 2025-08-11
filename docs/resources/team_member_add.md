# Resource: litellm_team_member_add

Add multiple members to a team with a single resource.

## Example Usage

### Basic Usage

```hcl
resource "litellm_team_member_add" "example" {
  team_id = "team-123"

  member {
    user_id = "user-456"
    role    = "admin"
  }

  member {
    user_email = "user@example.com"
    role       = "user"
  }

  max_budget_in_team = 100.0
}
```

### Dynamic Members Using Locals

```hcl
locals {
  team_members = [
    {
      user_id  = "user-123"
      role     = "admin"
    },
    {
      user_email = "developer1@company.com"
      role      = "user"
    },
    {
      user_email = "developer2@company.com"
      role      = "user"
    }
  ]
}

resource "litellm_team_member_add" "dynamic_example" {
  team_id = "team-456"

  dynamic "member" {
    for_each = local.team_members
    content {
      user_id    = lookup(member.value, "user_id", null)
      user_email = lookup(member.value, "user_email", null)
      role       = member.value.role
    }
  }

  max_budget_in_team = 200.0
}
```

## Argument Reference

- `team_id` - (Required) The ID of the team to add members to.
- `member` - (Required) One or more member blocks defining team members. Each block supports:
  - `user_id` - (Optional) The ID of the user to add to the team.
  - `user_email` - (Optional) The email of the user to add to the team.
  - `role` - (Required) The role of the user in the team. Valid values are:
    - `admin` - Team administrator with management permissions
    - `user` - Standard team member
- `max_budget_in_team` - (Optional) The maximum budget allocated for the team members.

## Note

This resource does not support import functionality. Team members must be managed through the resource configuration.
