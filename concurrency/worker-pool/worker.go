package main

import (
	"fmt"
	"go_design_pattern/concurrency/worker-pool/review"
)

func main() {
	bufferSize := 100
	var dispatcher = review.NewDispatcher(bufferSize)
	workers := 5
	for i := 0; i < workers; i++ {
		var w review.WorkerLauncher = &review.PreffixSuffixWorker{
			i,
			fmt.Sprintf("WorkerID: %d -> ", i),
			" World",
		}

		dispatcher.LaunchWorker(w)
	}

	//requests := 100
	//var wg sync.WaitGroup
	//wg.Add(requests)
	//
	//for i := 0; i < requests; i++ {
	//	req := review.NewStringRequest("(Msg_id: %d -> Hello", i, &wg)
	//	dispatcher.MakeRequest(req)
	//}
	//dispatcher.Stop()
	//wg.Wait()
}
