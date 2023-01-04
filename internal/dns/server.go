package dns

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	miekgdns "github.com/miekg/dns"
	"github.com/vitalvas/cloudrouter/lib/general"
	"github.com/vitalvas/cloudrouter/lib/logger"
	"github.com/vitalvas/cloudrouter/lib/multilisten"
	"golang.org/x/time/rate"
)

var (
	log = logger.NewConsole()

	defaultUpstream = []string{
		"8.8.8.8", "8.8.4.4", // Google
		"1.1.1.1", "1.0.0.1", // Cloudflare
		"94.140.14.14", "94.140.15.15", // AdGuard DNS
	}
)

type Server struct {
	Mux *miekgdns.ServeMux

	cfg DNS

	client    *miekgdns.Client
	ratelimit *rate.Limiter

	upstreamLock sync.RWMutex
	upstream     []string

	dnsUDPListeners *multilisten.Pool
	dnsTCPListeners *multilisten.Pool
}

type DNS struct {
	Upstream []string `json:"upstream"`
}

type listenerAdapter struct {
	*miekgdns.Server
}

func (srv *listenerAdapter) Close() error {
	return srv.Shutdown()
}

func NewServer() *Server {
	server := &Server{
		Mux: miekgdns.NewServeMux(),

		client: &miekgdns.Client{
			Timeout: 2 * time.Second,
		},

		ratelimit: rate.NewLimiter(rate.Every(time.Second), 1),

		dnsUDPListeners: multilisten.NewPool(),
		dnsTCPListeners: multilisten.NewPool(),
	}

	server.Mux.HandleFunc(".", server.handleRequest)

	return server
}

func (srv *Server) upstreams() []string {
	srv.upstreamLock.RLock()
	defer srv.upstreamLock.RUnlock()

	result := make([]string, len(srv.upstream))
	copy(result, srv.upstream)

	return result
}

func (srv *Server) Apply() error {
	b, err := os.ReadFile(filepath.Join(general.ConfigDir, "dns.json"))
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if err := json.Unmarshal(b, &srv.cfg); err != nil {
			return fmt.Errorf("cannot unmarshal config: %w", err)
		}
	}

	srv.upstreamLock.RLock()
	if len(srv.cfg.Upstream) > 0 {
		srv.upstream = srv.cfg.Upstream
	} else {
		srv.upstream = defaultUpstream
	}
	srv.upstreamLock.RUnlock()

	addrs, err := general.PrivateInterfaceAddrs()
	if err != nil {
		return err
	}

	srv.dnsUDPListeners.ListenAndServe(addrs, func(host string) multilisten.Listener {
		return &listenerAdapter{&miekgdns.Server{
			Addr:    net.JoinHostPort(host, "53"),
			Net:     "udp",
			Handler: srv.Mux,
		}}
	})

	srv.dnsTCPListeners.ListenAndServe(addrs, func(host string) multilisten.Listener {
		return &listenerAdapter{&miekgdns.Server{
			Addr:    net.JoinHostPort(host, "53"),
			Net:     "tcp",
			Handler: srv.Mux,
		}}
	})

	return nil
}

func (srv *Server) Shutdown() {
	srv.dnsTCPListeners.Close()
	srv.dnsUDPListeners.Close()
}
