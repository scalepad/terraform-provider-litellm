package vector

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

func DataSourceLiteLLMVectorStore() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLiteLLMVectorStoreRead,
		Schema:      dataSourceVectorStoreSchema(),
	}
}

func dataSourceLiteLLMVectorStoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)
	vectorStoreID := d.Get("vector_store_id").(string)

	vectorStore, err := getVectorStore(ctx, c, vectorStoreID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read vector store: %w", err))
	}

	if vectorStore == nil {
		return diag.FromErr(fmt.Errorf("vector store '%s' not found", vectorStoreID))
	}

	// Set the data source ID to the vector store ID
	d.SetId(vectorStore.VectorStoreID)

	if err := setVectorStoreResourceData(d, vectorStore); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
