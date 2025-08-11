# Example demonstrating the restored team_member_permissions functionality
# This example shows how to create a team with specific member permissions

terraform {
  required_providers {
    litellm = {
      source = "local/litellm"
    }
  }
}

provider "litellm" {
  api_base = "https://ai.bitop.dev"
  api_key  = var.litellm_api_key
}

variable "litellm_api_key" {
  description = "LiteLLM API Key"
  type        = string
  sensitive   = true
}

# Example team with member permissions
resource "litellm_team" "example" {
  team_alias = "example-team"
  
  # Team configuration
  max_budget      = 100.0
  budget_duration = "1mo"
  tpm_limit       = 1000
  rpm_limit       = 100
  
  # Models this team can access
  models = [
    "gpt-4",
    "gpt-3.5-turbo"
  ]
  
  # Team member permissions - this functionality has been restored
  team_member_permissions = [
    "create_key",
    "delete_key",
    "view_spend",
    "edit_team"
  ]
  
  # Additional metadata
  metadata = {
    department = "engineering"
    project    = "ai-integration"
  }
}

# Output to show the team ID and permissions
output "team_id" {
  value = litellm_team.example.id
}

output "team_permissions" {
  value = litellm_team.example.team_member_permissions
}
