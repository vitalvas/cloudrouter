package main

import (
	"log"

	"github.com/vitalvas/cloudrouter/internal/loader"
	"github.com/vitalvas/cloudrouter/lib/runner"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	runner.Execute(loader.New())
}
