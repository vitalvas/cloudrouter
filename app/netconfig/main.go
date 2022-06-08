package main

import (
	"github.com/vitalvas/cloudrouter/internal/netconfig"
	"github.com/vitalvas/cloudrouter/lib/runner"
)

func main() {
	runner.Execute(netconfig.NewNetConfig())
}
