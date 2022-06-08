package dhcp4server

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (this *Server) Shutdown() {

}

func (this *Server) Apply() error {
	return nil
}
