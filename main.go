package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/circa10a/terraform-provider-mcbroken/mcbroken"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mcbroken.Provider,
	})
}
