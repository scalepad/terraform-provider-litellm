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

# Random strings to ensure unique key aliases
resource "random_string" "basic_suffix" {
  length  = 8
  special = false
  upper   = false
}

resource "random_string" "advanced_suffix" {
  length  = 8
  special = false
  upper   = false
}

resource "random_string" "restricted_suffix" {
  length  = 8
  special = false
  upper   = false
}

provider "litellm" {
  api_base = "http://localhost:4000"
  api_key  = "sk-test-master-key-12345"
}

# Test basic key creation
resource "litellm_key" "basic_key" {
  models = ["gpt-4"]

  max_budget      = 25.0
  budget_duration = "30d"
  tpm_limit       = 2500
  rpm_limit       = 25

  key_alias = "basic-test-key-${random_string.basic_suffix.result}"
  duration  = "7d"

  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "basic_key"
    created_by  = "terraform"
  }

  key_type = "llm_api"
}

# Test key with comprehensive configuration
resource "litellm_key" "advanced_key" {
  models = [
    "gpt-4",
    "gpt-3.5-turbo",
    "claude-3-sonnet"
  ]

  max_budget      = 100.0
  soft_budget     = 80.0
  budget_duration = "30d"
  tpm_limit       = 10000
  rpm_limit       = 100


  aliases = {
    "primary-model"   = "gpt-4"
    "secondary-model" = "claude-3-sonnet"
    "fast-model"      = "gpt-3.5-turbo"
  }


  key_alias = "advanced-test-key-${random_string.advanced_suffix.result}"
  duration  = "60d"
  blocked   = false


  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "advanced_key"
    created_by  = "terraform"
    tier        = "premium"
  }

  key_type = "management"
}

# Test restricted key
resource "litellm_key" "restricted_key" {
  models = ["gpt-3.5-turbo"]

  max_budget      = 10.0
  budget_duration = "30d"
  tpm_limit       = 1000
  rpm_limit       = 10


  aliases = {
    "restricted-model" = "gpt-3.5-turbo"
  }

  key_alias = "restricted-test-key-${random_string.restricted_suffix.result}"
  duration  = "7d"
  blocked   = false


  metadata = {
    test_type   = "integration"
    environment = "ci"
    test_case   = "restricted_key"
    created_by  = "terraform"
    tier        = "basic"
  }

  key_type = "read_only"
}

# Outputs for verification
output "basic_key_token" {
  description = "Token for the basic test key"
  value       = litellm_key.basic_key.key
  sensitive   = true
}

output "advanced_key_token" {
  description = "Token for the advanced test key"
  value       = litellm_key.advanced_key.key
  sensitive   = true
}

output "restricted_key_token" {
  description = "Token for the restricted test key"
  value       = litellm_key.restricted_key.key
  sensitive   = true
}

output "key_info" {
  description = "Information about all test keys"
  value = {
    basic_key = {
      key_alias       = litellm_key.basic_key.key_alias
      max_budget      = litellm_key.basic_key.max_budget
      budget_duration = litellm_key.basic_key.budget_duration
      spend           = litellm_key.basic_key.spend
    }
    advanced_key = {
      key_alias       = litellm_key.advanced_key.key_alias
      max_budget      = litellm_key.advanced_key.max_budget
      soft_budget     = litellm_key.advanced_key.soft_budget
      budget_duration = litellm_key.advanced_key.budget_duration
      spend           = litellm_key.advanced_key.spend
    }
    restricted_key = {
      key_alias       = litellm_key.restricted_key.key_alias
      max_budget      = litellm_key.restricted_key.max_budget
      budget_duration = litellm_key.restricted_key.budget_duration
      spend           = litellm_key.restricted_key.spend
    }
  }
}

output "key_limits" {
  description = "Rate limits for all test keys"
  value = {
    basic_key = {
      tpm_limit = litellm_key.basic_key.tpm_limit
      rpm_limit = litellm_key.basic_key.rpm_limit
    }
    advanced_key = {
      tpm_limit             = litellm_key.advanced_key.tpm_limit
      rpm_limit             = litellm_key.advanced_key.rpm_limit
      max_parallel_requests = litellm_key.advanced_key.max_parallel_requests
    }
    restricted_key = {
      tpm_limit             = litellm_key.restricted_key.tpm_limit
      rpm_limit             = litellm_key.restricted_key.rpm_limit
      max_parallel_requests = litellm_key.restricted_key.max_parallel_requests
    }
  }
}
