package time_command

import (
	"fmt"
	"time"
)

type Command interface {
	Info() string
}

type TimePassed struct {
	Start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.Start).String() // time.Since function returns the time elapsed since the time stored in the provided paramter.
}

type ChainLogger interface {
	Next(command Command)
}

type Logger struct {
	NextChain ChainLogger
}

func (f *Logger) Next(c Command) {
	time.Sleep(time.Second)
	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}


