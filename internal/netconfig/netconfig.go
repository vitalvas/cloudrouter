package netconfig

import "github.com/vitalvas/cloudrouter/lib/logger"

var log = logger.NewConsole()

type NetConfig struct {
	interfaces *Interfaces
	wireguard  *Wireguard
	firewall   *Firewall
}

func NewNetConfig() *NetConfig {
	nc := &NetConfig{
		interfaces: NewInterfaces(),
	}

	var err error
	nc.wireguard, err = NewWireguard()
	if err != nil {
		log.Fatal(err)
	}

	nc.firewall, err = NewFirewall()
	if err != nil {
		log.Fatal(err)
	}

	return nc
}

func (nc *NetConfig) Shutdown() {
	nc.wireguard.Close()
}

func (nc *NetConfig) Apply() error {
	if err := applySysctl(); err != nil {
		log.Println("sysctl:", err)
	}

	if err := nc.interfaces.Apply(); err != nil {
		log.Println("interfaces:", err)
	}

	if err := nc.firewall.apply(); err != nil {
		log.Println("firewall:", err)
	}

	if err := nc.wireguard.apply(); err != nil {
		log.Println("wireguard:", err)
	}

	if err := applySysctl(); err != nil {
		log.Println("sysctl:", err)
	}

	return nil
}
