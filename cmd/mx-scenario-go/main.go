package main

import (
	scencli "github.com/multiversx/mx-chain-scenario-cli-go/cli"
)

const version = "v2.1.0-alpha2"

func main() {
	scencli.ScenariosCLI(version)
}
