package main

import (
	dhcp4server "github.com/vitalvas/cloudrouter/internal/dhcp4-server"
	"github.com/vitalvas/cloudrouter/lib/runner"
)

func main() {
	runner.Execute(dhcp4server.NewServer())
}
