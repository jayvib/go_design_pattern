package main

import (
	"fmt"
	"strings"
)

type WorkerLauncher interface {
	LaunchWorker(in chan *Request)
}

type PreffixSuffixWorker struct {
	id     int
	prefix string
	suffix string
}

func (w *PreffixSuffixWorker) LaunchWorker(in chan *Request) {
	w.applyPrefix(w.applySuffix(w.uppercase(in)))
}

func (w *PreffixSuffixWorker) uppercase(in <-chan *Request) <-chan *Request {
	out := make(chan *Request)

	go func() {
		defer close(out)
		for msg := range in {
			s := msg.Data.(string)
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
	}()

	return out
}

func (w *PreffixSuffixWorker) applySuffix(in <-chan *Request) <-chan *Request {
	out := make(chan *Request)
	go func() {
		defer close(out)
		for msg := range in {
			s := msg.Data.(string)
			msg.Data = fmt.Sprintf("%s%s", w.suffix, s)
			out <- msg
		}
	}()
	return out
}

func (w *PreffixSuffixWorker) applyPrefix(in <-chan *Request) {
	go func() {
		for msg := range in {
			s := msg.Data.(string)
			msg.Data = fmt.Sprintf("%s%s", s, w.prefix)
			msg.Handler(msg.Data)
		}
	}()
}
