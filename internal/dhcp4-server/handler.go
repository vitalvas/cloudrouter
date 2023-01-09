package dhcp4server

import (
	"log"
	"sync"

	"github.com/krolaw/dhcp4"
)

type Handler struct {
	log *log.Logger
	// serverIP    net.IP
	leasesMutex sync.Mutex
}

func (h *Handler) ServeDHCP(p dhcp4.Packet, msgType dhcp4.MessageType, options dhcp4.Options) dhcp4.Packet {
	hwAddr := p.CHAddr().String()

	switch msgType {
	case dhcp4.Discover:

	case dhcp4.Request:

	case dhcp4.Release, dhcp4.Decline:
		if h.expireLease(hwAddr) {
			h.log.Printf("expired lease for %v", hwAddr)
		}

		return nil
	}

	return nil
}

func (h *Handler) expireLease(hwAddr string) bool {
	h.leasesMutex.Lock()
	defer h.leasesMutex.Unlock()

	return true
}
