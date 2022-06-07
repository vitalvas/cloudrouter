package netconfig

import "github.com/vitalvas/cloudrouter/lib/logger"

var log = logger.NewConsole()

type NetConfig struct {
	wireguard *Wireguard
	firewall  *Firewall
}

func NewNetConfig() *NetConfig {
	this := &NetConfig{}

	var err error
	this.wireguard, err = NewWireguard()
	if err != nil {
		log.Fatal(err)
	}

	this.firewall, err = NewFirewall()
	if err != nil {
		log.Fatal(err)
	}

	return this
}

func (this *NetConfig) Apply() {
	if err := applySysctl(); err != nil {
		log.Println("sysctl:", err)
	}

	if err := this.firewall.apply(); err != nil {
		log.Println("firewall:", err)
	}

	if err := this.wireguard.apply(); err != nil {
		log.Println("wireguard:", err)
	}
}
