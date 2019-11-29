package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var (
	MaxWorkers = 5
	MaxQueue   = 5
)

type Processor struct {
	// Process the data
}

func (p *Processor) Process() {
	duration := time.Duration(rand.Intn(10)) * time.Second
	fmt.Println("Processing start...")
	time.Sleep(duration)
	fmt.Println("Processing end...")
}

type Job struct {
	Processor Processor
}

var JobQueue chan Job

func NewWorker(pool chan chan Job) *Worker {
	w := &Worker{
		WorkerPool: pool,
		JobChan:    make(chan Job),
	}

	ctx, cancel := context.WithCancel(context.Background())
	w.ctx = ctx
	w.cancel = cancel
	return w
}

type Worker struct {
	WorkerPool chan chan Job
	JobChan    chan Job
	ctx        context.Context
	cancel     context.CancelFunc
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChan // register this channel
			select {
			case <-w.ctx.Done():
				return
			case job := <-w.JobChan:
				job.Processor.Process()
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.cancel()
	}()
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, maxWorkers),
		MaxWorkers: maxWorkers,
	}
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChan := <-d.WorkerPool
				jobChan <- job
			}(job)
		}
	}
}

func main() {
	JobQueue = make(chan Job, MaxQueue)
	dispatcher := NewDispatcher(MaxWorkers)
	dispatcher.Run()

	for i := 0; i < 20; i++ {
		JobQueue <- Job{Processor: Processor{}}
	}
	
}
