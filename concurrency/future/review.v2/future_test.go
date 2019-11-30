package main

import (
	"sync"
	"testing"
	"time"
)

func timeout(wg *sync.WaitGroup, t *testing.T) {
	time.Sleep(5 * time.Second)
	t.Fail()
	wg.Done()
}

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		// In order to avoid deadlock
		go timeout(&wg, t)

		future.Success(func(s string) {
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "Hello, World", nil
		})

		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {
	})
}
