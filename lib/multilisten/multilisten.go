package multilisten

import (
	"sync"
	"time"

	"github.com/vitalvas/cloudrouter/lib/logger"
)

var log = logger.NewConsole()

type Listener interface {
	ListenAndServe() error
	Close() error
}

type Pool struct {
	lock         sync.Mutex
	listeners    map[string]Listener
	shutdown     bool
	RestartDelay time.Duration
}

func NewPool() *Pool {
	return &Pool{
		listeners:    make(map[string]Listener),
		RestartDelay: time.Second,
	}
}

func (pool *Pool) ListenAndServe(hosts []string, listenerFor func(host string) Listener) {
	pool.lock.Lock()
	defer pool.lock.Unlock()

	vanished := make(map[string]bool)

	for host := range pool.listeners {
		vanished[host] = false
	}

	for _, host := range hosts {
		if _, ok := pool.listeners[host]; ok {
			delete(vanished, host)
		} else {
			log.Printf("now listening on %s", host)

			ln := listenerFor(host)
			pool.listeners[host] = ln

			go func(host string, ln Listener) {
				for !pool.shutdown {
					if err := ln.ListenAndServe(); err != nil {
						log.Printf("listener for %q died: %v", host, err)
					}

					time.Sleep(pool.RestartDelay)
				}

				pool.lock.Lock()
				defer pool.lock.Unlock()

				delete(pool.listeners, host)
			}(host, ln)
		}
	}

	for host := range vanished {
		log.Printf("no longer listening on %s", host)

		pool.listeners[host].Close()
		delete(pool.listeners, host)
	}
}

func (pool *Pool) Close() {
	pool.shutdown = true

	for host := range pool.listeners {
		pool.listeners[host].Close()
	}
}
