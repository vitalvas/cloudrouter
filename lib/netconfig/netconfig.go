package netconfig

import "github.com/vitalvas/cloudrouter/lib/logger"

var log = logger.NewConsole()

func Apply() {
	if err := applySysctl(); err != nil {
		log.Println("sysctl:", err)
	}
}
