package momento

type state struct {
	decription string
}

type originator struct {
	s state
}

type momento struct {
	s state
}

func (o originator) NewMomento() momento {
	return momento{s: o.s}
}

type caretaker struct {
	momentos []momento
}

func (c caretaker) Push(m momento) {
	c.momentos = append(c.momentos, m)
}

func (c caretaker) Pop() momento {
	m := c.momentos[len(c.momentos)-1]
	c.momentos = c.momentos[:len(c.momentos)-2]
	return m
}

// use facade
type stateSaver struct {
	o originator
	c caretaker
}

func (s stateSaver) NewOriginator(o originator) {
	s.c.Push(s.o.NewMomento())
	s.o = o

}

func (s stateSaver) SaveMomento() {
	s.c.Push(s.o.NewMomento())
}

func (s stateSaver) RestoreMomento() {
	s.o.s = s.c.Pop().s
}
