package loader

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
)

type Loader struct {
	log   *log.Logger
	token string
}

func New() *Loader {
	return &Loader{}
}

func (lo *Loader) SetLogger(l *log.Logger) {
	lo.log = l
}

func (lo *Loader) Apply() error {
	if len(lo.token) < 10 {
		lo.setSessionToken()
	}

	return nil
}

func (lo *Loader) Shutdown() {
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
