package netconfig

import (
	"github.com/vishvananda/netlink"
	"golang.zx2c4.com/wireguard/wgctrl"
)

type Wireguard struct {
	client  *wgctrl.Client
	devices []WireGuardDevice
}

type WireGuardDevice struct {
	DeviceName string
	PrivateKey string
	ListenPort int
}

func NewWireguard() (*Wireguard, error) {
	this := &Wireguard{}

	var err error

	this.client, err = wgctrl.New()
	if err != nil {
		return nil, err
	}

	return this, nil
}

func (this *Wireguard) SetDevices(devs []WireGuardDevice) {
	this.devices = devs
}

func (this *Wireguard) Shutdown() {
	this.client.Close()
}

func (this *Wireguard) applyWireGuard() error {
	devs, err := this.client.Devices()
	if err != nil {
		log.Fatal(err)
	}

	var unknownDevices []string

	for _, dev := range devs {
		exists := false
		for _, row := range this.devices {
			if dev.Name == row.DeviceName {
				exists = true
			}
		}

		if !exists {
			unknownDevices = append(unknownDevices, dev.Name)
		}
	}

	for _, row := range unknownDevices {
		log.Println("deleting", row)

		link, err := netlink.LinkByName(row)
		if err != nil {
			return err
		}

		if err := netlink.LinkDel(link); err != nil {
			return err
		}
	}

	return nil
}
