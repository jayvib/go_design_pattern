package string_processor

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"sync"
)

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

// Request is the data model
type Request struct {
	Data interface{}
	Handler RequestHandler
}

type RequestHandler func(data interface{})

// WorkerLauncher represents the worker pool instance
// This will standby until the new data was come in to
// channel.
type WorkerLauncher interface {
	LaunchWorker(in <-chan Request)
}

func NewDispatcher(ctx context.Context, size int) Dispatcher {
	return &dispatcher{
		in: make(chan Request, size),
		ctx: ctx,
	}
}

// Dispatcher
type Dispatcher interface {
	// LaunchWorker will launch an instance of the worker.
	// LaunchWorker will not know whether the WorkerLauncher
	// implements a channel it will still detach from the main
	// and the channel in the WorkerLauncher will block until
	// the data come.
	LaunchWorker(w WorkerLauncher)
	// MakeRequest will send the request to the underlying channel
	// of the object
	MakeRequest(r Request)
	// Stop will close the channel and will signal the worker
	// that there will be no more data to be process.
	Stop()
}

// --------Implementors----------

// Implement the Dispatcher

type dispatcher struct {
	in chan Request
	ctx context.Context
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	go w.LaunchWorker(d.in)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.in <- r:
	case <-d.ctx.Done():
		return
	}
}

func (d *dispatcher) Stop() {
	close(d.in)
}

