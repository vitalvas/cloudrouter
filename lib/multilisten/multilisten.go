package multilisten

import (
	"sync"

	"github.com/vitalvas/cloudrouter/lib/logger"
)

var log = logger.NewConsole()

type Listener interface {
	ListenAndServe() error
	Close() error
}

type Pool struct {
	lock      sync.Mutex
	listeners map[string]Listener
}

func NewPool() *Pool {
	return &Pool{
		listeners: make(map[string]Listener),
	}
}

func (this *Pool) ListenAndServe(hosts []string, listenerFor func(host string) Listener) {
	this.lock.Lock()
	defer this.lock.Unlock()

	vanished := make(map[string]bool)

	for host := range this.listeners {
		vanished[host] = false
	}

	for _, host := range hosts {
		if _, ok := this.listeners[host]; ok {
			delete(vanished, host)
		} else {
			log.Printf("now listening on %s", host)

			ln := listenerFor(host)
			this.listeners[host] = ln

			go func(host string, ln Listener) {
				if err := ln.ListenAndServe(); err != nil {
					log.Printf("listener for %q died: %v", host, err)
				}

				this.lock.Lock()
				defer this.lock.Unlock()

				delete(this.listeners, host)
			}(host, ln)
		}
	}

	for host := range vanished {
		log.Printf("no longer listening on %s", host)

		this.listeners[host].Close()
		delete(this.listeners, host)
	}
}

func (this *Pool) Close() {
	for host := range this.listeners {
		this.listeners[host].Close()
	}
}
