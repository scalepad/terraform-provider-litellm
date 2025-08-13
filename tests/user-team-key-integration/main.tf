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

resource "random_string" "user_suffix" {
  length  = 6
  special = false
  upper   = false
}

resource "random_string" "team_suffix" {
  length  = 6
  special = false
  upper   = false
}

# =============================================================================
# TEAM CREATION
# =============================================================================

# Create a team first
resource "litellm_team" "integration_team" {
  team_alias = "integration-test-team-${random_string.team_suffix.result}"

  max_budget         = 200.0
  team_member_budget = 50.0
  budget_duration    = "30d"
  tpm_limit          = 20000
  rpm_limit          = 200

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
    test_type     = "integration"
    environment   = "ci"
    test_case     = "user_team_key_integration"
    created_by    = "terraform"
    purpose       = "comprehensive_integration_test"
    email_enabled = "true"
  }
}

# =============================================================================
# USER CREATION WITH EMAIL INVITE
# =============================================================================

# Create a user with email invite enabled
resource "litellm_user" "integration_user" {
  user_email = "integration-user-${random_string.user_suffix.result}@test-litellm.com"
  user_alias = "Integration Test User ${random_string.user_suffix.result}"
  user_role  = "internal_user"

  # Budget and limits
  max_budget            = 100.0
  soft_budget           = 80.0
  budget_duration       = "30d"
  tpm_limit             = 10000
  rpm_limit             = 100
  max_parallel_requests = 15

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

  # Model-specific limits removed - requires enterprise license
  # model_max_budget = {
  #   "gpt-4"           = "60"
  #   "gpt-3.5-turbo"   = "25"
  #   "claude-3-sonnet" = "15"
  # }

  # model_tpm_limit = {
  #   "gpt-4"           = "5000"
  #   "gpt-3.5-turbo"   = "3000"
  #   "claude-3-sonnet" = "2000"
  # }

  # model_rpm_limit = {
  #   "gpt-4"           = "50"
  #   "gpt-3.5-turbo"   = "30"
  #   "claude-3-sonnet" = "20"
  # }

  # Email and key settings - IMPORTANT: Enable email invite
  send_invite_email = true
  auto_create_key   = false # We'll create the key manually with team context
  blocked           = false

  # Cache controls
  allowed_cache_controls = ["no-cache", "max-age=3600"]

  metadata = {
    test_type     = "integration"
    environment   = "ci"
    test_case     = "user_team_key_integration"
    created_by    = "terraform"
    purpose       = "comprehensive_integration_test"
    email_enabled = "true"
    team_member   = "true"
  }
}

# =============================================================================
# TEAM MEMBERSHIP ASSIGNMENT
# =============================================================================

# Assign the user to the team
resource "litellm_team_member" "integration_member" {
  team_id            = litellm_team.integration_team.id
  user_id            = litellm_user.integration_user.user_id
  user_email         = litellm_user.integration_user.user_email
  role               = "user"
  max_budget_in_team = 45.0 # Slightly less than team member budget

  # Ensure user is created before adding to team
  depends_on = [
    litellm_user.integration_user,
    litellm_team.integration_team
  ]
}

# =============================================================================
# KEY GENERATION WITH EMAIL NOTIFICATION
# =============================================================================

# Create a key for the user within the team context with email notification
resource "litellm_key" "integration_user_key" {
  # Associate with both user and team
  user_id = litellm_user.integration_user.user_id
  team_id = litellm_team.integration_team.id

  # Model access (subset of what user and team allow)
  models = [
    "gpt-4",
    "gpt-3.5-turbo"
  ]

  # Budget and limits (within user and team constraints)
  max_budget            = 40.0 # Less than user's team budget
  soft_budget           = 30.0
  budget_duration       = "30d"
  tpm_limit             = 8000 # Less than user limit
  rpm_limit             = 80   # Less than user limit
  max_parallel_requests = 10   # Less than user limit

  # Model-specific limits removed - requires enterprise license
  # model_max_budget = {
  #   "gpt-4"         = "25"
  #   "gpt-3.5-turbo" = "15"
  # }

  # model_tpm_limit = {
  #   "gpt-4"         = "4000"
  #   "gpt-3.5-turbo" = "4000"
  # }

  # model_rpm_limit = {
  #   "gpt-4"         = "40"
  #   "gpt-3.5-turbo" = "40"
  # }

  # Key configuration
  key_alias = "integration-team-key-${random_string.test_suffix.result}"
  duration  = "30d"
  blocked   = false

  # Model aliases for this key
  aliases = {
    "team-gpt4" = "gpt-4"
    "team-fast" = "gpt-3.5-turbo"
  }

  # Cache controls
  allowed_cache_controls = ["no-cache"]

  # IMPORTANT: Enable email notification for key delivery
  send_invite_email = true

  # Key permissions removed - requires enterprise license
  # permissions = {
  #   "can_create_keys" = "false"
  #   "can_view_usage"  = "true"
  #   "can_view_keys"   = "true"
  # }

  # Guardrails removed - requires enterprise license
  # guardrails = ["content_filter"]

  # Tags removed - requires enterprise license
  # tags = [
  #   "integration-test",
  #   "team-member",
  #   "email-enabled"
  # ]

  metadata = {
    test_type     = "integration"
    environment   = "ci"
    test_case     = "user_team_key_integration"
    created_by    = "terraform"
    purpose       = "team_member_key"
    email_enabled = "true"
    user_context  = litellm_user.integration_user.user_id
    team_context  = litellm_team.integration_team.id
    member_role   = litellm_team_member.integration_member.role
  }

  key_type = "llm_api"

  # Ensure all dependencies are created first
  depends_on = [
    litellm_user.integration_user,
    litellm_team.integration_team,
    litellm_team_member.integration_member
  ]
}

