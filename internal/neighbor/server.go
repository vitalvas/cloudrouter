package neighbor

import (
	"encoding/binary"
	"sync"
	"time"
)

type Server struct {
	ID       []byte
	lock     sync.Mutex
	shutdown bool

	neighbors sync.Map
}

func NewServer() *Server {
	return &Server{
		ID: getMachineID(),
	}
}

func (s *Server) Apply() error {
	if s.lock.TryLock() {
		go s.server()
	}

	return nil
}

func (s *Server) Shutdown() {
	s.shutdown = true
}

func getMachineID() []byte {
	b := make([]byte, 8)
	id := time.Now().UnixNano()
	binary.LittleEndian.PutUint64(b, uint64(id))
	return b
}
