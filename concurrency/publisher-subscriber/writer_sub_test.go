package publisher_subscriber

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	testintFunc func(string)
}

func (m *mockWriter) Write(p []byte) (int, error) {
	m.testintFunc(string(p))
	return len(p), nil
}


func TestSubscriber(t *testing.T) {
	sub := NewWriterSubscriber(0, nil)
	msg := "Hello"
	var wg sync.WaitGroup
	wg.Add(1)
	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testintFunc: func(s string) {
			defer wg.Done()
			if !strings.Contains(s, msg) {
				t.Fatal(fmt.Errorf("incorrect string: %s", s))
			}
		},
	}
	err := sub.Notify(msg)
	if err != nil {
		wg.Done()
		t.Error(err)
	}
	wg.Wait()
	sub.Close()
}