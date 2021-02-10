package main

import (
	"github.com/GreyNoise-Intelligence/terraform-provider-clc/clc"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: clc.Provider})
}
