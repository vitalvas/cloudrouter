package wireguard

import (
	"log"

	"github.com/vishvananda/netlink"
)

func (this *Wireguard) SyncConfig() {
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
			log.Fatal(err)
		}

		if err := netlink.LinkDel(link); err != nil {
			log.Fatal(err)
		}
	}

}
