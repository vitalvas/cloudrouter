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
	config   config.Config
	shutdown bool
}

func NewVRouter() *VRouter {
	return &VRouter{}
}

func (this *VRouter) Execute() {
	this.config = config.NewConfig()

	// data, err := config.Marshal(this.config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// config.PackFile(this.config.ID, data)

	go this.background()

	notifyCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-notifyCtx.Done()

	log.Println("Shutdown signal received")
}

func (this *VRouter) background() {
	for {
		time.Sleep(time.Second)

		this.ApplySysctl()
	}
}
