terraform {
  required_providers {
    litellm = {
      source  = "registry.terraform.io/local/litellm"
      version = "1.0.0"
    }
  }
}

provider "litellm" {
  api_base = "http://localhost:4000"
  api_key  = "sk-test-master-key-12345"
}

# Test basic team creation
resource "litellm_team" "basic_team" {
  team_alias = "basic-test-team"

  max_budget      = 100.0
  budget_duration = "30d"
  tpm_limit       = 10000
  rpm_limit       = 100

  models = ["gpt-4", "gpt-3.5-turbo"]

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "basic_team"
  }
}

# Test team with restrictions
resource "litellm_team" "restricted_team" {
  team_alias = "restricted-test-team"

  max_budget         = 25.0
  team_member_budget = 10.0
  budget_duration    = "7d"
  tpm_limit          = 2500
  rpm_limit          = 25

  models = ["gpt-3.5-turbo"]

  team_member_permissions = [
    "read",
    "/key/generate"
  ]

  blocked = false

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "restricted_team"
    tier        = "basic"
  }
}

# Outputs for verification
output "basic_team_id" {
  description = "ID of the basic test team"
  value       = litellm_team.basic_team.id
}

output "restricted_team_id" {
  description = "ID of the restricted test team"
  value       = litellm_team.restricted_team.id
}

output "team_budget_info" {
  description = "Budget information for all test teams"
  value = {
    basic_team = {
      max_budget         = litellm_team.basic_team.max_budget
      team_member_budget = litellm_team.basic_team.team_member_budget
      budget_duration    = litellm_team.basic_team.budget_duration
    }
    restricted_team = {
      max_budget         = litellm_team.restricted_team.max_budget
      team_member_budget = litellm_team.restricted_team.team_member_budget
      budget_duration    = litellm_team.restricted_team.budget_duration
    }
  }
}

resource "litellm_team_member" "basic_user" {
  team_id            = litellm_team.basic_team.id
  user_id            = "test-user-member-001"
  user_email         = "user@test-basic-team.com"
  role               = "user"
  max_budget_in_team = 25.0
}

# Test team member without budget limit
resource "litellm_team_member" "restricted_user_no_budget" {
  team_id    = litellm_team.restricted_team.id
  user_id    = "test-restricted-user-001"
  user_email = "restricted@test-restricted-team.com"
  role       = "user"
  # No max_budget_in_team specified - should use team default
}

# Test dynamic team member management using locals
locals {
  restricted_team_members = [
    {
      user_id = "dynamic-user-001"
      role    = "user"
    },
    {
      user_email = "dynamic-user1@test-restricted-team.com"
      role       = "user"
    },
    {
      user_email = "dynamic-user2@test-restricted-team.com"
      role       = "user"
    }
  ]
}

resource "litellm_team_member_add" "restricted_team_dynamic" {
  team_id = litellm_team.restricted_team.id

  dynamic "member" {
    for_each = local.restricted_team_members
    content {
      user_id    = lookup(member.value, "user_id", null)
      user_email = lookup(member.value, "user_email", null)
      role       = member.value.role
    }
  }

  max_budget_in_team = 15.0
}

# Outputs for team member verification
output "individual_team_members" {
  description = "Information about individual team members"
  value = {
    basic_user = {
      id                 = litellm_team_member.basic_user.id
      team_id            = litellm_team_member.basic_user.team_id
      user_id            = litellm_team_member.basic_user.user_id
      user_email         = litellm_team_member.basic_user.user_email
      role               = litellm_team_member.basic_user.role
      max_budget_in_team = litellm_team_member.basic_user.max_budget_in_team
    }
    restricted_user_no_budget = {
      id                 = litellm_team_member.restricted_user_no_budget.id
      team_id            = litellm_team_member.restricted_user_no_budget.team_id
      user_id            = litellm_team_member.restricted_user_no_budget.user_id
      user_email         = litellm_team_member.restricted_user_no_budget.user_email
      role               = litellm_team_member.restricted_user_no_budget.role
      max_budget_in_team = litellm_team_member.restricted_user_no_budget.max_budget_in_team
    }
  }
}

output "bulk_team_members" {
  description = "Information about bulk team member resources"
  value = {
    restricted_team_dynamic = {
      id                 = litellm_team_member_add.restricted_team_dynamic.id
      team_id            = litellm_team_member_add.restricted_team_dynamic.team_id
      max_budget_in_team = litellm_team_member_add.restricted_team_dynamic.max_budget_in_team
    }
  }
}

output "team_member_test_summary" {
  description = "Summary of team member test configurations"
  value = {
    individual_members_count    = 2
    bulk_member_resources_count = 1
    roles_tested = [
      "user"
    ]
    test_scenarios = [
      "individual_member_management",
      "bulk_member_management",
      "dynamic_member_configuration",
      "mixed_user_identifiers",
      "budget_limits_testing",
      "role_permissions_testing"
    ]
  }
}
