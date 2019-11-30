package worker

import (
	"fmt"
	"sync"
)

type WorkLauncher interface {
	SetID(id int)
	Launch(wg *sync.WaitGroup, in chan Request)
}

func NewWorker(handler func(interface{}) (interface{}, error), id int) WorkLauncher {
	return &Worker{handler: handler, id: id}
}

type Worker struct {
	id      int
	handler func(interface{}) (interface{}, error)
}

func (w *Worker) SetID(id int) {
	w.id = id
}

func (w *Worker) Launch(wg *sync.WaitGroup, in chan Request) {
	go func() {
		fmt.Printf("Worker: %d up\n", w.id)
		defer func() {
			wg.Done()
			fmt.Printf("Worker %d exiting\n", w.id)
		}()
		for req := range in {
			res, _ := w.handler(req.Data)
			fmt.Println("Sending:", res, w.id)
		}
	}()
}
