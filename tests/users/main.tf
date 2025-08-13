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

# Test basic user creation
resource "litellm_user" "basic_user" {
  user_email = "basic@test.com"
  user_alias = "Basic Test User"
  user_role  = "internal_user"

  max_budget      = 50.0
  budget_duration = "1mo"
  tpm_limit       = 5000
  rpm_limit       = 50

  models = ["gpt-4"]

  send_invite_email = false
  auto_create_key   = true
  key_alias         = "basic-test-key"
  duration          = "30d"

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "basic_user"
  }
}

# Test admin user with comprehensive configuration
resource "litellm_user" "admin_user" {
  user_email = "admin@test.com"
  user_alias = "Admin Test User"
  user_role  = "proxy_admin"

  max_budget            = 500.0
  soft_budget           = 400.0
  budget_duration       = "1mo"
  tpm_limit             = 50000
  rpm_limit             = 500
  max_parallel_requests = 25

  models = [
    "gpt-4",
    "gpt-3.5-turbo",
    "claude-3-sonnet"
  ]

  aliases = {
    "admin-gpt4"   = "gpt-4"
    "admin-claude" = "claude-3-sonnet"
    "fast-model"   = "gpt-3.5-turbo"
  }

  allowed_cache_controls = ["no-cache", "no-store"]

  send_invite_email = false
  auto_create_key   = true
  key_alias         = "admin-test-key"
  duration          = "90d"
  blocked           = false

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "admin_user"
    role        = "administrator"
  }
}

# Test user without auto-created key
resource "litellm_user" "no_key_user" {
  user_email = "nokey@test.com"
  user_alias = "No Key Test User"
  user_role  = "internal_user"

  max_budget      = 25.0
  budget_duration = "1mo"
  tpm_limit       = 2500
  rpm_limit       = 25

  models = ["gpt-3.5-turbo"]

  send_invite_email = false
  auto_create_key   = false # No automatic key creation

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "no_key_user"
    key_policy  = "manual"
  }
}

# Outputs for verification
output "basic_user_id" {
  description = "ID of the basic test user"
  value       = litellm_user.basic_user.user_id
}

output "admin_user_id" {
  description = "ID of the admin test user"
  value       = litellm_user.admin_user.user_id
}

output "no_key_user_id" {
  description = "ID of the no key test user"
  value       = litellm_user.no_key_user.user_id
}

output "user_status" {
  description = "Status information for all test users"
  value = {
    basic_user = {
      spend           = litellm_user.basic_user.spend
      key_count       = litellm_user.basic_user.key_count
      created_at      = litellm_user.basic_user.created_at
      budget_reset_at = litellm_user.basic_user.budget_reset_at
    }
    admin_user = {
      spend           = litellm_user.admin_user.spend
      key_count       = litellm_user.admin_user.key_count
      created_at      = litellm_user.admin_user.created_at
      budget_reset_at = litellm_user.admin_user.budget_reset_at
    }
    no_key_user = {
      spend           = litellm_user.no_key_user.spend
      key_count       = litellm_user.no_key_user.key_count
      created_at      = litellm_user.no_key_user.created_at
      budget_reset_at = litellm_user.no_key_user.budget_reset_at
    }
  }
}

output "user_budget_info" {
  description = "Budget information for all test users"
  value = {
    basic_user = {
      max_budget      = litellm_user.basic_user.max_budget
      budget_duration = litellm_user.basic_user.budget_duration
      current_spend   = litellm_user.basic_user.spend
    }
    admin_user = {
      max_budget      = litellm_user.admin_user.max_budget
      soft_budget     = litellm_user.admin_user.soft_budget
      budget_duration = litellm_user.admin_user.budget_duration
      current_spend   = litellm_user.admin_user.spend
    }
    no_key_user = {
      max_budget      = litellm_user.no_key_user.max_budget
      budget_duration = litellm_user.no_key_user.budget_duration
      current_spend   = litellm_user.no_key_user.spend
    }
  }
}
