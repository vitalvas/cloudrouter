package neighbor

import (
	"bytes"
	"log"
	"net"
	"time"

	"github.com/vishvananda/netlink"
	"github.com/vitalvas/cloudrouter/internal/neighbor/message"
	"github.com/vitalvas/cloudrouter/lib/tools"
	"google.golang.org/protobuf/proto"
)

const (
	maxDatagramSize = 1280
)

type Neighbor struct {
	ID              []byte
	Interface       string
	Identity        string
	MACAddress      string
	RemoteInterface string
	Protocol        []string
	Address         []string
	LastContact     time.Time
}

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
		numBytes, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}

		s.receiverHander(link.Name, buffer[:numBytes])
	}

	return nil
}

func (s *Server) receiverHander(ifname string, data []byte) {
	var msg message.Neighbor
	if err := proto.Unmarshal(data, &msg); err != nil {
		return
	}

	if bytes.Equal(s.ID, msg.GetChassisId()) {
		return
	}

	neighbor := Neighbor{
		ID:              msg.GetChassisId(),
		Interface:       ifname,
		Identity:        msg.GetChassisName(),
		RemoteInterface: msg.GetChassisInterface(),
		MACAddress:      net.HardwareAddr(msg.GetChassisMac()).String(),
		Protocol:        []string{"crdp"},
		LastContact:     time.Now(),
	}

	for _, row := range msg.GetMgmtAddress() {
		addr := net.IP(row)
		if addr == nil {
			continue
		}

		neighbor.Address = append(neighbor.Address, addr.String())
	}

	s.neighbors.Store(msg.GetChassisId(), neighbor)

	log.Printf("%#v", neighbor)
}
