# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.4.1] - 2025-08-11
### :sparkles: New Features
- [`7d5f70a`](https://github.com/scalepad/terraform-provider-litellm/commit/7d5f70a466118602031a0a2b0300e697252eb19d) - add changelog summary to release workflow *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*

### :bug: Bug Fixes
- [`19b7bd1`](https://github.com/scalepad/terraform-provider-litellm/commit/19b7bd1653fefdb14f71e9a11d3ee7fb58c8d8df) - update documentation *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*
- [`3e10ef3`](https://github.com/scalepad/terraform-provider-litellm/commit/3e10ef39b180531913c128fa246f6af939e4051c) - update GITHUB_TOKEN reference in release workflow *(commit by [@sp-aaflalo](https://github.com/sp-aaflalo))*


## [0.3.3] - 2025-05-10

### Fixed
- Fixed issue where manual changes to team member permissions outside of Terraform weren't detected during plan/apply

## [0.3.2] - 2025-05-09

### Fixed
- Fixed issue where team member permissions weren't being applied to existing teams

## [0.3.1] - 2025-05-09

### Added
- Support for team member permissions in the `litellm_team` resource

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
