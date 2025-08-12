package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/key"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"github.com/scalepad/terraform-provider-litellm/internal/models"
	"github.com/scalepad/terraform-provider-litellm/internal/models/creds"
	"github.com/scalepad/terraform-provider-litellm/internal/team"
	"github.com/scalepad/terraform-provider-litellm/internal/team/member"
	"github.com/scalepad/terraform-provider-litellm/internal/tools/mcp"
	"github.com/scalepad/terraform-provider-litellm/internal/tools/vector"
	"github.com/scalepad/terraform-provider-litellm/internal/users"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"litellm_model":           models.ResourceModel(),
			"litellm_team":            team.ResourceTeam(),
			"litellm_team_member":     member.ResourceTeamMember(),
			"litellm_team_member_add": member.ResourceTeamMemberAdd(),
			"litellm_key":             key.ResourceKey(),
			"litellm_mcp_server":      mcp.ResourceLiteLLMMCPServer(),
			"litellm_credential":      creds.ResourceCredential(),
			"litellm_vector_store":    vector.ResourceLiteLLMVectorStore(),
			"litellm_user":            users.ResourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"litellm_credential":   creds.DataSourceLiteLLMCredential(),
			"litellm_vector_store": vector.DataSourceLiteLLMVectorStore(),
		},
		Schema: map[string]*schema.Schema{
			"api_base": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   false,
				DefaultFunc: schema.EnvDefaultFunc("LITELLM_API_BASE", nil),
				Description: "The base URL of the LiteLLM API",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LITELLM_API_KEY", nil),
				Description: "The API key for authenticating with LiteLLM",
			},
			"insecure_skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Skip TLS certificate verification when connecting to the LiteLLM API",
			},
		},
		ConfigureContextFunc: providerConfigureContext,
	}
}

// providerConfigureContext configures the provider with the given schema data and tests the connection.
func providerConfigureContext(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := litellm.ProviderConfig{
		APIBase:            d.Get("api_base").(string),
		APIKey:             d.Get("api_key").(string),
		InsecureSkipVerify: d.Get("insecure_skip_verify").(bool),
	}

	client := litellm.NewClient(config.APIBase, config.APIKey, config.InsecureSkipVerify)

	// Test the connection by calling GET /models
	if err := testConnection(ctx, client); err != nil {
		return nil, diag.FromErr(fmt.Errorf("failed to connect to LiteLLM API: %v", err))
	}

	return client, nil
}

// testConnection tests the connection to the LiteLLM API by calling GET /models
func testConnection(ctx context.Context, client *litellm.Client) error {
	_, err := client.SendRequest(ctx, http.MethodGet, "/models", nil)
	if err != nil {
		return fmt.Errorf("connection test failed: %w", err)
	}
	return nil
}
