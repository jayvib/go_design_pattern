package multi_logger

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type MultiLogger struct {
	loggers []ChainLogger
}

func (m *MultiLogger) Do(s string) {
	for _, logger := range m.loggers {
		logger.Next(s)
	}
}

// NewMultiLogger takes a type that satisfies ChainLogger then return a MultiLogger type
func NewMultiLogger(loggers ...ChainLogger) *MultiLogger {
	ml := new(MultiLogger)
	for _, l := range loggers {
		ml.loggers = append(ml.loggers, l)
	}
	return ml
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (se *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if se.NextChain != nil {
			se.NextChain.Next(s)
		}
		return
	}
	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}

}

type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}

	if c.NextChain != nil {
		c.NextChain.Next(s)
	}
}
