package dhcp4server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/vitalvas/cloudrouter/lib/general"
)

type Server struct {
	cfg DHCPInterfaces
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

func (this *Server) Shutdown() {
}

func (this *Server) Apply() error {
	b, err := ioutil.ReadFile(filepath.Join(general.ConfigDir, "dhcp4_server.json"))
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if err := json.Unmarshal(b, &this.cfg); err != nil {
			return fmt.Errorf("cannot unmarshal config: %w", err)
		}
	}

	return nil
}
