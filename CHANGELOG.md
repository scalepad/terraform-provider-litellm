# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.4.1] - 2025-08-11

### :sparkles: New Features

- [`7d5f70a`](https://github.com/scalepad/terraform-provider-litellm/commit/7d5f70a466118602031a0a2b0300e697252eb19d) - add changelog summary to release workflow _(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))_

### :bug: Bug Fixes

- [`19b7bd1`](https://github.com/scalepad/terraform-provider-litellm/commit/19b7bd1653fefdb14f71e9a11d3ee7fb58c8d8df) - update documentation _(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))_
- [`3e10ef3`](https://github.com/scalepad/terraform-provider-litellm/commit/3e10ef39b180531913c128fa246f6af939e4051c) - update GITHUB*TOKEN reference in release workflow *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))\_

## [0.3.11] - 2025-08-10

### Added

- **New Resource**: `litellm_credential` - Manage credentials for secure authentication
  - Support for storing sensitive credential values (API keys, tokens, etc.)
  - Non-sensitive credential information storage
  - Model ID association for credentials
  - Secure handling of sensitive data with Terraform's sensitive attribute
- **New Resource**: `litellm_vector_store` - Manage vector stores for embeddings and RAG
  - Support for multiple vector store providers (Pinecone, Weaviate, Chroma, Qdrant, etc.)
  - Integration with credential management for secure authentication
  - Configurable metadata and provider-specific parameters
  - Full CRUD operations for vector store lifecycle management
- **New Data Source**: `litellm_credential` - Retrieve information about existing credentials
  - Read-only access to credential metadata (sensitive values excluded for security)
  - Support for model ID filtering
  - Cross-stack and cross-configuration referencing capabilities
- **New Data Source**: `litellm_vector_store` - Retrieve information about existing vector stores
  - Complete vector store information retrieval
  - Support for monitoring, validation, and cross-referencing use cases
  - Metadata-based conditional logic support
- Enhanced API response handling for credential and vector store operations
- Comprehensive documentation and examples for new resources and data sources
- Example Terraform configurations for common use cases

### Changed

- Extended `utils.go` with specialized API response handlers for credentials and vector stores
- Updated provider configuration to include new resources and data sources
- Enhanced error handling for credential and vector store not found scenarios

## [0.3.10] - 2025-08-10

### Added

- **New Resource**: `litellm_mcp_server` - Manage MCP (Model Context Protocol) servers
  - Support for HTTP, SSE, and stdio transport types
  - Configurable authentication types (none, bearer, basic)
  - MCP access groups for permission management
  - Cost tracking configuration for MCP tools
  - Environment variables and command arguments for stdio transport
  - Health check status monitoring
  - Comprehensive documentation and examples

### Changed

- Updated provider to support MCP server management functionality
- Enhanced API response handling for MCP-specific operations

## [0.3.9] - 2025-08-10

### Fixed

- Fixed issue where omitting `budget_duration` in key resource caused API error "Invalid duration format"
- Added missing `omitempty` JSON tag to `BudgetDuration` field in Key struct to prevent sending empty strings to API

## [0.3.8] - 2025-08-08

### Added

- Added `additional_litellm_params` field to model resource for custom parameters beyond standard ones
- Support for passing custom parameters like `drop_params`, `timeout`, `max_retries`, `organization`, etc.
- Automatic type conversion for string values to appropriate types (boolean, integer, float)
- Full backward compatibility with existing model configurations
- Comprehensive example demonstrating various use cases with different providers

## [0.3.7] - 2025-08-08

### Fixed

- Fixed issue where changing max_budget_in_team didn't update existing team members with new budget
- Added budget change detection using d.HasChange to update ALL existing members when budget changes
- Implemented tracking to avoid duplicate API calls for members already updated
- Enhanced debug logging for budget update operations

## [0.3.6] - 2025-08-08

### Fixed

- Fixed issue where models deleted from LiteLLM proxy caused terraform plan to fail instead of planning recreation
- Enhanced ErrorResponse struct to properly parse LiteLLM proxy error format with Detail field
- Improved isModelNotFoundError function to detect "not found on litellm proxy" messages in Detail.Error field

## [0.3.5] - 2025-08-08

### Fixed

- Fixed team member update behavior to use member_update endpoint instead of delete/re-add
- Restored team_member_permissions functionality to litellm_team resource
- Enhanced team resource with proper permissions management endpoints

## [0.3.0] - 2025-04-23

### Fixed

- Implemented retry mechanism with exponential backoff for model read operations
- Added detailed logging for retry attempts
- Improved error handling for "model not found" errors

## [0.2.9] - 2025-04-23

### Fixed

- Increased delay after model creation from 2 to 5 seconds to fix "model not found" errors
- Added logging to confirm delay is working properly

## [0.2.8] - 2025-04-23

### Fixed

- Added delay after model creation to fix "model not found" errors when the LiteLLM proxy hasn't fully registered the model yet

## [0.2.7] - 2025-04-23

### Fixed

- Fixed issue where `thinking_enabled` and `merge_reasoning_content_in_choices` values were not being preserved in state, causing Terraform to want to modify them on every run

## [0.2.6] - 2025-03-13

### Added

- Added new `merge_reasoning_content_in_choices` option to model resource

## [0.2.5] - 2025-03-13

### Fixed

- Fixed issue where `thinking_budget_tokens` was being added to models that don't have `thinking_enabled = true`

## [0.2.4] - 2025-03-13

### Added

- Added new `thinking` capability to model resource with configurable parameters:
  - `thinking_enabled` - Boolean to enable/disable thinking capability (default: false)
  - `thinking_budget_tokens` - Integer to set token budget for thinking (default: 1024)

## [0.2.2] - 2025-02-06

### Added

- Added new `reasoning_effort` parameter to model resource with values: "low", "medium", "high"
- Added "chat" mode to model resource

### Changed

- Updated model mode options to: "completion", "embedding", "image_generation", "chat", "moderation", "audio_transcription"

## [1.0.0] - 2024-01-17

### Added

- Initial release of the LiteLLM Terraform Provider
- Support for managing LiteLLM models
- Support for managing teams and team members
- Comprehensive documentation for all resources
  [v0.4.1]: https://github.com/scalepad/terraform-provider-litellm/compare/v0.4.0...v0.4.1
