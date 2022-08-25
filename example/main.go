package main

import (
	"SimpleSkeleton/app"
	"SimpleSkeleton/example/boot"
	"flag"
)

var initConfigPath = flag.String(
	"etc_config",
	"./config.yaml",
	"config path")

func main() {
	boot.Boot(app.CreateConfiguration(initConfigPath))
}
