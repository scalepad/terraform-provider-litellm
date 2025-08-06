package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/scalepad/terraform-provider-litellm/litellm"
)

// main is the entry point for the plugin. It serves the provider
// using the Terraform plugin SDK.
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: litellm.Provider,
	})
}
