package netwatch

import (
	"net"
	"time"

	"github.com/digineo/go-ping"
)

type Task struct {
	Host     *net.IPAddr
	Interval time.Duration
	Timeout  time.Duration
	Attempts int
}

func NewTask() *Task {
	return &Task{
		Interval: time.Minute,
		Timeout:  time.Second,
		Attempts: 4,
	}
}

func (task *Task) Execute() (time.Duration, error) {
	pinger, err := ping.New("0.0.0.0", "::")
	if err != nil {
		return 0, err
	}

	defer pinger.Close()

	if pinger.PayloadSize() != 56 {
		pinger.SetPayloadSize(56)
	}

	rtt, err := pinger.PingAttempts(task.Host, task.Timeout, task.Attempts)
	if err != nil {
		return 0, err
	}

	return rtt, nil
}
