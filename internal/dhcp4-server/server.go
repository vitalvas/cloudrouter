package dhcp4server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vitalvas/cloudrouter/lib/general"
	"github.com/vitalvas/cloudrouter/lib/multilisten"
)

type Server struct {
	listeners *multilisten.Pool
	cfg       DHCPInterfaces
}

type DHCPInterfaces struct {
	Interfaces []DHCPInterface `json:"interfaces"`
}

type DHCPInterface struct {
	Name string `json:"name"`
}

func NewServer() *Server {
	return &Server{
		listeners: multilisten.NewPool(),
	}
}

func (this *Server) Shutdown() {
	this.listeners.Close()
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
