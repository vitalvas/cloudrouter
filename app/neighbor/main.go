package main

import (
	"log"

	"github.com/vitalvas/cloudrouter/internal/neighbor"
	"github.com/vitalvas/cloudrouter/lib/runner"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	runner.Execute(neighbor.NewServer())
}
