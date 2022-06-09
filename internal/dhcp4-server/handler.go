package dhcp4server

import (
	"log"
	"net"
	"sync"

	"github.com/krolaw/dhcp4"
)

type Handler struct {
	serverIP    net.IP
	leasesMutex sync.Mutex
}

func (this *Handler) ServeDHCP(p dhcp4.Packet, msgType dhcp4.MessageType, options dhcp4.Options) dhcp4.Packet {
	hwAddr := p.CHAddr().String()

	switch msgType {
	case dhcp4.Discover:

	case dhcp4.Request:

	case dhcp4.Release, dhcp4.Decline:
		if this.expireLease(hwAddr) {
			log.Printf("expired lease for %v", hwAddr)
		}

		return nil
	}

	return nil
}

func (this *Handler) expireLease(hwAddr string) bool {
	this.leasesMutex.Lock()
	defer this.leasesMutex.Unlock()

	return true
}
