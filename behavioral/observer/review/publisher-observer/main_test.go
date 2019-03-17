package pubob

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyNotifier struct {
	name string
	output string
}

func (d *dummyNotifier) Notify(m string) {
	d.output = m
}
func (d *dummyNotifier) Name() string { return d.name }

func TestObservers(t *testing.T) {
	var observers *Observers
	t.Run("Subscribe", func(t *testing.T){
		observers = &Observers{ notifiers: make(map[string]Notifier) }

		notifier1 := &dummyNotifier{ name: "notifier1" }
		notifier2 := &dummyNotifier{ name: "notifier2" }
		notifier3 := &dummyNotifier{ name: "notifier3" }
		notifier4 := &dummyNotifier{ name: "notifier4" }

		observers.Subscribe(notifier1, notifier2, notifier3, notifier4)

		l := len(observers.notifiers)
		assert.Equal(t, 4, len(observers.notifiers),
			"Expecting %d notifiers but got %d", 4, l)
	})

	t.Run("Unsubscribe", func(t *testing.T){
		notifierNames := []string{
			"notifier1",
			"notifier2",
		}

		observers.Unsubscribe(notifierNames...)

		l := len(observers.notifiers)
		assert.Equal(t, 2, l, "Expecting %d but got %d", 2, l)
	})

	t.Run("Trigger", func(t *testing.T){
		observers = &Observers{ notifiers: make(map[string]Notifier) }

		notifier1 := &dummyNotifier{ name: "notifier1" }
		notifier2 := &dummyNotifier{ name: "notifier2" }
		observers.Subscribe(notifier1, notifier2)
		message := "Hello Event Driven World!"
		observers.Trigger(message)

		assert.Equal(t, message, notifier1.output)
		assert.Equal(t, message, notifier2.output)
	})
}
