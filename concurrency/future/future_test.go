package future

import (
	"testing"
	"sync"
	"github.com/pkg/errors"
	"time"
)

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("timeout!")
	t.Fail()
	wg.Done()
}

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success result", func(t *testing.T){
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
	})

	t.Run("Failed result", func(t *testing.T){
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "", errors.New("error occurred")
		})
	})
}
