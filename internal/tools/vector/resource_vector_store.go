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

	request := buildVectorStoreGenerateRequest(d)

	createdVectorStoreResponse, err := createVectorStore(ctx, c, request)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating vector store: %s", err))
	}

	d.SetId(createdVectorStoreResponse.VectorStore.VectorStoreID)

	// Set the resource data with the created vector store information
	if err := setVectorStoreResourceDataFromGenerate(d, createdVectorStoreResponse); err != nil {
		return diag.FromErr(err)
	}

	return resourceLiteLLMVectorStoreRead(ctx, d, m)
}

func resourceLiteLLMVectorStoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	vectorStoreInfoResponse, err := getVectorStore(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading vector store: %s", err))
	}

	if vectorStoreInfoResponse == nil {
		d.SetId("")
		return nil
	}

	// Update resource data with API response, but preserve state values for certain fields
	if err := setVectorStoreResourceDataFromInfo(d, vectorStoreInfoResponse); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceLiteLLMVectorStoreUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	request := buildVectorStoreUpdateRequest(d)

	_, err := updateVectorStore(ctx, c, d.Id(), request)
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
