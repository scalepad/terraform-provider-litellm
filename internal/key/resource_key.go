package key

import (
	"context"
	"fmt"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKeyCreate,
		ReadContext:   resourceKeyRead,
		UpdateContext: resourceKeyUpdate,
		DeleteContext: resourceKeyDelete,
		Schema:        resourceKeySchema(),
	}
}

func resourceKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	request := buildKeyGenerateRequest(d)

	createdKeyResponse, err := createKey(ctx, c, request)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating key: %s", err))
	}

	d.SetId(createdKeyResponse.Key)

	// Set the resource data with the created key information
	// This includes the sensitive key which is only available during creation
	if err := setKeyResourceDataFromGenerate(d, createdKeyResponse); err != nil {
		return diag.FromErr(err)
	}

	return resourceKeyRead(ctx, d, m)
}

func resourceKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	keyInfoResponse, err := getKey(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading key: %s", err))
	}

	if keyInfoResponse == nil {
		d.SetId("")
		return nil
	}

	// Update resource data with API response, but preserve state values for certain fields
	if err := setKeyResourceDataFromInfo(d, keyInfoResponse); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceKeyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	request := buildKeyUpdateRequest(d)

	_, err := updateKey(ctx, c, d.Id(), request)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating key: %s", err))
	}

	return resourceKeyRead(ctx, d, m)
}

func resourceKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	err := deleteKey(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting key: %s", err))
	}

	d.SetId("")
	return nil
}
