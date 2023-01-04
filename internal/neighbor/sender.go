package neighbor

import (
	"log"
	"net"
	"time"

	"github.com/vishvananda/netlink"
	"github.com/vitalvas/cloudrouter/lib/tools"
)

var (
	allowedLinkType = []string{
		"device", "tuntap",
	}
)

func (s *Server) sender() {
	for !s.shutdown {
		list, err := netlink.LinkList()
		if err != nil {
			log.Println("error", err)
			return
		}

		for _, row := range list {
			if tools.SlicesContains(allowedLinkType, row.Type()) {
				attrs := row.Attrs()
				if attrs.EncapType != "ether" {
					continue
				}

				s.sendInterfaceNeighborPacket(attrs)
			}
		}

		time.Sleep(time.Minute)
	}
}

func (s *Server) sendInterfaceNeighborPacket(link *netlink.LinkAttrs) {
	iface, err := net.InterfaceByIndex(link.Index)
	if err != nil {
		return
	}

	addrs, err := iface.Addrs()
	if err != nil || addrs == nil {
		return
	}

	msg, err := s.makeNeighborPacket(link, addrs)
	if err != nil {
		return
	}

	if err := sendNeighborPacket(link, msg); err != nil {
		log.Println("err send:", err)
	}
}

func sendNeighborPacket(link *netlink.LinkAttrs, msg []byte) error {
	raddr := &net.UDPAddr{
		IP:   net.ParseIP(addrIPv6),
		Port: addrPort,
		Zone: link.Name,
	}

	conn, err := net.DialUDP("udp6", nil, raddr)
	if err != nil {
		return err
	}

	defer conn.Close()

	if _, err := conn.Write(msg); err != nil {
		return err
	}

	return nil
}
