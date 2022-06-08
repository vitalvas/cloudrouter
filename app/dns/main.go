package main

import (
	"github.com/vitalvas/cloudrouter/lib/dns"
	"github.com/vitalvas/cloudrouter/lib/runner"
)

func main() {
	runner.Execute(dns.NewServer())
}
