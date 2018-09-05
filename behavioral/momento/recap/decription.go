package recap

import "fmt"

type State struct {
	Description string
}

type momento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMomento() momento {
	return momento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m momento) {
	o.state = m.state
}

type careTaker struct {
	momentoList []momento
}

func (c *careTaker) Add(m momento) {
	c.momentoList = append(c.momentoList, m)
}

func (c *careTaker) Momento(i int) (momento, error) {
	if len(c.momentoList) < i || i < 0 {
		return momento{}, fmt.Errorf("Index not found\n")
	}
	return c.momentoList[i], nil
}


