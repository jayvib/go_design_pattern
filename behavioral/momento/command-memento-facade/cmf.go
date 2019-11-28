package main

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

type Memento struct {
	memento Command
}

type originator struct {
	Command Command
}

func (o *originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

type careTaker struct {
	mementoStack []Memento
}

func (c *careTaker) Add(m Memento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *careTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		memento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[:len(c.mementoStack)-1]
		return memento
	}
	return Memento{}
}

type MementoFacade struct {
	originator originator
	careTaker  careTaker
}

func (m *MementoFacade) SaveSetting(c Command) {
	m.originator.Command = c
	m.careTaker.Add(m.originator.NewMemento())
}

func (m *MementoFacade) RestoreSetting() Command {
	m.originator.ExtractAndStoreCommand(m.careTaker.Pop())
	return m.originator.Command
}
