package vector

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

func ResourceLiteLLMVectorStore() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLiteLLMVectorStoreCreate,
		ReadContext:   resourceLiteLLMVectorStoreRead,
		UpdateContext: resourceLiteLLMVectorStoreUpdate,
		DeleteContext: resourceLiteLLMVectorStoreDelete,
		Schema:        resourceVectorStoreSchema(),
	}
}

func resourceLiteLLMVectorStoreCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	vectorStoreData := buildVectorStoreData(d)
	vectorStore := buildVectorStoreForCreation(vectorStoreData)

	createdVectorStore, err := createVectorStore(ctx, c, vectorStore)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating vector store: %s", err))
	}

	d.SetId(createdVectorStore.VectorStoreID)
	return resourceLiteLLMVectorStoreRead(ctx, d, m)
}

func resourceLiteLLMVectorStoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	vectorStore, err := getVectorStore(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading vector store: %s", err))
	}

	if vectorStore == nil {
		d.SetId("")
		return nil
	}

	if err := setVectorStoreResourceData(d, vectorStore); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceLiteLLMVectorStoreUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	vectorStoreData := buildVectorStoreData(d)
	vectorStore := buildVectorStoreForCreation(vectorStoreData)
	vectorStore.VectorStoreID = d.Id() // Set the vector store ID for update

	_, err := updateVectorStore(ctx, c, vectorStore)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating vector store: %s", err))
	}

	return resourceLiteLLMVectorStoreRead(ctx, d, m)
}

func resourceLiteLLMVectorStoreDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	err := deleteVectorStore(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting vector store: %s", err))
	}

	d.SetId("")
	return nil
}
