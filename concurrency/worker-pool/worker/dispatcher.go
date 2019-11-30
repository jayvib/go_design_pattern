package worker

import (
	"fmt"
	"sync"
	"time"
)

type Dispatcher interface {
	Dispatch(handler func(interface{}) (interface{}, error))
	MakeRequest(r Request)
	Stop()
}

func NewDispatcher(nworkers int) Dispatcher {
	return &dispatcher{
		inCh:       make(chan Request), // maybe use buffered channel?
		numWorkers: nworkers,
		wg:         &sync.WaitGroup{},
	}
}

type dispatcher struct {
	inCh       chan Request
	numWorkers int
	wg         *sync.WaitGroup
}

func (d *dispatcher) Dispatch(handler func(interface{}) (interface{}, error)) {
	for i := 0; i < d.numWorkers; i++ {
		d.wg.Add(1)
		worker := NewWorker(handler, i)
		worker.Launch(d.wg, d.inCh) // should I use context here?
	}
}

func (d *dispatcher) launchWorker(wg *sync.WaitGroup, w WorkLauncher) {
	w.Launch(wg, d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
}

func (d *dispatcher) Stop() {
	close(d.inCh)
	d.wg.Wait()
}
