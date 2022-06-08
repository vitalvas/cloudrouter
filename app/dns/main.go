package main

import (
	"github.com/vitalvas/cloudrouter/internal/dns"
	"github.com/vitalvas/cloudrouter/lib/runner"
)

func main() {
	runner.Execute(dns.NewServer())
}
