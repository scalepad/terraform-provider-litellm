# litellm_model Resource

Manages a LiteLLM model configuration. This resource allows you to create, update, and delete model configurations in your LiteLLM instance.

## Example Usage

```hcl
resource "litellm_model" "gpt4" {
  model_name          = "gpt-4-proxy"
  custom_llm_provider = "openai"
  model_api_key       = var.openai_api_key
  model_api_base      = "https://api.openai.com/v1"
  api_version         = "2023-05-15"
  base_model          = "gpt-4"
  tier                = "paid"
  team_id             = "team-123"
  mode                = "chat"
  reasoning_effort    = "medium"
  thinking_enabled    = true
  thinking_budget_tokens = 1024
  merge_reasoning_content_in_choices = true
  tpm                 = 100000
  rpm                 = 1000

  # Cost configuration (per million tokens)
  input_cost_per_million_tokens  = 30.0    # $0.03 per 1k tokens = $30 per million
  output_cost_per_million_tokens = 60.0    # $0.06 per 1k tokens = $60 per million

  # AWS-specific configuration (if applicable)
  aws_access_key_id     = var.aws_access_key_id
  aws_secret_access_key = var.aws_secret_access_key
  aws_region_name       = var.aws_region
}
```

## Argument Reference

The following arguments are supported:

- `model_name` - (Required) The name of the model configuration. This will be used to identify the model in API calls.

- `custom_llm_provider` - (Required) The LLM provider for this model (e.g., "openai", "anthropic", "azure", "bedrock").

- `model_api_key` - (Optional) The API key for the underlying model provider.

- `model_api_base` - (Optional) The base URL for the model provider's API.

- `api_version` - (Optional) The API version to use for the model provider.

- `base_model` - (Required) The actual model identifier from the provider (e.g., "gpt-4", "claude-2").

- `tier` - (Optional) The usage tier for this model. Valid values are "free" or "paid". Default is "free".

- `team_id` - (Optional) Associate the model with a specific team.

- `mode` - (Optional) The intended use of the model. Valid values are:

  - `completion`
  - `embedding`
  - `image_generation`
  - `chat`
  - `moderation`
  - `audio_transcription`

- `tpm` - (Optional) Tokens per minute limit for this model.

- `rpm` - (Optional) Requests per minute limit for this model.

- `reasoning_effort` - (Optional) Configures the model's reasoning effort level. Valid values are:

  - `low`
  - `medium`
  - `high`

- `thinking_enabled` - (Optional) Enables the model's thinking capability. Default is `false`.

- `thinking_budget_tokens` - (Optional) Sets the token budget for the model's thinking capability. Default is `1024`.

- `merge_reasoning_content_in_choices` - (Optional) When set to `true`, merges reasoning content into the model's choices.

- `input_cost_per_million_tokens` - (Optional) Cost per million input tokens. This will be automatically converted to the per-token cost required by the API.

- `output_cost_per_million_tokens` - (Optional) Cost per million output tokens. This will be automatically converted to the per-token cost required by the API.

### AWS-specific Configuration

- `aws_access_key_id` - (Optional) AWS access key ID for AWS-based models.

- `aws_secret_access_key` - (Optional) AWS secret access key for AWS-based models.

- `aws_region_name` - (Optional) AWS region name for AWS-based models.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The ID of the model configuration.

## Import

Model configurations can be imported using the model ID:

```shell
terraform import litellm_model.gpt4 <model-id>
```

### Import Blocks (Terraform 1.5+)

For Terraform 1.5 and later, you can use import blocks to declaratively import models:

```hcl
import {
  to = litellm_model.gpt4
  id = "<model-id>"
}

resource "litellm_model" "gpt4" {
  # Configuration will be populated after import
  model_name          = "gpt-4-proxy"
  custom_llm_provider = "openai"
  base_model          = "gpt-4"
  # ... other configuration
}
```

Import blocks are the recommended approach for Terraform 1.5+ as they allow you to manage imports declaratively as part of your configuration.

Note: The model ID is generated when the model is created and is different from the `model_name`.

## Security Note

When using this resource, ensure that sensitive information such as API keys and AWS credentials are stored securely. It's recommended to use environment variables or a secure secret management solution rather than hardcoding these values in your Terraform configuration files.
