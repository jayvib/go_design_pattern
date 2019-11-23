package main

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID      int
	Message string
}

func (p *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message '%s' received\n", p.ID, m)
	p.Message = m
}

func TestSubject(t *testing.T) {
	testObserver1 := &TestObserver{1, ""}
	testObserver2 := &TestObserver{2, ""}
	testObserver3 := &TestObserver{3, ""}
	publisher := Publisher{}

	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(testObserver1)
		publisher.AddObserver(testObserver2)
		publisher.AddObserver(testObserver3)

		if len(publisher.ObserversList) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(testObserver2)

		if len(publisher.ObserversList) != 2 {
			t.Errorf("The size of the observer is not the size expected\n")
		}
		for _, observer := range publisher.ObserversList {
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
			}
			if testObserver.ID == 2 {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {
		if len(publisher.ObserversList) == 0 {
			t.Error("Nothing to test.")
		}
		for _, observer := range publisher.ObserversList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fatal()
			}
			if printObserver.Message != "" {
				t.Errorf("The observer's message werent empty.")
			}
		}

		message := "Hello World!"
		publisher.NotifyObservers(message)
		for _, observer := range publisher.ObserversList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}
			if printObserver.Message != message {
				t.Errorf("Expected message on observer %d was not met. got: %s", printObserver.ID, printObserver.Message)
			}
		}
	})
}