# =============================================================================
# VERIFICATION OUTPUTS
# =============================================================================

output "test_summary" {
  description = "Summary of the integration test setup"
  value = {
    test_name        = "User-Team-Key Integration with Email Notifications"
    test_id          = random_string.test_suffix.result
    user_created     = true
    team_created     = true
    membership_added = true
    key_generated    = true
    email_invites = {
      user_invite  = litellm_user.integration_user.send_invite_email
      key_delivery = litellm_key.integration_user_key.send_invite_email
    }
  }
}

output "user_details" {
  description = "Details of the created user"
  value = {
    user_id           = litellm_user.integration_user.user_id
    user_email        = litellm_user.integration_user.user_email
    user_alias        = litellm_user.integration_user.user_alias
    user_role         = litellm_user.integration_user.user_role
    max_budget        = litellm_user.integration_user.max_budget
    current_spend     = litellm_user.integration_user.spend
    key_count         = litellm_user.integration_user.key_count
    created_at        = litellm_user.integration_user.created_at
    send_invite_email = litellm_user.integration_user.send_invite_email
  }
}

output "team_details" {
  description = "Details of the created team"
  value = {
    team_id            = litellm_team.integration_team.id
    team_alias         = litellm_team.integration_team.team_alias
    max_budget         = litellm_team.integration_team.max_budget
    team_member_budget = litellm_team.integration_team.team_member_budget
    budget_duration    = litellm_team.integration_team.budget_duration
    tpm_limit          = litellm_team.integration_team.tpm_limit
    rpm_limit          = litellm_team.integration_team.rpm_limit
    models             = litellm_team.integration_team.models
    member_permissions = litellm_team.integration_team.team_member_permissions
    blocked            = litellm_team.integration_team.blocked
  }
}

output "team_membership_details" {
  description = "Details of the team membership"
  value = {
    membership_id      = litellm_team_member.integration_member.id
    team_id            = litellm_team_member.integration_member.team_id
    user_id            = litellm_team_member.integration_member.user_id
    user_email         = litellm_team_member.integration_member.user_email
    role               = litellm_team_member.integration_member.role
    max_budget_in_team = litellm_team_member.integration_member.max_budget_in_team
  }
}

output "key_details" {
  description = "Details of the generated key"
  value = {
    key_alias         = litellm_key.integration_user_key.key_alias
    user_id           = litellm_key.integration_user_key.user_id
    team_id           = litellm_key.integration_user_key.team_id
    models            = litellm_key.integration_user_key.models
    max_budget        = litellm_key.integration_user_key.max_budget
    soft_budget       = litellm_key.integration_user_key.soft_budget
    current_spend     = litellm_key.integration_user_key.spend
    tpm_limit         = litellm_key.integration_user_key.tpm_limit
    rpm_limit         = litellm_key.integration_user_key.rpm_limit
    duration          = litellm_key.integration_user_key.duration
    send_invite_email = litellm_key.integration_user_key.send_invite_email
    blocked           = litellm_key.integration_user_key.blocked
    # Enterprise features removed:
    # tags              = litellm_key.integration_user_key.tags
    # guardrails        = litellm_key.integration_user_key.guardrails
  }
}

# Sensitive output for the actual key token
output "key_token" {
  description = "The actual API key token (sensitive)"
  value       = litellm_key.integration_user_key.key
  sensitive   = true
}

output "email_notifications" {
  description = "Email notification status"
  value = {
    user_invite_sent = {
      enabled    = litellm_user.integration_user.send_invite_email
      recipient  = litellm_user.integration_user.user_email
      user_alias = litellm_user.integration_user.user_alias
    }
    key_delivery_sent = {
      enabled      = litellm_key.integration_user_key.send_invite_email
      recipient    = litellm_user.integration_user.user_email
      key_alias    = litellm_key.integration_user_key.key_alias
      team_context = litellm_team.integration_team.team_alias
    }
  }
}

output "integration_test_verification" {
  description = "Verification checklist for the integration test"
  sensitive   = true
  value = {
    checklist = {
      "1_user_created"          = litellm_user.integration_user.user_id != ""
      "2_team_created"          = litellm_team.integration_team.id != ""
      "3_user_assigned_to_team" = litellm_team_member.integration_member.id != ""
      "4_key_generated"         = litellm_key.integration_user_key.key != ""
      "5_user_invite_enabled"   = litellm_user.integration_user.send_invite_email
      "6_key_email_enabled"     = litellm_key.integration_user_key.send_invite_email
      "7_key_has_team_context"  = litellm_key.integration_user_key.team_id == litellm_team.integration_team.id
      "8_key_has_user_context"  = litellm_key.integration_user_key.user_id == litellm_user.integration_user.user_id
    }
    all_requirements_met = (
      litellm_user.integration_user.user_id != "" &&
      litellm_team.integration_team.id != "" &&
      litellm_team_member.integration_member.id != "" &&
      litellm_key.integration_user_key.key != "" &&
      litellm_user.integration_user.send_invite_email &&
      litellm_key.integration_user_key.send_invite_email
    )
  }
}
