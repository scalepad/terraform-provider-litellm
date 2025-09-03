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

# Random strings to ensure unique aliases
resource "random_string" "team_suffix" {
  length  = 8
  special = false
  upper   = false
}

resource "random_string" "service_account_suffix" {
  length  = 8
  special = false
  upper   = false
}

provider "litellm" {
  api_base = "http://localhost:4000"
  api_key  = "sk-test-master-key-12345"
}

# Test teams for service accounts
resource "litellm_team" "production_team" {
  team_alias = "production-service-team-${random_string.team_suffix.result}"

  max_budget      = 1000.0
  budget_duration = "30d"
  tpm_limit       = 50000
  rpm_limit       = 500

  models = ["gpt-4", "gpt-3.5-turbo", "claude-3.5-sonnet"]

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "service_accounts"
    team_type   = "production"
    created_by  = "terraform"
  }
}

resource "litellm_team" "staging_team" {
  team_alias = "staging-service-team-${random_string.team_suffix.result}"

  max_budget      = 500.0
  budget_duration = "7d"
  tpm_limit       = 25000
  rpm_limit       = 250

  models = ["gpt-4", "gpt-3.5-turbo"]

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "service_accounts"
    team_type   = "staging"
    created_by  = "terraform"
  }
}

resource "litellm_team" "development_team" {
  team_alias = "development-service-team-${random_string.team_suffix.result}"

  max_budget      = 100.0
  budget_duration = "1d"
  tpm_limit       = 5000
  rpm_limit       = 50

  models = ["gpt-3.5-turbo"]

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "service_accounts"
    team_type   = "development"
    created_by  = "terraform"
  }
}

# Service accounts for production team
resource "litellm_service_account" "production_api_service" {
  team_id   = litellm_team.production_team.id
  key_alias = "production-api-service-${random_string.service_account_suffix.result}"

  models          = ["all-team-models"]
  max_budget      = 500.0
  tpm_limit       = 25000
  rpm_limit       = 250
  budget_duration = "30d"

  metadata = {
    test_type    = "integration"
    environment  = "ci"
    test_case    = "service_accounts"
    service_type = "api"
    team_type    = "production"
    created_by   = "terraform"
    owner        = "platform-team"
  }

  service_account_id = "prod-api-service-${random_string.service_account_suffix.result}"
}

resource "litellm_service_account" "production_worker_service" {
  team_id   = litellm_team.production_team.id
  key_alias = "production-worker-service-${random_string.service_account_suffix.result}"

  models          = ["gpt-4", "claude-3.5-sonnet"]
  max_budget      = 300.0
  tpm_limit       = 15000
  rpm_limit       = 150
  budget_duration = "30d"

  metadata = {
    test_type    = "integration"
    environment  = "ci"
    test_case    = "service_accounts"
    service_type = "worker"
    team_type    = "production"
    created_by   = "terraform"
    owner        = "data-team"
  }
}

# Service accounts for staging team
resource "litellm_service_account" "staging_api_service" {
  team_id   = litellm_team.staging_team.id
  key_alias = "staging-api-service-${random_string.service_account_suffix.result}"

  models          = ["all-team-models"]
  max_budget      = 200.0
  tpm_limit       = 10000
  rpm_limit       = 100
  budget_duration = "7d"

  metadata = {
    test_type    = "integration"
    environment  = "ci"
    test_case    = "service_accounts"
    service_type = "api"
    team_type    = "staging"
    created_by   = "terraform"
    owner        = "platform-team"
  }
}

# Service accounts for development team
resource "litellm_service_account" "development_service" {
  team_id   = litellm_team.development_team.id
  key_alias = "development-service-${random_string.service_account_suffix.result}"

  models          = ["gpt-3.5-turbo"]
  max_budget      = 50.0
  tpm_limit       = 2500
  rpm_limit       = 25
  budget_duration = "1d"

  metadata = {
    test_type    = "integration"
    environment  = "ci"
    test_case    = "service_accounts"
    service_type = "development"
    team_type    = "development"
    created_by   = "terraform"
    owner        = "developer"
  }
}

# Service account with minimal configuration (auto-generated service_account_id)
resource "litellm_service_account" "minimal_service" {
  team_id   = litellm_team.development_team.id
  key_alias = "minimal-service-${random_string.service_account_suffix.result}"

  models     = ["gpt-3.5-turbo"]
  max_budget = 10.0

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "minimal_service_account"
    created_by  = "terraform"
  }
  # service_account_id will be auto-generated
}

# Outputs for verification
output "team_ids" {
  description = "IDs of all test teams"
  value = {
    production_team  = litellm_team.production_team.id
    staging_team     = litellm_team.staging_team.id
    development_team = litellm_team.development_team.id
  }
}

