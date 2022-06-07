package dns

import (
	"net"
	"sync"
	"time"

	"github.com/miekg/dns"
	miekgdns "github.com/miekg/dns"
	"github.com/vitalvas/cloudrouter/lib/general"
	"github.com/vitalvas/cloudrouter/lib/logger"
	"github.com/vitalvas/cloudrouter/lib/multilisten"
	"golang.org/x/time/rate"
)

var log = logger.NewConsole()

type Server struct {
	Mux *dns.ServeMux

	client    *dns.Client
	ratelimit *rate.Limiter

	upstreamLock sync.RWMutex
	upstream     []string

	dnsUDPListeners *multilisten.Pool
	dnsTCPListeners *multilisten.Pool
}

type listenerAdapter struct {
	*miekgdns.Server
}

func (this *listenerAdapter) Close() error {
	return this.Shutdown()
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

		dnsUDPListeners: multilisten.NewPool(),
		dnsTCPListeners: multilisten.NewPool(),
	}

	server.Mux.HandleFunc(".", server.handleRequest)

	return server
}

func (this *Server) upstreams() []string {
	this.upstreamLock.RLock()
	defer this.upstreamLock.RUnlock()

	result := make([]string, len(this.upstream))
	copy(result, this.upstream)

	return result
}

func (this *Server) Apply() error {
	addrs, err := general.PrivateInterfaceAddrs()
	if err != nil {
		return err
	}

	this.dnsUDPListeners.ListenAndServe(addrs, func(host string) multilisten.Listener {
		return &listenerAdapter{&miekgdns.Server{
			Addr:    net.JoinHostPort(host, "53"),
			Net:     "udp",
			Handler: this.Mux,
		}}
	})

	this.dnsTCPListeners.ListenAndServe(addrs, func(host string) multilisten.Listener {
		return &listenerAdapter{&miekgdns.Server{
			Addr:    net.JoinHostPort(host, "53"),
			Net:     "tcp",
			Handler: this.Mux,
		}}
	})

	return nil
}

func (this *Server) Shutdown() {
	this.dnsTCPListeners.Close()
	this.dnsUDPListeners.Close()
}
