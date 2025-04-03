package main

import (
	"github.com/davejbax/packer-plugin-arm/builder"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"log"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, builder.NewBuilder())
	if err := pps.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}
