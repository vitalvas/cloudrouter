package netconfig

import "log"

type NetConfig struct {
	log *log.Logger

	interfaces *Interfaces
	wireguard  *Wireguard
	firewall   *Firewall
}

func NewNetConfig() *NetConfig {
	nc := &NetConfig{
		interfaces: NewInterfaces(),
	}

	var err error
	nc.wireguard, err = NewWireguard(nc.log)
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

func (nc *NetConfig) SetLogger(l *log.Logger) {
	nc.log = l
}

func (nc *NetConfig) Apply() error {
	if err := nc.applySysctl(); err != nil {
		nc.log.Println("sysctl:", err)
	}

	if err := nc.interfaces.Apply(); err != nil {
		nc.log.Println("interfaces:", err)
	}

	if err := nc.firewall.apply(); err != nil {
		nc.log.Println("firewall:", err)
	}

	if err := nc.wireguard.apply(); err != nil {
		nc.log.Println("wireguard:", err)
	}

	if err := nc.applySysctl(); err != nil {
		nc.log.Println("sysctl:", err)
	}

	return nil
}
