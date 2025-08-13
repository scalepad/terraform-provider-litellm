# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.6.1] - 2025-08-13
### :bug: Bug Fixes
- [`c1f9e1e`](https://github.com/scalepad/terraform-provider-litellm/commit/c1f9e1eb55bb99c22a102ca09af839680b3fb1e8) - remove deprecated role options from team member resource documentation *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*


## [v0.6.0] - 2025-08-13
### :sparkles: New Features
- [`fb6e2eb`](https://github.com/scalepad/terraform-provider-litellm/commit/fb6e2eb9ce520fb790877dc6fa50fba0853de3c6) - add connection check *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`26627f2`](https://github.com/scalepad/terraform-provider-litellm/commit/26627f25d9cf387057758264a1badfe60871eb9a) - **user**: add user management to litellm *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`30189ba`](https://github.com/scalepad/terraform-provider-litellm/commit/30189babf511415e4c91c775d3cf007795827c3c) - **users**: make user module able to create and update existing users. *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*

### :bug: Bug Fixes
- [`deebbec`](https://github.com/scalepad/terraform-provider-litellm/commit/deebbec8f55a5a83d7053744215ca4c8a4445e3f) - **key**: fix issue with setting send_email for key *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`69508e9`](https://github.com/scalepad/terraform-provider-litellm/commit/69508e9d8d64bb80f1ca4870ed514ee64a421bdf) - **credentials**: fix the datasource credentials to not use deprecated methods *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`401fe4b`](https://github.com/scalepad/terraform-provider-litellm/commit/401fe4b2ce1987f78e34a2ace5559e85c2f73886) - **credentials**: fix operation on resource credential *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`363ca7b`](https://github.com/scalepad/terraform-provider-litellm/commit/363ca7b4cd3907d2a3ee32fac6e93e232067a4e0) - **models**: refactor the code to split the logic into proper files *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`1cfd062`](https://github.com/scalepad/terraform-provider-litellm/commit/1cfd062ce586b93fb89f86eeacfa036aac7ee010) - **team**: refactor how team are created and handled *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`b0f5043`](https://github.com/scalepad/terraform-provider-litellm/commit/b0f5043adc586429296bfc50e39cac8da9b871a8) - **team_member**: Refactor fully the team member to be consistent *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`b6c1771`](https://github.com/scalepad/terraform-provider-litellm/commit/b6c1771540878e30b1705727c1ab175b19cdf355) - **credentials**: fix issue with creds *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`164db5f`](https://github.com/scalepad/terraform-provider-litellm/commit/164db5fef35626cf3a6cdf040e24143efca816be) - **mcp**: refactor mcp and add tests *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`9425c0c`](https://github.com/scalepad/terraform-provider-litellm/commit/9425c0ccf9fae1396b7dfbbeda3b1ff31c83f54d) - **vector**: refactor vector *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`c51de99`](https://github.com/scalepad/terraform-provider-litellm/commit/c51de9940fa40371cac27d0dfd02745743d7bdc8) - **user**: fix user type and deleting users *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`5a5ce96`](https://github.com/scalepad/terraform-provider-litellm/commit/5a5ce961689183f4df3fcaf2e905d42ca71379aa) - **user**: issue with new fields not being propagated to state *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`fcb1f70`](https://github.com/scalepad/terraform-provider-litellm/commit/fcb1f700a795a613d7464490dcbaa8b3f4870b12) - **models**: be sure to have validation for tier of model *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`be36535`](https://github.com/scalepad/terraform-provider-litellm/commit/be365356bb4397844283bef197d63906d2b81335) - **key**: add proper duration validation *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`5136d3e`](https://github.com/scalepad/terraform-provider-litellm/commit/5136d3ef83f43b9473a043197924ef85422eb290) - **key**: fix the key module *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`ad507ab`](https://github.com/scalepad/terraform-provider-litellm/commit/ad507ab047a1c6fc251968218a810983ceb7b9ba) - **key**: keys are now only updating fields that have changed. *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`0260291`](https://github.com/scalepad/terraform-provider-litellm/commit/0260291016a8165f1438a6b0101ffe61aa2c80b0) - **key**: For security reason, migrate from using the key as ID to use the token_id *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`4381ff7`](https://github.com/scalepad/terraform-provider-litellm/commit/4381ff759575fe27067be9693eea778e555631c7) - **team**: only send the field that changed *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`e50c77e`](https://github.com/scalepad/terraform-provider-litellm/commit/e50c77ebccf7af59c5656dcfd2b3a5af05d2bc95) - **vector**: have proper way to set them *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`9e9351f`](https://github.com/scalepad/terraform-provider-litellm/commit/9e9351fcbf9b6e684e16d8efa6f9455acd46059c) - **user**: be sure to url-encode the ids *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`8134087`](https://github.com/scalepad/terraform-provider-litellm/commit/813408719397bd6c34aa5f5faf010bee3f3bf985) - add URL escaping for model ID parameter in resource_model_crud.go *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`932ea48`](https://github.com/scalepad/terraform-provider-litellm/commit/932ea489237d665878f896006abb018c7558547c) - add URL escaping for key ID and key alias parameters in key_crud.go *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`c8dd17c`](https://github.com/scalepad/terraform-provider-litellm/commit/c8dd17ccbec6b0454a8665ba3077750f6d33c82b) - add URL escaping for team ID parameters in team_crud.go *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`fa7d88c`](https://github.com/scalepad/terraform-provider-litellm/commit/fa7d88c8870c7efca548cba4b39a7fcb91174599) - update resourceKeyCreate to set ID using TokenID instead of Key *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*


## [v0.5.2] - 2025-08-11
### :bug: Bug Fixes
- [`c10484b`](https://github.com/scalepad/terraform-provider-litellm/commit/c10484b3f5060a08852481350d0923a90beeaa4b) - don't report on null field *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*


## [v0.5.1] - 2025-08-11
### :sparkles: New Features
- [`c2e1a2d`](https://github.com/scalepad/terraform-provider-litellm/commit/c2e1a2d7b9886d8d1afda59cdd1380a3c966b770) - add team member buget *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*

### :wrench: Chores
- [`ea054fb`](https://github.com/scalepad/terraform-provider-litellm/commit/ea054fbb806913f86c2b2a4295e20eb29de9ed44) - update CHANGELOG.md for release v0.5.0 [skip ci] *(commit by [@github-actions[bot]](https://github.com/apps/github-actions))*


## [v0.5.0] - 2025-08-11
### :sparkles: New Features
- [`bbd0888`](https://github.com/scalepad/terraform-provider-litellm/commit/bbd088896ef7045c45f567ab3c4b3ee0fa5a3e1c) - add additional_litellm_params support for custom model parameters
- [`82a5500`](https://github.com/scalepad/terraform-provider-litellm/commit/82a550013e4aac422ceac724a6980f4d50c9ff09) - Add MCP server resource support
- [`0ed4066`](https://github.com/scalepad/terraform-provider-litellm/commit/0ed4066649751a5787a37ecf281df1eff94d2bc1) - Add credential and vector store resources with data sources

### :bug: Bug Fixes
- [`98218d0`](https://github.com/scalepad/terraform-provider-litellm/commit/98218d05e1f0ab7593ca1ba436ef57168ed90ade) - handle model not found error and recreate
- [`59e43a1`](https://github.com/scalepad/terraform-provider-litellm/commit/59e43a10a82aa8cd56c499d9e0871c243864b695) - handle max_budget_in_team updates for existing team members

### :wrench: Chores
- [`6417bd5`](https://github.com/scalepad/terraform-provider-litellm/commit/6417bd59bb6351d44004b6af60555b0f7100ad99) - update CHANGELOG.md for release v0.4.1 [skip ci] *(commit by [@github-actions[bot]](https://github.com/apps/github-actions))*


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
[v0.5.0]: https://github.com/scalepad/terraform-provider-litellm/compare/v0.4.1...v0.5.0
[v0.5.1]: https://github.com/scalepad/terraform-provider-litellm/compare/v0.5.0...v0.5.1
[v0.5.2]: https://github.com/scalepad/terraform-provider-litellm/compare/v0.5.1...v0.5.2
[v0.6.0]: https://github.com/scalepad/terraform-provider-litellm/compare/v0.5.2...v0.6.0
[v0.6.1]: https://github.com/scalepad/terraform-provider-litellm/compare/v0.6.0...v0.6.1
