package runner

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vitalvas/cloudrouter/lib/logger"
)

type Runner interface {
	Shutdown()
	Apply() error
}

func Execute(srv Runner) {
	var log = logger.NewConsole()

	defer srv.Shutdown()

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
