package main

import (
	"log"
	"os"
	"path/filepath"

	dhcp4server "github.com/vitalvas/cloudrouter/internal/dhcp4-server"
	"github.com/vitalvas/cloudrouter/internal/dns"
	"github.com/vitalvas/cloudrouter/internal/neighbor"
	"github.com/vitalvas/cloudrouter/lib/runner"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	var app runner.Runner

	switch getExecName() {
	case "cloudrouter-dhcp4-server", "dhcp4-server":
		app = dhcp4server.NewServer()

	case "cloudrouter-dns", "dns":
		app = dns.NewServer()

	case "cloudrouter-neighbor", "neighbor":
		app = neighbor.NewServer()

	default:
		log.Fatal("program name recognition error")
	}

	runner.Execute(app)
}

func getExecName() string {
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Base(path)
}
