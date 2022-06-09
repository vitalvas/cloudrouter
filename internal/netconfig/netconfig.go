package netconfig

import "github.com/vitalvas/cloudrouter/lib/logger"

var log = logger.NewConsole()

type NetConfig struct {
	interfaces *Interfaces
	wireguard  *Wireguard
	firewall   *Firewall
}

func NewNetConfig() *NetConfig {
	this := &NetConfig{
		interfaces: NewInterfaces(),
	}

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

func (this *NetConfig) Shutdown() {
	this.wireguard.Close()
}

func (this *NetConfig) Apply() error {
	if err := applySysctl(); err != nil {
		log.Println("sysctl:", err)
	}

	if err := this.interfaces.Apply(); err != nil {
		log.Println("interfaces:", err)
	}

	if err := this.firewall.apply(); err != nil {
		log.Println("firewall:", err)
	}

	if err := this.wireguard.apply(); err != nil {
		log.Println("wireguard:", err)
	}

	if err := applySysctl(); err != nil {
		log.Println("sysctl:", err)
	}

	return nil
}
