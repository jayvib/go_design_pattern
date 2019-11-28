package main

import (
	"sync"
	"time"
)

func main() {
	channel := make(chan string, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func(wg *sync.WaitGroup, ch chan<- string) {
		time.Sleep(3 * time.Second)
		ch <- "Hello Jayson"
		//close(ch)
		println("finished goroutine 1")
		defer wg.Done()
	}(&wg, channel)

	go func(wg *sync.WaitGroup, ch chan<- string) {
		time.Sleep(4 * time.Second)
		ch <- "Hello Channel Concurrency World!!!"
		println("finished goroutine 2")
		defer wg.Done()
	}(&wg, channel)

	println(<-channel)
	println(<-channel)
	wg.Wait()
	println("existing main")
}
