package icmp

import (
	"fmt"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func ICMP() {
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}

	fmt.Printf("msg: %v\n", msg)
}
