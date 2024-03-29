package dhcp4server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/vitalvas/cloudrouter/lib/general"
)

type Server struct {
	cfg DHCPInterfaces
	log *log.Logger
}

type DHCPInterfaces struct {
	Interfaces []DHCPInterface `json:"interfaces"`
}

type DHCPInterface struct {
	Name      string        `json:"name"`
	LeaseTime time.Duration `json:"lease_time"`
}

func NewServer() *Server {
	return &Server{}
}

func (h *Server) Shutdown() {
}

func (h *Server) SetLogger(l *log.Logger) {
	h.log = l
}

func (h *Server) Apply() error {
	b, err := os.ReadFile(filepath.Join(general.ConfigDir, "dhcp4_server.json"))
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if err := json.Unmarshal(b, &h.cfg); err != nil {
			return fmt.Errorf("cannot unmarshal config: %w", err)
		}
	}

	return nil
}
