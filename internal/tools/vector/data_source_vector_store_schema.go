package vector

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVectorStoreSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vector_store_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Unique identifier for the vector store to retrieve",
		},
		"vector_store_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the vector store",
		},
		"custom_llm_provider": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Custom LLM provider for the vector store",
		},
		"vector_store_description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Description of the vector store",
		},
		"vector_store_metadata": {
			Type:        schema.TypeMap,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Metadata associated with the vector store",
		},
		"litellm_credential_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the LiteLLM credential used",
		},
		"litellm_params": {
			Type:        schema.TypeMap,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Additional LiteLLM parameters",
		},
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the vector store was created",
		},
		"updated_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the vector store was last updated",
		},
	}
}
