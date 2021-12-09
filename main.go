package main

import (
	"github.com/g42cloud-terraform/terraform-provider-g42vbcloud/g42vbcloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: g42vbcloud.Provider})
}
