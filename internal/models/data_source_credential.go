package models

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

func DataSourceLiteLLMCredential() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLiteLLMCredentialRead,

		Schema: map[string]*schema.Schema{
			"credential_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the credential to retrieve",
			},
			"model_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Model ID associated with this credential",
			},
			"credential_info": {
				Type:        schema.TypeMap,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Additional information about the credential",
			},
			// Note: credential_values are not exposed in data sources for security reasons
		},
	}
}

func dataSourceLiteLLMCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)
	credentialName := d.Get("credential_name").(string)
	modelID := d.Get("model_id").(string)

	// Use the same endpoint as the resource read operation
	endpoint := fmt.Sprintf("/credentials/by_name/%s", credentialName)
	if modelID != "" {
		endpoint += fmt.Sprintf("?model_id=%s", modelID)
	}

	resp, err := c.SendRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read credential: %w", err))
	}

	if resp == nil {
		return diag.FromErr(fmt.Errorf("credential '%s' not found", credentialName))
	}

	credentialResp, err := parseCredentialResponse(resp)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse credential response: %w", err))
	}

	// Set the data source ID to the credential name
	d.SetId(credentialResp.CredentialName)
	d.Set("credential_name", credentialResp.CredentialName)
	d.Set("credential_info", credentialResp.CredentialInfo)
	// Note: We don't expose credential_values in data sources for security reasons

	return nil
}

func parseCredentialResponse(resp map[string]interface{}) (*litellm.CredentialResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	credentialResp := &litellm.CredentialResponse{}

	if credentialName, ok := resp["credential_name"].(string); ok {
		credentialResp.CredentialName = credentialName
	}

	if credentialInfo, ok := resp["credential_info"].(map[string]interface{}); ok {
		credentialResp.CredentialInfo = credentialInfo
	}

	// Note: We don't parse credential_values for security reasons in data sources

	return credentialResp, nil
}
