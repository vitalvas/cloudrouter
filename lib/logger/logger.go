package logger

import (
	"io"
	"log"
	"os"
)

type nonBlockingWriter struct {
	W chan<- string
}

func (w *nonBlockingWriter) Write(p []byte) (n int, _ error) {
	select {
	case w.W <- string(p):
	default:
		// channel unavailable, ignore
	}

	return len(p), nil
}

func New() *log.Logger {
	w := io.Discard

	if console, err := os.OpenFile("/dev/console", os.O_RDWR, 0600); err == nil {
		ch := make(chan string, 1)

		go func() {
			for buf := range ch {
				console.Write([]byte(buf))
			}
		}()

		w = &nonBlockingWriter{W: ch}
	}

	return log.New(io.MultiWriter(os.Stderr, w), "", log.LstdFlags|log.Lshortfile)
}
