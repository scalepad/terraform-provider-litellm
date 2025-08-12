# Example Terraform configuration for LiteLLM User resource

terraform {
  required_providers {
    litellm = {
      source = "registry.terraform.io/your-org/litellm"
    }
  }
}

provider "litellm" {
  api_base = "https://your-litellm-instance.com"
  api_key  = var.litellm_api_key
}

# Basic Internal User
resource "litellm_user" "basic_internal" {
  user_email = "john.doe@company.com"
  user_alias = "John Doe"
  user_role  = "internal_user"

  # Basic budget and rate limits
  max_budget      = 100.0
  budget_duration = "1mo"
  tpm_limit       = 10000
  rpm_limit       = 100

  # Basic model access
  models = ["gpt-3.5-turbo", "gpt-4"]

  # User management
  send_invite_email = true
  auto_create_key   = true
  key_alias         = "john-doe-key"
  duration          = "30d"

  # Basic metadata
  metadata = {
    department  = "engineering"
    team        = "backend"
    environment = "production"
  }
}

# Admin User with Comprehensive Configuration
resource "litellm_user" "admin_user" {
  user_id    = "admin-user-001"
  user_email = "admin@company.com"
  user_alias = "System Administrator"
  user_role  = "proxy_admin"

  # Generous budget and limits for admin
  max_budget            = 5000.0
  soft_budget           = 4000.0
  budget_duration       = "1mo"
  tpm_limit             = 100000
  rpm_limit             = 1000
  max_parallel_requests = 50

  # Full model access
  models = [
    "gpt-4",
    "gpt-4-turbo",
    "gpt-3.5-turbo",
    "claude-3.5-sonnet",
    "claude-3-haiku",
    "gemini-pro"
  ]

  # Model aliases for convenience
  aliases = {
    "admin-gpt4"   = "gpt-4"
    "admin-claude" = "claude-3.5-sonnet"
    "admin-gemini" = "gemini-pro"
    "fast-model"   = "gpt-3.5-turbo"
  }

  # Administrative permissions
  permissions = {
    "admin_ui_access"   = "true"
    "model_management"  = "true"
    "user_management"   = "true"
    "budget_management" = "true"
    "system_monitoring" = "true"
  }

  # Object-level permissions
  object_permission = {
    "all_teams"     = "admin"
    "all_projects"  = "admin"
    "system_config" = "write"
  }

  # Security features
  guardrails             = ["admin_audit_log", "privileged_access_monitor"]
  allowed_cache_controls = ["no-cache", "no-store", "must-revalidate"]

  # User management
  send_invite_email = true
  auto_create_key   = true
  key_alias         = "admin-master-key"
  duration          = "90d"
  blocked           = false

  # Advanced features
  prompts       = ["admin_system_prompt", "security_prompt"]
  organizations = ["org_main", "org_security"]

  # Comprehensive metadata
  metadata = {
    department         = "it"
    team               = "platform"
    role               = "administrator"
    security_clearance = "high"
    cost_center        = "IT-OPS"
    environment        = "production"
  }
}

# Customer User with Budget Limits and Model Restrictions
resource "litellm_user" "customer_user" {
  user_email = "customer@external.com"
  user_alias = "External Customer"
  user_role  = "customer"

  # Strict budget controls
  max_budget            = 50.0
  soft_budget           = 40.0
  budget_duration       = "1mo"
  tpm_limit             = 5000
  rpm_limit             = 50
  max_parallel_requests = 5

  # Limited model access
  models = ["gpt-3.5-turbo", "claude-3-haiku"]

  # Model-specific budget limits
  model_max_budget = {
    "gpt-3.5-turbo"  = "30"
    "claude-3-haiku" = "20"
  }

  # Model-specific rate limits
  model_tpm_limit = {
    "gpt-3.5-turbo"  = "3000"
    "claude-3-haiku" = "2000"
  }

  model_rpm_limit = {
    "gpt-3.5-turbo"  = "30"
    "claude-3-haiku" = "20"
  }

  # Customer-specific aliases
  aliases = {
    "fast-chat"    = "gpt-3.5-turbo"
    "efficient-ai" = "claude-3-haiku"
  }

  # Limited permissions
  permissions = {
    "api_access"     = "true"
    "basic_features" = "true"
  }

  # Security and content filtering
  guardrails = [
    "content_filter",
    "pii_detection",
    "toxicity_filter",
    "customer_safety_check"
  ]

  # Cache restrictions
  allowed_cache_controls = ["no-store"]

  # User management
  send_invite_email = true
  auto_create_key   = true
  key_alias         = "customer-api-key"
  duration          = "30d"
  blocked           = false

  # Customer metadata
  metadata = {
    account_type      = "external"
    subscription_tier = "basic"
    industry          = "technology"
    region            = "north-america"
    support_level     = "standard"
  }
}

