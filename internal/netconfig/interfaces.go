package netconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vishvananda/netlink"
	"github.com/vitalvas/cloudrouter/lib/general"
)

type Interfaces struct {
}

type Interface struct {
	Name         string `json:"name"`
	HardwareAddr string `json:"hardware_addr"`
	MTU          int    `json:"mtu"`
}

func NewInterfaces() *Interfaces {
	return &Interfaces{}
}

func (iface *Interfaces) Apply() error {
	_, err := os.ReadFile(filepath.Join(general.ConfigDir, "interfaces.json"))
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := iface.generateInterfaces(); err != nil {
			return err
		}
	}

	return nil
}

func (iface *Interfaces) generateInterfaces() error {
	links, err := netlink.LinkList()
	if err != nil {
		return err
	}

	for _, link := range links {
		attrs := link.Attrs()
		if attrs.EncapType == "ether" {
			iface := Interface{
				Name:         attrs.Name,
				HardwareAddr: attrs.HardwareAddr.String(),
			}

			fmt.Printf("%#v\n", iface)
		}
	}

	return nil
}
