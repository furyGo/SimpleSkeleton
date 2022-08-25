package main

import (
	"flag"
	"github.com/furyGo/SimpleSkeleton/app"
	"github.com/furyGo/SimpleSkeleton/example/boot"
)

var initConfigPath = flag.String(
	"etc_config",
	"./config.yaml",
	"config path")

func main() {
	boot.Boot(app.CreateConfiguration(initConfigPath))
}
