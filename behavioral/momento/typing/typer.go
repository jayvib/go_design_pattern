package typing

type state struct {
	s string
}

type memento struct {
	state state
}

type originator struct {
	memento memento
}

func (o *originator) getmemento() memento {
	return memento{}
}

func (o *originator) extractAndStoreState(m memento) {
	o.memento = m
}

type careTaker struct {
	mementos []memento
}

func (c *careTaker) add(m memento) {
	c.mementos = append(c.mementos, m)
}

func (c *careTaker) memento(i int) (memento, error) {
	return memento{}, nil
}
