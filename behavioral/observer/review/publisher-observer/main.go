package pubob

import "sync"

type Notifier interface {
	Notify(m string)
	Name() string
}

type Observers struct {
	mux       sync.Mutex
	notifiers map[string]Notifier
}

func (o *Observers) Subscribe(n ...Notifier) {
	for _, param := range n {
		if _, ok := o.notifiers[param.Name()]; !ok {
			o.notifiers[param.Name()] = param
		}
	}
}

func (o *Observers) Unsubscribe(n ...string) {
	o.mux.Lock()
	defer o.mux.Unlock()

	for _, param := range n {
		delete(o.notifiers, param)
	}
}

func (o *Observers) Trigger(message string) {

	for _, observer := range o.notifiers {
		observer.Notify(message)
	}
}
