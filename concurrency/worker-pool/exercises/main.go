package main

import (
	"fmt"
	sp "go_design_pattern/concurrency/worker-pool/exercises/string-processor"
	"golang.org/x/net/context"
	"strings"
	"sync"
	"time"
)

type PrefixSuffixStringLaunchWorker struct {
	prefixS string
	suffixS string
	id int
}

func (p *PrefixSuffixStringLaunchWorker) LaunchWorker(in chan sp.Request) {
	fmt.Printf("Worker %d is up...\n", p.id)
	p.suffix(p.prefix(p.uppercase(in)))
}

// Create a pipeline.
// Steps:
// 1. Uppercase the string
// 2. Put prefix to the string
// 3. Put suffix to the string
func (p *PrefixSuffixStringLaunchWorker) uppercase(in <-chan sp.Request) (<-chan sp.Request) {
	outChan := make(chan sp.Request)
	go func() {
		defer close(outChan)
		for v := range in {
			time.Sleep(2 * time.Second)
			s, ok := v.Data.(string)
			if !ok {
				v.Handler(nil)
				return
			}
			trans := strings.ToUpper(s)
			output := fmt.Sprintf("Request_ID: %d -> %s", p.id, trans)
			v.Data = output
			outChan <- v
		}
	}()
	return outChan
}

func (p *PrefixSuffixStringLaunchWorker) prefix(in <-chan sp.Request) (<-chan sp.Request) {
	outChan := make(chan sp.Request)
	go func() {
		defer close(outChan)
		for r := range in {
			s, ok := r.Data.(string)
			if !ok {
				r.Handler(nil)
				return
			}
			r.Data = fmt.Sprintf("%s-%s", p.prefixS, s)
			outChan <- r
		}
	}()
	return outChan
}

func (p *PrefixSuffixStringLaunchWorker) suffix(in <-chan sp.Request) {
	for r := range in {
		s, ok := r.Data.(string)
		if !ok {
			r.Handler(nil)
			return
		}
		trans := fmt.Sprintf("%s-%s", s, p.suffixS)
		r.Handler(trans)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// Launch the dispatcher
	var dptcher sp.Dispatcher = sp.NewDispatcher(ctx, 100)
	workerSize := 5
	for i := 1; i < workerSize; i++ {
		var lworker sp.WorkerLauncher = &PrefixSuffixStringLaunchWorker{
			prefixS: "Go",
			suffixS: "Gopher",
			id: i,
		}
		dptcher.LaunchWorker(lworker)
	}

	request := 100
	var wg sync.WaitGroup
	wg.Add(request)
	for i := 0; i < request; i++ {
		req := sp.NewStringRequest(
			"Request_ID: %d -> Hello Gopher!", i, &wg,
		)

		dptcher.MakeRequest(req)
	}
	dptcher.Stop()
	wg.Wait()
}
