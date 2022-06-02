package vrouter

import (
	"log"
	"os"

	"github.com/vitalvas/cloudrouter/app/vrouter/sysctl"
)

func (this *VRouter) ApplySysctl() {
	for key, val := range sysctl.Params(this.config.System.Sysctl) {
		current, err := sysctl.Get(key)
		if err != nil {
			if err == os.ErrNotExist {
				continue
			} else {
				log.Fatal(err)
			}
		}

		if val != current {
			log.Println("changing", key, "from", current, "to", val)
			sysctl.Set(key, val)
		}
	}
}
