package dns

import (
	"sync"
	"time"

	"github.com/miekg/dns"
	"github.com/vitalvas/cloudrouter/lib/logger"
	"golang.org/x/time/rate"
)

var log = logger.NewConsole()

type Server struct {
	Mux *dns.ServeMux

	client    *dns.Client
	ratelimit *rate.Limiter

	upstreamLock sync.RWMutex
	upstream     []string
}

func NewServer() *Server {
	server := &Server{
		Mux: dns.NewServeMux(),
		client: &dns.Client{
			Timeout: 2 * time.Second,
		},
		ratelimit: rate.NewLimiter(rate.Every(time.Second), 10),
		upstream: []string{
			"8.8.8.8:53", "8.8.4.4:53",
			"1.1.1.1:53", "1.0.0.1:53",
		},
	}

	server.Mux.HandleFunc(".", server.handleRequest)

	return server
}

func (s *Server) upstreams() []string {
	s.upstreamLock.RLock()
	defer s.upstreamLock.RUnlock()

	result := make([]string, len(s.upstream))
	copy(result, s.upstream)

	return result
}
