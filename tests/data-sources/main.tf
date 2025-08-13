terraform {
  required_providers {
    litellm = {
      source  = "registry.terraform.io/local/litellm"
      version = "1.0.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.1"
    }
  }
}

# Generate random suffix for unique naming
resource "random_id" "test_suffix" {
  byte_length = 4
}

provider "litellm" {
  api_base = "http://localhost:4000"
  api_key  = "sk-test-master-key-12345"
}

# Create a credential first for testing data source
resource "litellm_credential" "test_credential" {
  credential_name = "test-openai-credential-${random_id.test_suffix.hex}"

  credential_info = {
    provider    = "openai"
    environment = "test"
    purpose     = "data_source_testing"
  }

  credential_values = {
    api_key = "sk-test-openai-key-for-datasource"
  }
}

# Create a vector store first for testing data source
resource "litellm_vector_store" "test_vector_store" {
  vector_store_name   = "test-integration-vector-store-${random_id.test_suffix.hex}"
  custom_llm_provider = "openai"

  vector_store_description = "Test vector store for data source integration tests"

  litellm_credential_name = litellm_credential.test_credential.credential_name

  litellm_params = {
    model = "text-embedding-ada-002"
  }

  vector_store_metadata = {
    test_type   = "integration"
    environment = "ci"
    purpose     = "data_source_testing"
  }
}

# Test credential data source
data "litellm_credential" "test_cred_data" {
  credential_name = litellm_credential.test_credential.credential_name

  depends_on = [litellm_credential.test_credential]
}

# Test vector store data source
data "litellm_vector_store" "test_vector_data" {
  vector_store_id = litellm_vector_store.test_vector_store.vector_store_id

  depends_on = [litellm_vector_store.test_vector_store]
}

# Outputs for verification
output "credential_data_source" {
  description = "Information from credential data source"
  value = {
    credential_name = data.litellm_credential.test_cred_data.credential_name
    credential_info = data.litellm_credential.test_cred_data.credential_info
    # Note: credential_values are not exposed in data sources for security
  }
  sensitive = true
}

output "vector_store_data_source" {
  description = "Information from vector store data source"
  value = {
    vector_store_id          = data.litellm_vector_store.test_vector_data.vector_store_id
    vector_store_name        = data.litellm_vector_store.test_vector_data.vector_store_name
    custom_llm_provider      = data.litellm_vector_store.test_vector_data.custom_llm_provider
    vector_store_description = data.litellm_vector_store.test_vector_data.vector_store_description
    litellm_credential_name  = data.litellm_vector_store.test_vector_data.litellm_credential_name
    vector_store_metadata    = data.litellm_vector_store.test_vector_data.vector_store_metadata
    litellm_params           = data.litellm_vector_store.test_vector_data.litellm_params
    created_at               = data.litellm_vector_store.test_vector_data.created_at
    updated_at               = data.litellm_vector_store.test_vector_data.updated_at
  }
}

# Data source comparison and validation
output "data_source_validation" {
  description = "Validation that data sources return correct information"
  value = {
    # Credential validation
    credential_name_matches = (
      litellm_credential.test_credential.credential_name ==
      data.litellm_credential.test_cred_data.credential_name
    )

    # Vector store validation
    vector_store_id_matches = (
      litellm_vector_store.test_vector_store.vector_store_id ==
      data.litellm_vector_store.test_vector_data.vector_store_id
    )
    vector_store_name_matches = (
      litellm_vector_store.test_vector_store.vector_store_name ==
      data.litellm_vector_store.test_vector_data.vector_store_name
    )
    custom_llm_provider_matches = (
      litellm_vector_store.test_vector_store.custom_llm_provider ==
      data.litellm_vector_store.test_vector_data.custom_llm_provider
    )
    credential_reference_matches = (
      litellm_vector_store.test_vector_store.litellm_credential_name ==
      data.litellm_vector_store.test_vector_data.litellm_credential_name
    )
  }
}

# Cross-reference test: Use credential data source to reference in another vector store
data "litellm_credential" "shared_cred" {
  credential_name = litellm_credential.test_credential.credential_name

  depends_on = [litellm_credential.test_credential]
}

resource "litellm_vector_store" "cross_reference_store" {
  vector_store_name   = "cross-reference-test-store-${random_id.test_suffix.hex}"
  custom_llm_provider = "openai"

  vector_store_description = "Vector store using credential from data source"

  # Reference credential through data source
  litellm_credential_name = data.litellm_credential.shared_cred.credential_name

  litellm_params = {
    model = "text-embedding-ada-002"
  }

  vector_store_metadata = {
    test_type = "cross_reference"
    source    = "data_source"
  }
}

# Test that we can read the cross-referenced vector store
data "litellm_vector_store" "cross_reference_data" {
  vector_store_id = litellm_vector_store.cross_reference_store.vector_store_id

  depends_on = [litellm_vector_store.cross_reference_store]
}

output "cross_reference_validation" {
  description = "Validation of cross-referencing between data sources and resources"
  value = {
    credential_name_from_data_source = data.litellm_credential.shared_cred.credential_name
    credential_name_in_vector_store  = data.litellm_vector_store.cross_reference_data.litellm_credential_name
    credentials_match = (
      data.litellm_credential.shared_cred.credential_name ==
      data.litellm_vector_store.cross_reference_data.litellm_credential_name
    )
  }
}

# Summary output for test results
output "test_summary" {
  description = "Summary of all data source tests"
  value = {
    credential_data_source_working   = data.litellm_credential.test_cred_data.credential_name != ""
    vector_store_data_source_working = data.litellm_vector_store.test_vector_data.vector_store_id != ""
    cross_reference_working = (
      data.litellm_credential.shared_cred.credential_name ==
      data.litellm_vector_store.cross_reference_data.litellm_credential_name
    )
    all_tests_passed = (
      data.litellm_credential.test_cred_data.credential_name != "" &&
      data.litellm_vector_store.test_vector_data.vector_store_id != "" &&
      data.litellm_credential.shared_cred.credential_name ==
      data.litellm_vector_store.cross_reference_data.litellm_credential_name
    )
  }
}
