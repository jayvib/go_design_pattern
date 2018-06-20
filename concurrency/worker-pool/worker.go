package worker

import (
	"log"
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
				log.Fatal("invalid casting to string")
			}
		},
	}
}

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(id int, w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(5 * time.Second):
		return
	}
}