# Team User with Advanced Permissions and Guardrails
resource "litellm_user" "team_user" {
  user_email = "team-lead@company.com"
  user_alias = "AI Team Lead"
  user_role  = "team"

  # Team-level budget and limits
  max_budget            = 1000.0
  soft_budget           = 800.0
  budget_duration       = "1mo"
  tpm_limit             = 25000
  rpm_limit             = 250
  max_parallel_requests = 15

  # Balanced model access
  models = [
    "gpt-4",
    "gpt-3.5-turbo",
    "claude-3.5-sonnet",
    "claude-3-haiku"
  ]

  # Team-specific model limits
  model_max_budget = {
    "gpt-4"             = "600"
    "claude-3.5-sonnet" = "300"
    "gpt-3.5-turbo"     = "80"
    "claude-3-haiku"    = "20"
  }

  model_tpm_limit = {
    "gpt-4"             = "8000"
    "claude-3.5-sonnet" = "10000"
    "gpt-3.5-turbo"     = "5000"
    "claude-3-haiku"    = "2000"
  }

  # Team aliases
  aliases = {
    "team-primary"   = "gpt-4"
    "team-secondary" = "claude-3.5-sonnet"
    "team-fast"      = "gpt-3.5-turbo"
    "team-efficient" = "claude-3-haiku"
  }

  # Team permissions
  permissions = {
    "team_management"   = "true"
    "model_switching"   = "true"
    "usage_analytics"   = "true"
    "prompt_management" = "true"
  }

  # Team object permissions
  object_permission = {
    "team_ai_project" = "admin"
    "shared_prompts"  = "write"
    "team_models"     = "read"
  }

  # Advanced guardrails
  guardrails = [
    "content_filter",
    "code_safety_check",
    "team_compliance",
    "usage_monitoring"
  ]

  # Cache controls
  allowed_cache_controls = ["no-cache", "private"]

  # User management
  send_invite_email = true
  auto_create_key   = true
  key_alias         = "team-lead-key"
  duration          = "60d"
  blocked           = false

  # Team-specific prompts
  prompts = [
    "team_coding_prompt",
    "team_review_prompt",
    "safety_guidelines"
  ]

  # Organizations
  organizations = ["org_ai_team", "org_engineering"]

  # Team metadata
  metadata = {
    department  = "engineering"
    team        = "ai-platform"
    role        = "team_lead"
    project     = "ai-integration"
    cost_center = "R&D"
    environment = "production"
  }
}

# Enterprise User with SSO Integration and Complex Configuration
resource "litellm_user" "enterprise_user" {
  user_id    = "enterprise-sso-user-001"
  user_email = "enterprise.user@corp.com"
  user_alias = "Enterprise SSO User"
  user_role  = "internal_user"

  # Enterprise-level budget and limits
  max_budget            = 2000.0
  soft_budget           = 1500.0
  budget_duration       = "1mo"
  tpm_limit             = 50000
  rpm_limit             = 500
  max_parallel_requests = 25

  # Full enterprise model access
  models = [
    "gpt-4",
    "gpt-4-turbo",
    "gpt-3.5-turbo",
    "claude-3.5-sonnet",
    "claude-3-opus",
    "claude-3-haiku",
    "gemini-pro",
    "gemini-pro-vision"
  ]

  # Enterprise model budgets
  model_max_budget = {
    "gpt-4"             = "800"
    "claude-3.5-sonnet" = "600"
    "claude-3-opus"     = "400"
    "gemini-pro"        = "200"
  }

  # Enterprise aliases
  aliases = {
    "enterprise-premium"  = "gpt-4"
    "enterprise-balanced" = "claude-3.5-sonnet"
    "enterprise-fast"     = "gpt-3.5-turbo"
    "enterprise-vision"   = "gemini-pro-vision"
  }

  # Enterprise permissions
  permissions = {
    "enterprise_features" = "true"
    "advanced_analytics"  = "true"
    "custom_models"       = "true"
    "integration_access"  = "true"
    "compliance_tools"    = "true"
  }

  # Complex object permissions
  object_permission = {
    "enterprise_project_alpha"    = "admin"
    "enterprise_project_beta"     = "write"
    "shared_enterprise_resources" = "read"
    "compliance_data"             = "read"
  }

  # Enterprise-grade guardrails
  guardrails = [
    "enterprise_content_filter",
    "compliance_check",
    "data_classification",
    "audit_logging",
    "pii_detection",
    "enterprise_security_scan"
  ]

  # Enterprise cache controls
  allowed_cache_controls = [
    "no-cache",
    "no-store",
    "private",
    "must-revalidate"
  ]

  # SSO integration
  sso_user_id = "sso_enterprise_id_12345"

  # User management
  send_invite_email = false # SSO users don't need email invites
  auto_create_key   = true
  key_alias         = "enterprise-sso-key"
  duration          = "180d" # Longer duration for enterprise users
  blocked           = false

  # Enterprise prompts
  prompts = [
    "enterprise_system_prompt",
    "compliance_prompt",
    "security_guidelines",
    "enterprise_best_practices"
  ]

  # Multiple organizations
  organizations = [
    "org_enterprise_main",
    "org_compliance",
    "org_security",
    "org_innovation"
  ]

  # Comprehensive enterprise metadata
  metadata = {
    department         = "enterprise_solutions"
    team               = "digital_transformation"
    role               = "senior_architect"
    employee_id        = "EMP12345"
    cost_center        = "ENTERPRISE-TECH"
    security_clearance = "confidential"
    compliance_level   = "sox_compliant"
    region             = "global"
    business_unit      = "enterprise_services"
    environment        = "production"
  }

  # Legacy config for backward compatibility
  config = {
    "legacy_feature_flag" = "enabled"
    "enterprise_mode"     = "true"
  }
}

