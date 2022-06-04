package vrouter

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vitalvas/cloudrouter/app/vrouter/config"
)

type VRouter struct {
	config   *config.Config
	shutdown bool
}

func NewVRouter() *VRouter {
	this := &VRouter{}

	var err error

	this.config, err = LoadOrCreateConfig()
	if err != nil {
		log.Fatal(err)
	}

	return this
}

func (this *VRouter) Execute() {
	go this.background()

	log.Println("started")

	notifyCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-notifyCtx.Done()

	log.Println("Shutdown signal received")

	this.saveConfig()
}

func (this *VRouter) background() {
	for {
		time.Sleep(time.Second)

		this.ApplySysctl()
	}
}
