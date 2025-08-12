package models

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func buildCredentialData(d *schema.ResourceData) map[string]interface{} {
	credentialData := make(map[string]interface{})

	// String fields
	utils.GetValueDefault[string](d, "credential_name", credentialData)
	utils.GetValueDefault[string](d, "model_id", credentialData)

	// Map fields
	utils.GetValueDefault[map[string]interface{}](d, "credential_info", credentialData)
	utils.GetValueDefault[map[string]interface{}](d, "credential_values", credentialData)

	return credentialData
}

func setCredentialResourceData(d *schema.ResourceData, credential *Credential) error {
	fields := map[string]interface{}{
		"credential_name": credential.CredentialName,
		"model_id":        credential.ModelID,
		"credential_info": credential.CredentialInfo,
		// Note: We don't set credential_values for security reasons
	}

	for field, value := range fields {
		if err := d.Set(field, value); err != nil {
			log.Printf("[WARN] Error setting %s: %s", field, err)
			return fmt.Errorf("error setting %s: %s", field, err)
		}
	}

	return nil
}

func parseCredentialAPIResponse(resp map[string]interface{}) (*Credential, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	credential := &Credential{}

	for k, v := range resp {
		if v == nil {
			continue
		}

		switch k {
		case "credential_name":
			if s, ok := v.(string); ok {
				credential.CredentialName = s
			}
		case "model_id":
			if s, ok := v.(string); ok {
				credential.ModelID = s
			}
		case "credential_info":
			if m, ok := v.(map[string]interface{}); ok {
				credential.CredentialInfo = m
			}
		case "credential_values":
			if m, ok := v.(map[string]interface{}); ok {
				credential.CredentialValues = m
			}
		}
	}

	return credential, nil
}

func buildCredentialForCreation(data map[string]interface{}) *Credential {
	credential := &Credential{}

	if v, ok := data["credential_name"].(string); ok {
		credential.CredentialName = v
	}
	if v, ok := data["model_id"].(string); ok {
		credential.ModelID = v
	}
	if v, ok := data["credential_info"].(map[string]interface{}); ok {
		credential.CredentialInfo = v
	}
	if v, ok := data["credential_values"].(map[string]interface{}); ok {
		credential.CredentialValues = v
	}

	return credential
}
