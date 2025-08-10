# Example LiteLLM Credential Configuration

# OpenAI credential for GPT models
resource "litellm_credential" "openai_cred" {
  credential_name = "openai-api-key"
  model_id        = "gpt-4"
  
  credential_info = {
    provider = "openai"
    region   = "us-east-1"
    purpose  = "chat-completions"
  }
  
  credential_values = {
    api_key = var.openai_api_key
    org_id  = var.openai_org_id
  }
}

# Pinecone credential for vector store
resource "litellm_credential" "pinecone_cred" {
  credential_name = "pinecone-production"
  
  credential_info = {
    provider    = "pinecone"
    environment = "production"
    region      = "us-east-1"
  }
  
  credential_values = {
    api_key    = var.pinecone_api_key
    index_name = "document-embeddings"
  }
}

# Anthropic credential for Claude models
resource "litellm_credential" "anthropic_cred" {
  credential_name = "anthropic-api-key"
  
  credential_info = {
    provider = "anthropic"
    purpose  = "text-generation"
  }
  
  credential_values = {
    api_key = var.anthropic_api_key
  }
}

# Variables for sensitive values
variable "openai_api_key" {
  description = "OpenAI API key"
  type        = string
  sensitive   = true
}

variable "openai_org_id" {
  description = "OpenAI organization ID"
  type        = string
  sensitive   = true
}

variable "pinecone_api_key" {
  description = "Pinecone API key"
  type        = string
  sensitive   = true
}

variable "anthropic_api_key" {
  description = "Anthropic API key"
  type        = string
  sensitive   = true
}
