package neighbor

import (
	"net"
	"os"

	"github.com/vishvananda/netlink"
	"github.com/vitalvas/cloudrouter/internal/neighbor/message"
	"google.golang.org/protobuf/proto"
)

func (s *Server) makeNeighborPacket(link *netlink.LinkAttrs, addrs []net.Addr) ([]byte, error) {
	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	msg := &message.Neighbor{
		ChassisId:        s.ID,
		ChassisName:      host,
		ChassisMac:       link.HardwareAddr,
		ChassisInterface: link.Name,
	}

	for _, row := range addrs {
		addr := row.(*net.IPNet)
		if addr == nil {
			continue
		}

		msg.MgmtAddress = append(msg.MgmtAddress, addr.IP)
	}

	return proto.Marshal(msg)
}
