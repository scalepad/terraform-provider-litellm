package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/key"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"github.com/scalepad/terraform-provider-litellm/internal/models"
	"github.com/scalepad/terraform-provider-litellm/internal/models/creds"
	"github.com/scalepad/terraform-provider-litellm/internal/team"
	"github.com/scalepad/terraform-provider-litellm/internal/team/member"
	"github.com/scalepad/terraform-provider-litellm/internal/tools/mcp"
	"github.com/scalepad/terraform-provider-litellm/internal/tools/vector"
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
		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure configures the provider with the given schema data.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := litellm.ProviderConfig{
		APIBase:            d.Get("api_base").(string),
		APIKey:             d.Get("api_key").(string),
		InsecureSkipVerify: d.Get("insecure_skip_verify").(bool),
	}

	return litellm.NewClient(config.APIBase, config.APIKey, config.InsecureSkipVerify), nil
}
