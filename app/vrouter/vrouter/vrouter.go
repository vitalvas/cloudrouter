package vrouter

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vitalvas/cloudrouter/app/vrouter/config"
	"github.com/vitalvas/cloudrouter/app/vrouter/wireguard"
)

type VRouter struct {
	config    *config.Config
	wireguard *wireguard.Wireguard
	shutdown  bool
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
	var err error

	this.wireguard, err = wireguard.New()
	if err != nil {
		log.Fatal(err)
	}

	this.wireguard.SetDevices(this.config.VPN.WireGuard)

	go this.background()

	log.Println("started")

	notifyCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-notifyCtx.Done()

	log.Println("Shutdown signal received")

	this.saveConfig()
	this.wireguard.Shutdown()
}

func (this *VRouter) background() {
	for {
		this.wireguard.SyncConfig()

		time.Sleep(10 * time.Second)
	}
}
