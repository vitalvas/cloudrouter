package dns

import (
	"net"

	"github.com/miekg/dns"
)

func (srv *Server) handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	for idx, u := range srv.upstreams() {
		if _, _, err := net.SplitHostPort(u); err != nil {
			u = net.JoinHostPort(u, "53")
		}

		in, _, err := srv.client.Exchange(r, u)
		if err != nil {
			srv.log.Printf("resolving %v failed: %v", r.Question, err)

			continue
		}

		w.WriteMsg(in)

		if idx > 0 {
			// re-order srv upstream to the front of srv.upstream
			srv.upstreamLock.Lock()

			if srv.upstream[idx] == u {
				srv.upstream = append(append([]string{u}, srv.upstream[:idx]...), srv.upstream[idx+1:]...)
			}

			srv.upstreamLock.Unlock()
		}
	}
}
