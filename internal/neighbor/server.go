package neighbor

import "sync"

type Server struct {
	lock sync.Mutex
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Apply() error {
	if s.lock.TryLock() {
		go s.server()
	}

	return nil
}

func (s *Server) Shutdown() {

}
