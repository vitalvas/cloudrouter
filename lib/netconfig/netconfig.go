package netconfig

import "github.com/vitalvas/cloudrouter/lib/logger"

var log = logger.NewConsole()

type NetConfig struct {
}

func NewNetConfig() *NetConfig {
	return &NetConfig{}
}

func (this *NetConfig) Apply() {
	if err := applySysctl(); err != nil {
		log.Println("sysctl:", err)
	}

	if err := applyWireGuard(); err != nil {
		log.Println("wireguard:", err)
	}
}
