package icmp

import (
	"net"
	"time"
)

type Task struct {
	Host     net.IPAddr
	Interval time.Duration
	Timeout  time.Duration
}
