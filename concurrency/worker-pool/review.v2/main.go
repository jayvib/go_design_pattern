package main

import (
	"fmt"
	"sync"
)

func main() {
	bufferSize := 100
	dispatcher := NewDispatcher(bufferSize)

	workers := 3

	for i := 0; i < workers; i++ {
		w := &PreffixSuffixWorker{
			prefix: fmt.Sprintf("Worker ID: %d -> ", i),
			suffix: " World",
			id:     i,
		}
		dispatcher.LaunchWorker(w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()
	wg.Wait()
}
