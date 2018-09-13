package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}

	for i := 0; i < 10; i++ {
		go func(i int) { // deploy 10 goroutines
			fmt.Printf("Its %d turn\n", i)
			//counter.Lock() // lock the value
			//defer counter.Unlock() // then unlock afterwards
			counter.value++ // increment the value while it is lock by the mutex
		}(i)
	}
	time.Sleep(time.Second)

	counter.Lock()
	defer counter.Unlock()
	println(counter.value)
}
