package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/vitalvas/cloudrouter/lib/logger"
	"github.com/vitalvas/cloudrouter/lib/netconfig"
)

var log = logger.NewConsole()

func main() {
	srv := netconfig.NewNetConfig()

	srv.Apply()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR1, os.Interrupt)

	for sig := range ch {
		if sig == syscall.SIGUSR1 {
			srv.Apply()
		} else if sig == os.Interrupt {
			break
		}
	}

	log.Println("shutdown")

}