# Output the user IDs for reference
output "user_ids" {
  description = "IDs of the created users"
  value = {
    basic_internal  = litellm_user.basic_internal.user_id
    admin_user      = litellm_user.admin_user.user_id
    customer_user   = litellm_user.customer_user.user_id
    team_user       = litellm_user.team_user.user_id
    enterprise_user = litellm_user.enterprise_user.user_id
  }
}

# Output user status and spend information
output "user_status" {
  description = "Status and spend information for users"
  value = {
    basic_internal = {
      spend           = litellm_user.basic_internal.spend
      key_count       = litellm_user.basic_internal.key_count
      created_at      = litellm_user.basic_internal.created_at
      budget_reset_at = litellm_user.basic_internal.budget_reset_at
    }
    admin_user = {
      spend           = litellm_user.admin_user.spend
      key_count       = litellm_user.admin_user.key_count
      created_at      = litellm_user.admin_user.created_at
      budget_reset_at = litellm_user.admin_user.budget_reset_at
    }
    customer_user = {
      spend           = litellm_user.customer_user.spend
      key_count       = litellm_user.customer_user.key_count
      created_at      = litellm_user.customer_user.created_at
      budget_reset_at = litellm_user.customer_user.budget_reset_at
    }
    team_user = {
      spend           = litellm_user.team_user.spend
      key_count       = litellm_user.team_user.key_count
      created_at      = litellm_user.team_user.created_at
      budget_reset_at = litellm_user.team_user.budget_reset_at
    }
    enterprise_user = {
      spend           = litellm_user.enterprise_user.spend
      key_count       = litellm_user.enterprise_user.key_count
      created_at      = litellm_user.enterprise_user.created_at
      budget_reset_at = litellm_user.enterprise_user.budget_reset_at
    }
  }
}

# Output budget and usage information
output "user_budget_info" {
  description = "Budget and usage information for users"
  value = {
    basic_internal = {
      max_budget      = litellm_user.basic_internal.max_budget
      budget_duration = litellm_user.basic_internal.budget_duration
      current_spend   = litellm_user.basic_internal.spend
    }
    admin_user = {
      max_budget      = litellm_user.admin_user.max_budget
      soft_budget     = litellm_user.admin_user.soft_budget
      budget_duration = litellm_user.admin_user.budget_duration
      current_spend   = litellm_user.admin_user.spend
    }
    customer_user = {
      max_budget      = litellm_user.customer_user.max_budget
      soft_budget     = litellm_user.customer_user.soft_budget
      budget_duration = litellm_user.customer_user.budget_duration
      current_spend   = litellm_user.customer_user.spend
    }
    team_user = {
      max_budget      = litellm_user.team_user.max_budget
      soft_budget     = litellm_user.team_user.soft_budget
      budget_duration = litellm_user.team_user.budget_duration
      current_spend   = litellm_user.team_user.spend
    }
    enterprise_user = {
      max_budget      = litellm_user.enterprise_user.max_budget
      soft_budget     = litellm_user.enterprise_user.soft_budget
      budget_duration = litellm_user.enterprise_user.budget_duration
      current_spend   = litellm_user.enterprise_user.spend
    }
  }
}

# Variable for API key
variable "litellm_api_key" {
  description = "API key for LiteLLM"
  type        = string
  sensitive   = true
}
