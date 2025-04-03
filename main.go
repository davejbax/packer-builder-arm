package main

import (
	"github.com/davejbax/packer-plugin-arm/builder"
	"github.com/davejbax/packer-plugin-arm/version"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"log"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, builder.NewBuilder())
	pps.SetVersion(version.PluginVersion)
	if err := pps.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
