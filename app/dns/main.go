package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vitalvas/cloudrouter/lib/dns"
	"github.com/vitalvas/cloudrouter/lib/logger"
)

var log = logger.NewConsole()

func main() {
	srv := dns.NewServer()

	defer srv.Shutdown()

	if err := srv.Apply(); err != nil {
		log.Fatal("error start server: %w", err)
	}

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
			if err := srv.Apply(); err != nil {
				log.Printf("error update listener: %v", err)
			}
		} else if sig == syscall.SIGINT {
			break
		}
	}

	log.Println("shutdown")
}
