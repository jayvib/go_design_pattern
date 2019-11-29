package description

import "container/list"

type State struct {
	Description string
}

type Memento struct {
	state State
}

type originator struct {
	s State
}

func (o *originator) Snapshot() Memento {
	return Memento{state:o.s}
}

func (o *originator) Extract(m Memento) {
	o.s = m.state
}

type caretaker struct {
	l *list.List
}

func (c *caretaker) Push(m Memento) {
	c.l.PushBack(m)
}

func (c *caretaker) Pop() Memento {
	e := c.l.Back()
	c.l.Remove(e)

	return e.Value.(Memento)
}
