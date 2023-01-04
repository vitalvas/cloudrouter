package neighbor

const (
	addrIPv6 = "ff02::f2"
	addrPort = 1024
)

func (s *Server) server() {
	defer s.lock.Unlock()

	go s.sender()

	go s.receiver()
}
