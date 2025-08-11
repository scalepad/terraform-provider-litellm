# Example Terraform configuration for LiteLLM MCP Server resource

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

# Basic HTTP MCP Server
resource "litellm_mcp_server" "github_server" {
  server_name = "github-mcp-server"
  alias       = "github"
  description = "GitHub MCP server for repository operations"
  url         = "https://api.github.com/mcp"
  transport   = "http"
  auth_type   = "bearer"
  
  mcp_access_groups = ["dev_team", "devops_team"]
}

# SSE MCP Server with detailed configuration and cost tracking
resource "litellm_mcp_server" "zapier_server" {
  server_name  = "zapier-automation"
  alias        = "zapier"
  description  = "Zapier MCP server for workflow automation"
  url          = "https://actions.zapier.com/mcp/sk-xxxxx/sse"
  transport    = "sse"
  auth_type    = "bearer"
  spec_version = "2024-11-05"
  
  mcp_access_groups = ["automation_team", "marketing_team"]
  
  mcp_info {
    server_name = "Zapier Integration Server"
    description = "Provides automation tools through Zapier's MCP interface"
    logo_url    = "https://zapier.com/assets/images/zapier-logo.png"
    
    mcp_server_cost_info {
      default_cost_per_query = 0.01
      
      tool_name_to_cost_per_query = {
        "send_email"           = 0.05
        "create_document"      = 0.03
        "update_spreadsheet"   = 0.02
        "post_to_slack"        = 0.01
        "create_calendar_event" = 0.04
      }
    }
  }
}

# Stdio MCP Server for local development
resource "litellm_mcp_server" "local_dev_server" {
  server_name = "local-development-tools"
  alias       = "local-dev"
  description = "Local MCP server for development tools"
  url         = "stdio://local-dev"
  transport   = "stdio"
  auth_type   = "none"
  
  command = "python3"
  args    = ["/opt/mcp-servers/dev-tools/server.py", "--verbose"]
  
  env = {
    "PYTHONPATH"    = "/opt/mcp-servers/dev-tools"
    "DEBUG"         = "true"
    "LOG_LEVEL"     = "info"
    "WORKSPACE_DIR" = "/workspace"
  }
  
  mcp_access_groups = ["local_developers"]
  
  mcp_info {
    server_name = "Development Tools"
    description = "Local development utilities and tools"
    
    mcp_server_cost_info {
      default_cost_per_query = 0.0  # Free for local development
    }
  }
}

# Enterprise MCP Server with comprehensive configuration
resource "litellm_mcp_server" "enterprise_api_server" {
  server_name  = "enterprise-api-gateway"
  alias        = "enterprise"
  description  = "Enterprise API gateway MCP server"
  url          = "https://api.enterprise.com/mcp/v1"
  transport    = "http"
  auth_type    = "bearer"
  spec_version = "2024-11-05"
  
  mcp_access_groups = [
    "enterprise_users",
    "api_consumers",
    "integration_team"
  ]
  
  mcp_info {
    server_name = "Enterprise API Gateway"
    description = "Provides access to enterprise APIs and services"
    logo_url    = "https://enterprise.com/logo.png"
    
    mcp_server_cost_info {
      default_cost_per_query = 0.10
      
      tool_name_to_cost_per_query = {
        "query_database"       = 0.25
        "generate_report"      = 0.50
        "send_notification"    = 0.05
        "create_user"          = 0.15
        "update_permissions"   = 0.20
        "audit_log_query"      = 0.30
      }
    }
  }
}

# Output the server IDs for reference
output "mcp_server_ids" {
  description = "IDs of the created MCP servers"
  value = {
    github     = litellm_mcp_server.github_server.server_id
    zapier     = litellm_mcp_server.zapier_server.server_id
    local_dev  = litellm_mcp_server.local_dev_server.server_id
    enterprise = litellm_mcp_server.enterprise_api_server.server_id
  }
}

# Output server status information
output "mcp_server_status" {
  description = "Status information for MCP servers"
  value = {
    github = {
      status              = litellm_mcp_server.github_server.status
      last_health_check   = litellm_mcp_server.github_server.last_health_check
      health_check_error  = litellm_mcp_server.github_server.health_check_error
    }
    zapier = {
      status              = litellm_mcp_server.zapier_server.status
      last_health_check   = litellm_mcp_server.zapier_server.last_health_check
      health_check_error  = litellm_mcp_server.zapier_server.health_check_error
    }
    local_dev = {
      status              = litellm_mcp_server.local_dev_server.status
      last_health_check   = litellm_mcp_server.local_dev_server.last_health_check
      health_check_error  = litellm_mcp_server.local_dev_server.health_check_error
    }
    enterprise = {
      status              = litellm_mcp_server.enterprise_api_server.status
      last_health_check   = litellm_mcp_server.enterprise_api_server.last_health_check
      health_check_error  = litellm_mcp_server.enterprise_api_server.health_check_error
    }
  }
}

# Variable for API key
variable "litellm_api_key" {
  description = "API key for LiteLLM"
  type        = string
  sensitive   = true
}
