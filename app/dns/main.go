package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	miekgdns "github.com/miekg/dns"
	"github.com/vitalvas/cloudrouter/lib/dns"
	"github.com/vitalvas/cloudrouter/lib/general"
	"github.com/vitalvas/cloudrouter/lib/logger"
	"github.com/vitalvas/cloudrouter/lib/multilisten"
)

var (
	log = logger.NewConsole()

	dnsUDPListeners = multilisten.NewPool()
	dnsTCPListeners = multilisten.NewPool()
)

type listenerAdapter struct {
	*miekgdns.Server
}

func (a *listenerAdapter) Close() error {
	return a.Shutdown()
}

func updateListeners(mux *miekgdns.ServeMux) error {
	addrs, err := general.PrivateInterfaceAddrs()
	if err != nil {
		return err
	}

	dnsUDPListeners.ListenAndServe(addrs, func(host string) multilisten.Listener {
		return &listenerAdapter{&miekgdns.Server{
			Addr:    net.JoinHostPort(host, "53"),
			Net:     "udp",
			Handler: mux,
		}}
	})

	dnsTCPListeners.ListenAndServe(addrs, func(host string) multilisten.Listener {
		return &listenerAdapter{&miekgdns.Server{
			Addr:    net.JoinHostPort(host, "53"),
			Net:     "tcp",
			Handler: mux,
		}}
	})

	return nil
}

func main() {
	srv := dns.NewServer()

	if err := updateListeners(srv.Mux); err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR1, os.Interrupt)

	for sig := range ch {
		if sig == syscall.SIGUSR1 {
			if err := updateListeners(srv.Mux); err != nil {
				log.Printf("updateListeners: %v", err)
			}
		} else if sig == os.Interrupt {
			break
		}
	}

	log.Println("shutdown")

	dnsTCPListeners.Close()
	dnsUDPListeners.Close()
}
