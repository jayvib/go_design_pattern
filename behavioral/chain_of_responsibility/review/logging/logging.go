package logging

import (
	"io"

	"github.com/sirupsen/logrus"
)

// The object that will implement this is completely
// independent from others... in other words..
// Object has a single responsibility method Next(string)
type ChainLogger interface {
	Next(string)
}

type consoleLogger struct {
	chain ChainLogger
}

func (cl *consoleLogger) Next(str string) {
	logrus.Printf("Console: %s\n", str)
	if cl.chain != nil {
		cl.chain.Next(str)
	}
}

type writeLogger struct {
	chain ChainLogger
	w     io.Writer
}

func (wl *writeLogger) Next(str string) {
	wl.w.Write([]byte(str))

	if wl.chain != nil {
		wl.chain.Next(str)
	}
}
