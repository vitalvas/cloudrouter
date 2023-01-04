package neighbor

import (
	"log"
	"net"

	"github.com/vishvananda/netlink"
	"github.com/vitalvas/cloudrouter/internal/neighbor/message"
	"github.com/vitalvas/cloudrouter/lib/tools"
	"google.golang.org/protobuf/proto"
)

const (
	maxDatagramSize = 1024
)

func (s *Server) receiver() {
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

			go s.receiverInterface(attrs)
		}
	}

}

func (s *Server) receiverInterface(link *netlink.LinkAttrs) error {
	raddr := &net.UDPAddr{
		IP:   net.ParseIP(addrIPv6),
		Port: addrPort,
		Zone: link.Name,
	}

	conn, err := net.ListenMulticastUDP("udp6", nil, raddr)
	if err != nil {
		return err
	}

	defer conn.Close()

	conn.SetReadBuffer(maxDatagramSize)

	for !s.shutdown {
		buffer := make([]byte, maxDatagramSize)
		numBytes, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}

		s.receiverHander(src, buffer[:numBytes])
	}

	return nil
}

func (s *Server) receiverHander(addr *net.UDPAddr, data []byte) {
	var msg message.Neighbor
	if err := proto.Unmarshal(data, &msg); err != nil {
		return
	}

	log.Println(msg.String())
}
