package loader

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
)

var apps = []string{
	"/apps/bin/dns",
	"/apps/bin/neighbor",
}

type Loader struct {
	log     *log.Logger
	token   string
	runOnce sync.Once

	apps map[string]*os.Process
}

func New() *Loader {
	return &Loader{
		apps: make(map[string]*os.Process),
	}
}

func (lo *Loader) SetLogger(l *log.Logger) {
	lo.log = l
}

func (lo *Loader) Apply() error {
	if len(lo.token) < 10 {
		lo.setSessionToken()
	}

	lo.runOnce.Do(func() {
		for _, name := range apps {
			go lo.execute(name)
		}
	})

	return nil
}

func (lo *Loader) Shutdown() {
	for name, child := range lo.apps {
		if err := child.Signal(syscall.SIGINT); err != nil {
			lo.log.Printf("error shutdown process '%s': %v", name, err)
		}
	}
}

func (lo *Loader) setSessionToken() {
	const size = 24

	token := make([]byte, size)
	if _, err := rand.Read(token); err != nil {
		lo.log.Fatal(err)
	}

	lo.token = base64.RawURLEncoding.EncodeToString(token)
}

func (lo *Loader) generateEnvData() []string {
	return []string{
		"CR_BROKER=tcp://[::1]:1883",
		fmt.Sprintf("CR_TOKEN=%s", lo.token),
	}
}
