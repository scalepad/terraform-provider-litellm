# Service Accounts Integration Tests

This directory contains comprehensive integration tests for the `litellm_service_account` Terraform resource.

## Overview

These tests create multiple teams and service accounts to validate the service account functionality across different scenarios:

- **Multi-environment testing**: Production, staging, and development teams
- **Service account variations**: Different configurations and use cases
- **Team association**: Service accounts properly linked to teams
- **Auto-generated vs custom IDs**: Testing both automatic UUID generation and custom service account IDs

## Test Structure

### Teams Created

1. **Production Team** - High limits, multiple models, 30-day budget cycle
2. **Staging Team** - Medium limits, subset of models, 7-day budget cycle
3. **Development Team** - Low limits, single model, daily budget cycle

### Service Accounts Created

1. **Production API Service** - Full team access, custom service account ID
2. **Production Worker Service** - Restricted models, auto-generated ID
3. **Staging API Service** - Team access with staging limits
4. **Development Service** - Minimal configuration for development
5. **Minimal Service** - Basic setup with auto-generated ID

## Running the Tests

### Prerequisites

- LiteLLM server running on `http://localhost:4000`
- Valid API key: `sk-test-master-key-12345`
- Terraform installed

### Execute Tests

```bash
cd tests/service-accounts
terraform init
terraform plan
terraform apply
```

### Verify Results

```bash
terraform output service_account_keys    # View generated keys (sensitive)
terraform output service_account_info    # View service account details
terraform output team_ids               # View created team IDs
terraform output test_summary           # View test coverage summary
```

## Test Scenarios Covered

### ✅ Core Functionality

- Service account creation with team association
- Automatic UUIDv7 generation for service account IDs
- Custom service account ID assignment
- Proper metadata handling and tagging

### ✅ Configuration Variations

- Different budget limits and durations
- Rate limiting (TPM/RPM) configurations
- Model access restrictions
- Team inheritance of permissions

### ✅ Multi-Environment Setup

- Production-grade configurations
- Staging environment testing
- Development sandbox testing
- Environment-specific metadata

### ✅ Integration Testing

- Team-service account relationships
- Budget inheritance and limits
- Rate limit enforcement
- Metadata propagation

## Outputs

The test provides comprehensive outputs for validation:

- **service_account_keys**: Generated API keys (marked sensitive)
- **service_account_info**: Detailed service account information
- **service_account_limits**: Rate limiting configurations
- **service_account_metadata**: Metadata validation
- **team_ids**: Created team identifiers
- **test_summary**: Coverage and scenario summary

## Cleanup

To clean up test resources:

```bash
terraform destroy
```

## Notes

- All resources are tagged with test metadata for easy identification
- Random suffixes ensure unique resource names across test runs
- Service accounts are automatically associated with their respective teams
- Budget and rate limits are configured appropriately for each environment
