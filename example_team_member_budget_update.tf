# Example demonstrating the max_budget_in_team update fix
# This example shows how changing max_budget_in_team will now properly update all existing team members

terraform {
  required_providers {
    litellm = {
      source = "registry.terraform.io/ncecere/litellm"
    }
  }
}

provider "litellm" {
  api_base = "https://your-litellm-proxy.com"
  api_key  = var.litellm_api_key
}

# Create a team first
resource "litellm_team" "example" {
  team_alias = "example-team"
  max_budget = 500.0
  models     = ["gpt-4", "gpt-3.5-turbo"]
}

# Add team members with initial budget
resource "litellm_team_member_add" "example_members" {
  team_id = litellm_team.example.id
  
  # Initial budget of $100 per member
  max_budget_in_team = 100.0
  
  member {
    user_email = "user1@example.com"
    role       = "admin"
  }
  
  member {
    user_email = "user2@example.com"
    role       = "user"
  }
  
  member {
    user_id = "user123"
    role    = "user"
  }
}

# To test the fix:
# 1. Apply the above configuration
# 2. Change max_budget_in_team from 100.0 to 120.0
# 3. Run terraform plan - it should show the budget change
# 4. Run terraform apply - all existing members will be updated with the new budget

# Example of the change that will now work correctly:
# max_budget_in_team = 120.0  # Changed from 100.0

variable "litellm_api_key" {
  description = "API key for LiteLLM"
  type        = string
  sensitive   = true
}
