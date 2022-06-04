package wireguard

import (
	"github.com/vitalvas/cloudrouter/app/vrouter/config"
	"golang.zx2c4.com/wireguard/wgctrl"
)

type Wireguard struct {
	client  *wgctrl.Client
	devices []config.WireGuard
}

func New() (*Wireguard, error) {
	this := &Wireguard{}

	var err error

	this.client, err = wgctrl.New()
	if err != nil {
		return nil, err
	}

	return this, nil
}

func (this *Wireguard) SetDevices(devs []config.WireGuard) {
	this.devices = devs
}

func (this *Wireguard) Shutdown() {
	this.client.Close()
}
