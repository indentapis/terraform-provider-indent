package main

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	provider "go.indent.com/terraform-provider-indent/provider"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		Debug:        os.Getenv("INDENT_PROVIDER_DEBUG") == "true",
		ProviderAddr: "terraform.io/indentapis/indent",
		ProviderFunc: provider.Provider,
	})
}
