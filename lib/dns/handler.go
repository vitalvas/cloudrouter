package dns

import "github.com/miekg/dns"

func (this *Server) handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	for idx, u := range this.upstreams() {
		in, _, err := this.client.Exchange(r, u)
		if err != nil {
			if this.ratelimit.Allow() {
				log.Printf("resolving %v failed: %v", r.Question, err)
			}

			continue
		}

		w.WriteMsg(in)

		if idx > 0 {
			// re-order this upstream to the front of this.upstream
			this.upstreamLock.Lock()

			if this.upstream[idx] == u {
				this.upstream = append(append([]string{u}, this.upstream[:idx]...), this.upstream[idx+1:]...)
			}

			this.upstreamLock.Unlock()
		}
	}
}
