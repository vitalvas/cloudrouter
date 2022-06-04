package netwatch

import (
	"crypto/rand"
	"fmt"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func ICMP(task Task) {
	token := make([]byte, 16)
	rand.Read(token)

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: token,
		},
	}

	fmt.Printf("msg: %v\n", msg)
}
