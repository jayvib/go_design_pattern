package event_driven

import "sync"

const eventChanBufferSize = 10

type eventChannel chan *Event

type Eventer interface {
	Events() (eventNames map[string]string)
	Event(name string) string
	AddEvent(name string)
	DeleteEvent(name string)
	Publish(name string, data interface{})
	Subscribe() (events eventChannel)
	Unsubscribe(events eventChannel)
	On(name string, f func(s interface{})) (err error)
	Once(name string, f func(s interface{})) (err error)
}

type Event struct {
	Name string
	Data interface{}
}

func NewEvent(name string, data interface{}) *Event {
	return &Event{
		Name: name,
		Data: data,
	}
}

// eventer is an implementation for Eventer
type eventer struct {
	eventnames map[string]string
	in eventChannel
	outs map[eventChannel]eventChannel
	eventsMutex sync.Mutex
}

func (e *eventer) Events() (eventNames map[string]string) {
	return e.eventnames
}

func (e *eventer) Event(name string) (eventName string) {
	return e.eventnames[name]
}

func (e *eventer) AddEvent(name string) {
	e.eventnames[name] = name
}

func (e *eventer) DeleteEvent(name string) {
	delete(e.eventnames, name)
}

func (e *eventer) Publish(name string, data interface{}) {
	event := NewEvent(name, data)
	e.in <- event
}

func (e *eventer) Subscribe() eventChannel {
	e.eventsMutex.Lock()
	defer e.eventsMutex.Unlock()
	out := make(eventChannel, eventChanBufferSize)
	e.outs[out] = out
	return out
}

func (e *eventer) Unsubscribe(events eventChannel) {
	e.eventsMutex.Lock()
	defer e.eventsMutex.Unlock()
	delete(e.outs, events)
}

func (e *eventer) On(n string, f func(s interface{})) (err error) {
	out := e.Subscribe()
	go func() {
		for {
			select {
			case evt := <-out:
				if evt.Name == n {
					f(evt.Data)
				}
			}
		}
	}()
	return
}

func (e *eventer) Once(n string, f func(s interface{})) (err error) {
	out := e.Subscribe()
	go func() {
	ProcessEvents:
		for evt := range out {
			if evt.Name == n {
				f(evt.Data)
			}
			e.Unsubscribe(out)
			break ProcessEvents
		}
	}()
	return
}