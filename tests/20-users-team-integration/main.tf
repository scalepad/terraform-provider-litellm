terraform {
  required_providers {
    litellm = {
      source  = "registry.terraform.io/local/litellm"
      version = "1.0.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.1"
    }
  }
}

provider "litellm" {
  api_base = "http://localhost:4000"
  api_key  = "sk-test-master-key-12345"
}

# Random strings to ensure unique identifiers
resource "random_string" "test_suffix" {
  length  = 8
  special = false
  upper   = false
}

resource "random_string" "team_suffix" {
  length  = 6
  special = false
  upper   = false
}

# Generate unique suffixes for each of the 20 users
resource "random_string" "user_suffix" {
  for_each = local.users
  length   = 6
  special  = false
  upper    = false
}

# =============================================================================
# LOCAL VALUES - USER DEFINITIONS
# =============================================================================

locals {
  # Define all 20 users with their configuration
  users = {
    user_01 = {
      number      = 1
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_02 = {
      number      = 2
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_03 = {
      number      = 3
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_04 = {
      number      = 4
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_05 = {
      number      = 5
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_06 = {
      number      = 6
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_07 = {
      number      = 7
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_08 = {
      number      = 8
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_09 = {
      number      = 9
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_10 = {
      number      = 10
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_11 = {
      number      = 11
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_12 = {
      number      = 12
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_13 = {
      number      = 13
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_14 = {
      number      = 14
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_15 = {
      number      = 15
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_16 = {
      number      = 16
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_17 = {
      number      = 17
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_18 = {
      number      = 18
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_19 = {
      number      = 19
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
    user_20 = {
      number      = 20
      role        = "internal_user"
      max_budget  = 100.0
      soft_budget = 80.0
      tpm_limit   = 5000
      rpm_limit   = 50
    }
  }
}

# =============================================================================
# TEAM CREATION
# =============================================================================

# Create a team to hold all 20 users
resource "litellm_team" "multi_user_team" {
  team_alias = "multi-user-test-team-${random_string.team_suffix.result}"

  max_budget         = 1000.0 # Higher budget to accommodate 20 users
  team_member_budget = 50.0
  budget_duration    = "30d"
  tpm_limit          = 100000 # Higher limits for multiple users
  rpm_limit          = 1000

  models = [
    "gpt-4",
    "gpt-3.5-turbo",
    "claude-3-sonnet"
  ]

  team_member_permissions = [
    "read",
    "/key/generate",
    "/key/info",
    "/user/info"
  ]

  blocked = false

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "20_users_team_integration"
    created_by  = "terraform"
    purpose     = "multi_user_team_test"
    user_count  = "20"
  }
}

# =============================================================================
# USER CREATION (20 USERS)
# =============================================================================

# Create 20 users using for_each
resource "litellm_user" "test_users" {
  for_each = local.users

  user_email = "test-${each.key}-${random_string.user_suffix[each.key].result}@test-litellm.com"
  user_alias = "Test User ${each.value.number} ${random_string.user_suffix[each.key].result}"
  user_role  = each.value.role

  # Budget and limits from local values
  max_budget            = each.value.max_budget
  soft_budget           = each.value.soft_budget
  budget_duration       = "30d"
  tpm_limit             = each.value.tpm_limit
  rpm_limit             = each.value.rpm_limit
  max_parallel_requests = 10

  # Model access
  models = [
    "gpt-4",
    "gpt-3.5-turbo",
    "claude-3-sonnet"
  ]

  # Model aliases for user convenience
  aliases = {
    "primary-gpt"  = "gpt-4"
    "fast-gpt"     = "gpt-3.5-turbo"
    "claude-smart" = "claude-3-sonnet"
  }

  # Email and key settings
  send_invite_email = true
  auto_create_key   = false
  blocked           = false

  # Cache controls
  allowed_cache_controls = ["no-cache", "max-age=3600"]

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "20_users_team_integration"
    created_by  = "terraform"
    purpose     = "multi_user_team_test"
    user_key    = each.key
    user_number = tostring(each.value.number)
    total_users = "20"
  }
}

# =============================================================================
# TEAM MEMBERSHIP ASSIGNMENT (ALL 20 USERS)
# =============================================================================

# Assign all 20 users to the team using for_each
resource "litellm_team_member" "team_members" {
  for_each = local.users

  team_id            = litellm_team.multi_user_team.id
  user_id            = litellm_user.test_users[each.key].user_id
  user_email         = litellm_user.test_users[each.key].user_email
  role               = "user"
  max_budget_in_team = 45.0

  # Ensure users and team are created before adding to team
  depends_on = [
    litellm_user.test_users,
    litellm_team.multi_user_team
  ]
}

# =============================================================================
# VERIFICATION OUTPUTS
# =============================================================================

output "test_summary" {
  description = "Summary of the 20-user team integration test"
  value = {
    test_name           = "20 Users Team Integration Test"
    test_id             = random_string.test_suffix.result
    team_created        = true
    users_created_count = length(litellm_user.test_users)
    memberships_added   = length(litellm_team_member.team_members)
    all_users_in_team   = length(litellm_team_member.team_members) == 20
  }
}

output "team_details" {
  description = "Details of the created team"
  value = {
    team_id            = litellm_team.multi_user_team.id
    team_alias         = litellm_team.multi_user_team.team_alias
    max_budget         = litellm_team.multi_user_team.max_budget
    team_member_budget = litellm_team.multi_user_team.team_member_budget
    budget_duration    = litellm_team.multi_user_team.budget_duration
    tpm_limit          = litellm_team.multi_user_team.tpm_limit
    rpm_limit          = litellm_team.multi_user_team.rpm_limit
    models             = litellm_team.multi_user_team.models
    member_permissions = litellm_team.multi_user_team.team_member_permissions
    blocked            = litellm_team.multi_user_team.blocked
  }
}

output "users_summary" {
  description = "Summary of all created users"
  value = {
    total_users  = length(litellm_user.test_users)
    user_emails  = { for k, v in litellm_user.test_users : k => v.user_email }
    user_ids     = { for k, v in litellm_user.test_users : k => v.user_id }
    user_aliases = { for k, v in litellm_user.test_users : k => v.user_alias }
  }
}

output "team_memberships_summary" {
  description = "Summary of all team memberships"
  value = {
    total_memberships = length(litellm_team_member.team_members)
    membership_ids    = { for k, v in litellm_team_member.team_members : k => v.id }
    member_emails     = { for k, v in litellm_team_member.team_members : k => v.user_email }
    member_roles      = { for k, v in litellm_team_member.team_members : k => v.role }
  }
}

# Detailed verification for each user membership
output "detailed_user_memberships" {
  description = "Detailed information for each user's team membership"
  value = {
    for user_key in keys(local.users) : user_key => {
      user_id            = litellm_user.test_users[user_key].user_id
      user_email         = litellm_user.test_users[user_key].user_email
      user_alias         = litellm_user.test_users[user_key].user_alias
      user_number        = local.users[user_key].number
      membership_id      = litellm_team_member.team_members[user_key].id
      team_id            = litellm_team_member.team_members[user_key].team_id
      role               = litellm_team_member.team_members[user_key].role
      max_budget_in_team = litellm_team_member.team_members[user_key].max_budget_in_team
      is_team_member     = litellm_team_member.team_members[user_key].team_id == litellm_team.multi_user_team.id
    }
  }
}

output "verification_checklist" {
  description = "Comprehensive verification checklist for the test"
  value = {
    team_created                = litellm_team.multi_user_team.id != ""
    expected_users_created      = length(litellm_user.test_users) == 20
    expected_memberships_added  = length(litellm_team_member.team_members) == 20
    all_users_have_membership   = length([for member in litellm_team_member.team_members : member.id if member.team_id == litellm_team.multi_user_team.id]) == 20
    all_memberships_valid       = length([for member in litellm_team_member.team_members : member.id if member.team_id != ""]) == 20
    all_users_have_email_invite = length([for user in litellm_user.test_users : user.user_id if user.send_invite_email]) == 20

    # Verification that all user IDs match their corresponding membership user IDs
    user_membership_consistency = length([
      for user_key in keys(local.users) : user_key
      if litellm_user.test_users[user_key].user_id == litellm_team_member.team_members[user_key].user_id
    ]) == 20

    # Overall test success
    test_passed = (
      litellm_team.multi_user_team.id != "" &&
      length(litellm_user.test_users) == 20 &&
      length(litellm_team_member.team_members) == 20 &&
      length([for member in litellm_team_member.team_members : member.id if member.team_id == litellm_team.multi_user_team.id]) == 20
    )
  }
}

# Email notifications summary
output "email_notifications_summary" {
  description = "Summary of email invitations sent to all users"
  value = {
    total_invites_sent  = length([for user in litellm_user.test_users : user.user_id if user.send_invite_email])
    invite_recipients   = { for k, v in litellm_user.test_users : k => v.user_email if v.send_invite_email }
    all_invites_enabled = length([for user in litellm_user.test_users : user.user_id if user.send_invite_email]) == 20
  }
}

# Performance and scale metrics
output "scale_metrics" {
  description = "Metrics related to the scale of this test"
  value = {
    users_per_team          = length(litellm_user.test_users)
    total_team_budget       = litellm_team.multi_user_team.max_budget
    total_user_budgets      = sum([for user in litellm_user.test_users : user.max_budget])
    team_tpm_limit          = litellm_team.multi_user_team.tpm_limit
    combined_user_tpm_limit = sum([for user in litellm_user.test_users : user.tpm_limit])
    team_rpm_limit          = litellm_team.multi_user_team.rpm_limit
    combined_user_rpm_limit = sum([for user in litellm_user.test_users : user.rpm_limit])
  }
}
