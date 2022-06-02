package main

import (
	"log"

	"github.com/vitalvas/cloudrouter/app/vrouter/vrouter"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	vrouter.NewVRouter().Execute()
}
