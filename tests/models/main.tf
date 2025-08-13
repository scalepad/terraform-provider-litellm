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

# Test basic model configuration
resource "litellm_model" "basic_model" {
  model_name          = "test-gpt-3.5-turbo"
  custom_llm_provider = "openai"
  base_model          = "gpt-3.5-turbo"

  mode = "chat"
  tier = "free"

  input_cost_per_million_tokens  = 1500.0
  output_cost_per_million_tokens = 2000.0

  tpm = 10000
  rpm = 100
}

# Test advanced model configuration
resource "litellm_model" "advanced_model" {
  model_name          = "test-gpt-4-advanced"
  custom_llm_provider = "openai"
  base_model          = "gpt-4"

  mode = "chat"
  tier = "paid"

  input_cost_per_million_tokens  = 30000.0
  output_cost_per_million_tokens = 60000.0

  tpm = 50000
  rpm = 500

  reasoning_effort                   = "medium"
  thinking_enabled                   = true
  thinking_budget_tokens             = 2048
  merge_reasoning_content_in_choices = true

  additional_litellm_params = {
    "temperature" = "0.7"
    "max_tokens"  = "8192"
  }
}

# Test Claude model configuration
resource "litellm_model" "claude_model" {
  model_name          = "test-claude-3-sonnet"
  custom_llm_provider = "anthropic"
  base_model          = "claude-3-sonnet-20240229"

  mode = "chat"
  tier = "paid"

  input_cost_per_million_tokens  = 3000.0
  output_cost_per_million_tokens = 15000.0

  tpm = 25000
  rpm = 250

  additional_litellm_params = {
    "temperature" = "0.5"
    "max_tokens"  = "4096"
  }
}

# Test model with AWS configuration
resource "litellm_model" "aws_model" {
  model_name          = "test-bedrock-claude"
  custom_llm_provider = "bedrock"
  base_model          = "anthropic.claude-3-sonnet-20240229-v1:0"

  mode = "chat"
  tier = "paid"

  input_cost_per_million_tokens  = 3000.0
  output_cost_per_million_tokens = 15000.0

  tpm = 30000
  rpm = 300

  aws_region_name = "us-east-1"

  additional_litellm_params = {
    "temperature" = "0.3"
    "max_tokens"  = "2048"
  }
}

# Outputs for verification
output "basic_model_id" {
  description = "ID of the basic test model"
  value       = litellm_model.basic_model.id
}

output "advanced_model_id" {
  description = "ID of the advanced test model"
  value       = litellm_model.advanced_model.id
}

output "claude_model_id" {
  description = "ID of the Claude test model"
  value       = litellm_model.claude_model.id
}

output "aws_model_id" {
  description = "ID of the AWS test model"
  value       = litellm_model.aws_model.id
}

output "model_info" {
  description = "Information about all test models"
  value = {
    basic_model = {
      model_name          = litellm_model.basic_model.model_name
      custom_llm_provider = litellm_model.basic_model.custom_llm_provider
      base_model          = litellm_model.basic_model.base_model
      mode                = litellm_model.basic_model.mode
      tier                = litellm_model.basic_model.tier
    }
    advanced_model = {
      model_name          = litellm_model.advanced_model.model_name
      custom_llm_provider = litellm_model.advanced_model.custom_llm_provider
      base_model          = litellm_model.advanced_model.base_model
      mode                = litellm_model.advanced_model.mode
      tier                = litellm_model.advanced_model.tier
      reasoning_effort    = litellm_model.advanced_model.reasoning_effort
      thinking_enabled    = litellm_model.advanced_model.thinking_enabled
    }
    claude_model = {
      model_name          = litellm_model.claude_model.model_name
      custom_llm_provider = litellm_model.claude_model.custom_llm_provider
      base_model          = litellm_model.claude_model.base_model
      mode                = litellm_model.claude_model.mode
      tier                = litellm_model.claude_model.tier
    }
    aws_model = {
      model_name          = litellm_model.aws_model.model_name
      custom_llm_provider = litellm_model.aws_model.custom_llm_provider
      base_model          = litellm_model.aws_model.base_model
      mode                = litellm_model.aws_model.mode
      tier                = litellm_model.aws_model.tier
      aws_region_name     = litellm_model.aws_model.aws_region_name
    }
  }
}

output "model_costs" {
  description = "Cost information for all test models"
  value = {
    basic_model = {
      input_cost_per_million_tokens  = litellm_model.basic_model.input_cost_per_million_tokens
      output_cost_per_million_tokens = litellm_model.basic_model.output_cost_per_million_tokens
      tpm                            = litellm_model.basic_model.tpm
      rpm                            = litellm_model.basic_model.rpm
    }
    advanced_model = {
      input_cost_per_million_tokens  = litellm_model.advanced_model.input_cost_per_million_tokens
      output_cost_per_million_tokens = litellm_model.advanced_model.output_cost_per_million_tokens
      tpm                            = litellm_model.advanced_model.tpm
      rpm                            = litellm_model.advanced_model.rpm
    }
    claude_model = {
      input_cost_per_million_tokens  = litellm_model.claude_model.input_cost_per_million_tokens
      output_cost_per_million_tokens = litellm_model.claude_model.output_cost_per_million_tokens
      tpm                            = litellm_model.claude_model.tpm
      rpm                            = litellm_model.claude_model.rpm
    }
    aws_model = {
      input_cost_per_million_tokens  = litellm_model.aws_model.input_cost_per_million_tokens
      output_cost_per_million_tokens = litellm_model.aws_model.output_cost_per_million_tokens
      tpm                            = litellm_model.aws_model.tpm
      rpm                            = litellm_model.aws_model.rpm
    }
  }
}