output "service_account_keys" {
  description = "Generated service account keys (sensitive)"
  value = {
    production_api_service = {
      key      = litellm_service_account.production_api_service.key
      token_id = litellm_service_account.production_api_service.token_id
      key_name = litellm_service_account.production_api_service.key_name
    }
    production_worker_service = {
      key      = litellm_service_account.production_worker_service.key
      token_id = litellm_service_account.production_worker_service.token_id
      key_name = litellm_service_account.production_worker_service.key_name
    }
    staging_api_service = {
      key      = litellm_service_account.staging_api_service.key
      token_id = litellm_service_account.staging_api_service.token_id
      key_name = litellm_service_account.staging_api_service.key_name
    }
    development_service = {
      key      = litellm_service_account.development_service.key
      token_id = litellm_service_account.development_service.token_id
      key_name = litellm_service_account.development_service.key_name
    }
    minimal_service = {
      key      = litellm_service_account.minimal_service.key
      token_id = litellm_service_account.minimal_service.token_id
      key_name = litellm_service_account.minimal_service.key_name
    }
  }
  sensitive = true
}

output "service_account_info" {
  description = "Information about all service accounts"
  value = {
    production_api_service = {
      key_alias          = litellm_service_account.production_api_service.key_alias
      team_id            = litellm_service_account.production_api_service.team_id
      service_account_id = litellm_service_account.production_api_service.service_account_id
      max_budget         = litellm_service_account.production_api_service.max_budget
      spend              = litellm_service_account.production_api_service.spend
      models             = litellm_service_account.production_api_service.models
      created_at         = litellm_service_account.production_api_service.created_at
    }
    production_worker_service = {
      key_alias          = litellm_service_account.production_worker_service.key_alias
      team_id            = litellm_service_account.production_worker_service.team_id
      service_account_id = litellm_service_account.production_worker_service.service_account_id
      max_budget         = litellm_service_account.production_worker_service.max_budget
      spend              = litellm_service_account.production_worker_service.spend
      models             = litellm_service_account.production_worker_service.models
      created_at         = litellm_service_account.production_worker_service.created_at
    }
    staging_api_service = {
      key_alias          = litellm_service_account.staging_api_service.key_alias
      team_id            = litellm_service_account.staging_api_service.team_id
      service_account_id = litellm_service_account.staging_api_service.service_account_id
      max_budget         = litellm_service_account.staging_api_service.max_budget
      spend              = litellm_service_account.staging_api_service.spend
      models             = litellm_service_account.staging_api_service.models
      created_at         = litellm_service_account.staging_api_service.created_at
    }
    development_service = {
      key_alias          = litellm_service_account.development_service.key_alias
      team_id            = litellm_service_account.development_service.team_id
      service_account_id = litellm_service_account.development_service.service_account_id
      max_budget         = litellm_service_account.development_service.max_budget
      spend              = litellm_service_account.development_service.spend
      models             = litellm_service_account.development_service.models
      created_at         = litellm_service_account.development_service.created_at
    }
    minimal_service = {
      key_alias          = litellm_service_account.minimal_service.key_alias
      team_id            = litellm_service_account.minimal_service.team_id
      service_account_id = litellm_service_account.minimal_service.service_account_id
      max_budget         = litellm_service_account.minimal_service.max_budget
      spend              = litellm_service_account.minimal_service.spend
      models             = litellm_service_account.minimal_service.models
      created_at         = litellm_service_account.minimal_service.created_at
    }
  }
}

output "service_account_limits" {
  description = "Rate limits for all service accounts"
  value = {
    production_api_service = {
      tpm_limit = litellm_service_account.production_api_service.tpm_limit
      rpm_limit = litellm_service_account.production_api_service.rpm_limit
    }
    production_worker_service = {
      tpm_limit = litellm_service_account.production_worker_service.tpm_limit
      rpm_limit = litellm_service_account.production_worker_service.rpm_limit
    }
    staging_api_service = {
      tpm_limit = litellm_service_account.staging_api_service.tpm_limit
      rpm_limit = litellm_service_account.staging_api_service.rpm_limit
    }
    development_service = {
      tpm_limit = litellm_service_account.development_service.tpm_limit
      rpm_limit = litellm_service_account.development_service.rpm_limit
    }
    minimal_service = {
      tpm_limit = litellm_service_account.minimal_service.tpm_limit
      rpm_limit = litellm_service_account.minimal_service.rpm_limit
    }
  }
}

output "service_account_metadata" {
  description = "Metadata for all service accounts"
  value = {
    production_api_service    = litellm_service_account.production_api_service.metadata
    production_worker_service = litellm_service_account.production_worker_service.metadata
    staging_api_service       = litellm_service_account.staging_api_service.metadata
    development_service       = litellm_service_account.development_service.metadata
    minimal_service           = litellm_service_account.minimal_service.metadata
  }
}

output "test_summary" {
  description = "Summary of service account test configurations"
  value = {
    teams_created            = 3
    service_accounts_created = 5
    environments_tested      = ["production", "staging", "development"]
    service_types_tested     = ["api", "worker", "development"]
    features_tested = [
      "auto_generated_service_account_id",
      "custom_service_account_id",
      "team_association",
      "budget_limits",
      "rate_limits",
      "metadata_tagging",
      "model_restrictions",
      "multiple_teams"
    ]
    test_scenarios = [
      "production_multi_service",
      "staging_single_service",
      "development_minimal_config",
      "auto_generated_vs_custom_ids",
      "team_budget_inheritance"
    ]
  }
}
