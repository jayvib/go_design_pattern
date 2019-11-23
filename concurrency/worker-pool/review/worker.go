package review

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	return Request{
		Data: "Hello",
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Panic("invalid casting to string")
			}
			fmt.Println(s)
		},
	}
}

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(request Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh) // this will detach even tough the implementor is using a channel
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}

type PreffixSuffixWorker struct {
	Id      int
	PrefixS string
	SuffixS string
}

func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	fmt.Printf("Worker %d is up...\n", w.Id)
	w.prefix(w.append(w.uppercase(in))) // pipeline
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		defer close(out)
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		defer close(out)
		for msg := range in {
			upperCaseString, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				return
			}
			msg.Data = fmt.Sprintf("%s%s", upperCaseString, w.SuffixS)
			out <- msg
		}
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			upperCasedStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				return
			}
			msg.Handler(fmt.Sprintf("%s%s", w.PrefixS, upperCasedStringWithSuffix))
		}
	}()
}
