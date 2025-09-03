package serviceaccount

import (
	"context"
	"fmt"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceServiceAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceAccountCreate,
		ReadContext:   resourceServiceAccountRead,
		UpdateContext: resourceServiceAccountUpdate,
		DeleteContext: resourceServiceAccountDelete,
		Schema:        resourceServiceAccountSchema(),
	}
}

func resourceServiceAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	// Validate required fields
	if teamID := d.Get("team_id").(string); teamID == "" {
		return diag.Errorf("team_id is required for service account creation")
	}

	request := buildServiceAccountGenerateRequest(d)

	createdServiceAccountResponse, err := CreateServiceAccount(ctx, c, request)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating service account: %s", err))
	}

	d.SetId(createdServiceAccountResponse.TokenID)

	// Set the resource data with the created service account information
	if err := setServiceAccountResourceData(d, createdServiceAccountResponse); err != nil {
		return diag.FromErr(err)
	}

	return resourceServiceAccountRead(ctx, d, m)
}

func resourceServiceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	serviceAccountInfoResponse, err := GetServiceAccount(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading service account: %s", err))
	}

	if serviceAccountInfoResponse == nil {
		d.SetId("")
		return nil
	}

	// Update resource data with API response
	if err := setServiceAccountResourceDataFromInfo(d, serviceAccountInfoResponse); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceServiceAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	request := buildServiceAccountUpdateRequest(d)

	_, err := UpdateServiceAccount(ctx, c, d.Id(), request)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating service account: %s", err))
	}

	return resourceServiceAccountRead(ctx, d, m)
}

func resourceServiceAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	err := DeleteServiceAccount(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting service account: %s", err))
	}

	d.SetId("")
	return nil
}
