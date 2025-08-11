# Example demonstrating the improved team member update behavior
# This example shows how modifying a team member's role will now use
# the /team/member_update endpoint instead of delete/add

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

# Example team member resource
resource "litellm_team_member_add" "example" {
  team_id = "your-team-id"
  
  member {
    user_id = "user-123"
    role    = "admin"  # Change this to "user" to test the update behavior
  }

  member {
    user_email = "developer@example.com"
    role       = "user"  # Change this to "admin" to test the update behavior
  }

  max_budget_in_team = 100.0
}

# Output to show the resource ID
output "team_member_resource_id" {
  value = litellm_team_member_add.example.id
}
