package main

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before looop, current count is %d\n", singleton.GetCount())

	var val int
	if val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	singleton.Stop()
	fmt.Println(val)
}
