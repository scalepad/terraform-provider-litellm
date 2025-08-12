package models

import (
	"context"
	"fmt"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCredential() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCredentialCreate,
		ReadContext:   resourceCredentialRead,
		UpdateContext: resourceCredentialUpdate,
		DeleteContext: resourceCredentialDelete,
		Schema:        resourceCredentialSchema(),
	}
}

func resourceCredentialSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"credential_name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Name of the credential",
		},
		"model_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Model ID associated with this credential",
		},
		"credential_info": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Additional information about the credential",
		},
		"credential_values": {
			Type:        schema.TypeMap,
			Required:    true,
			Sensitive:   true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Sensitive credential values (API keys, tokens, etc.)",
		},
	}
}

func resourceCredentialCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	credentialData := buildCredentialData(d)
	credential := buildCredentialForCreation(credentialData)

	createdCredential, err := createCredential(ctx, c, credential)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating credential: %s", err))
	}

	d.SetId(createdCredential.CredentialName)
	return resourceCredentialRead(ctx, d, m)
}

func resourceCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	credential, err := getCredential(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading credential: %s", err))
	}

	if credential == nil {
		d.SetId("")
		return nil
	}

	if err := setCredentialResourceData(d, credential); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceCredentialUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	credentialData := buildCredentialData(d)
	credential := buildCredentialForCreation(credentialData)
	credential.CredentialName = d.Id() // Set the credential name for update

	_, err := updateCredential(ctx, c, credential)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating credential: %s", err))
	}

	return resourceCredentialRead(ctx, d, m)
}

func resourceCredentialDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	err := deleteCredential(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting credential: %s", err))
	}

	d.SetId("")
	return nil
}
