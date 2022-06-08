package dns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/miekg/dns"
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
	Mux *dns.ServeMux

	cfg DNS

	client    *dns.Client
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

func (this *listenerAdapter) Close() error {
	return this.Shutdown()
}

func NewServer() *Server {
	server := &Server{
		Mux: dns.NewServeMux(),

		client: &dns.Client{
			Timeout: 2 * time.Second,
		},

		ratelimit: rate.NewLimiter(rate.Every(time.Second), 1),

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
	b, err := ioutil.ReadFile(filepath.Join(general.ConfigDir, "dns.json"))
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if err := json.Unmarshal(b, &this.cfg); err != nil {
			return fmt.Errorf("cannot unmarshal config: %w", err)
		}
	}

	this.upstreamLock.RLock()
	if len(this.cfg.Upstream) > 0 {
		this.upstream = this.cfg.Upstream
	} else {
		this.upstream = defaultUpstream
	}
	this.upstreamLock.RUnlock()

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
