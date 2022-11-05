package neighbor

import (
	"log"
	"net"
	"time"

	"github.com/vishvananda/netlink"
	"golang.org/x/exp/slices"
)

const (
	addrIPv4 = "224.0.0.242"
	addrIPv6 = "ff02::f2"
	addrPort = 1024
)

var (
	ipv4Addr = &net.UDPAddr{
		IP:   net.ParseIP(addrIPv4),
		Port: addrPort,
	}
	ipv6Addr = &net.UDPAddr{
		IP:   net.ParseIP(addrIPv6),
		Port: addrPort,
	}

	allowedLinkType = []string{
		"device", "tuntap",
	}
)

func (s *Server) server() {
	defer s.lock.Unlock()

	for {
		list, err := netlink.LinkList()
		if err != nil {
			log.Println("error", err)
			continue
		}

		for _, row := range list {
			if slices.Contains(allowedLinkType, row.Type()) {
				sendNeighborPacket(row.Attrs())
			}
		}

		time.Sleep(time.Minute)
	}
}

func sendNeighborPacket(link *netlink.LinkAttrs) {
	iface, err := net.InterfaceByIndex(link.Index)
	if err != nil {
		return
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return
	}

	ipv4, ipv6 := false, false

	for _, row := range addrs {
		addr := row.(*net.IPNet)

		if !ipv4 {
			ipv4 = addr.IP.To4() != nil
		}

		if !ipv6 {
			ipv6 = len(addr.IP.To16()) == net.IPv6len
		}
	}

	if ipv4 {
		sendNeighborPacketIPv4(link)
	}

	if ipv6 {
		sendNeighborPacketIPv6(link)
	}
}

func sendNeighborPacketIPv4(link *netlink.LinkAttrs) {

}

func sendNeighborPacketIPv6(link *netlink.LinkAttrs) {

}
