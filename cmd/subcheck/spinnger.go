package main

import (
	"fmt"
	"time"
)

type Spinner struct {
	stop chan struct{}
}

func NewSpinner(msg string) *Spinner {
	s := &Spinner{stop: make(chan struct{})}

	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		for {
			select {
			case <-s.stop:
				fmt.Print("\r")
				return
			default:
				fmt.Printf("\r%s %s", msg, frames[i%len(frames)])
				time.Sleep(80 * time.Millisecond)
				i++
			}
		}
	}()

	return s
}

func (s *Spinner) Stop(successMsg string) {
	close(s.stop)
}
