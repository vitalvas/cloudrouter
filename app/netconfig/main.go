package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vitalvas/cloudrouter/lib/logger"
	"github.com/vitalvas/cloudrouter/lib/netconfig"
)

var log = logger.NewConsole()

func main() {
	srv := netconfig.NewNetConfig()

	srv.Apply()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR1, syscall.SIGINT)

	go func() {
		for {
			time.Sleep(time.Hour)
			ch <- syscall.SIGUSR1
		}
	}()

	for sig := range ch {
		if sig == syscall.SIGUSR1 {
			srv.Apply()
		} else if sig == syscall.SIGINT {
			break
		}
	}

	log.Println("shutdown")
}
