package key

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// resourceKeyResourceV0 returns the schema for version 0 of the key resource
func resourceKeyResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: resourceKeySchema(),
	}
}

// resourceKeyStateUpgradeV0 handles migration from schema version 0 to 1
func resourceKeyStateUpgradeV0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	// Check if token_id exists in state and is not empty
	if tokenID, exists := rawState["token_id"]; exists && tokenID != nil && tokenID.(string) != "" {
		// Use token_id as the new ID
		rawState["id"] = tokenID.(string)
		return rawState, nil
	}

	// If no token_id, try to find the key using key_alias
	keyAlias, aliasExists := rawState["key_alias"]
	if !aliasExists || keyAlias == nil || keyAlias.(string) == "" {
		return nil, fmt.Errorf("cannot migrate state: no token_id found and no key_alias available to lookup the key")
	}

	// Use the LiteLLM client to find the key by alias
	c := meta.(*litellm.Client)
	keyItem, err := findKeyByAlias(ctx, c, keyAlias.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to find key by alias during state migration: %w", err)
	}

	// Update the ID to use the token from the API response
	rawState["id"] = keyItem.Token
	rawState["token_id"] = keyItem.Token

	return rawState, nil
}
