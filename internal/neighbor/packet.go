package neighbor

import (
	"net"
	"os"

	"github.com/vishvananda/netlink"
	"github.com/vitalvas/cloudrouter/internal/neighbor/message"
	"google.golang.org/protobuf/proto"
)

func makeNeighborPacket(link *netlink.LinkAttrs, addrs []net.Addr) ([]byte, error) {
	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	msg := &message.Neighbor{
		Chassis: &message.Chassis{
			Name: host,
			Mac:  link.HardwareAddr,
		},
	}

	for _, row := range addrs {
		addr := row.(*net.IPNet)
		if addr == nil {
			continue
		}

		mgmt := &message.MgmtAddress{}

		if addr.IP.To4() != nil {
			mgmt.Version = 4
			mgmt.Address = addr.IP.To4()
		} else if len(addr.IP.To16()) == net.IPv6len {
			mgmt.Version = 6
			mgmt.Address = addr.IP.To16()
		}

		if mgmt.Version > 0 {
			msg.MgmtAddress = append(msg.MgmtAddress, mgmt)
		}
	}

	return proto.Marshal(msg)
}
