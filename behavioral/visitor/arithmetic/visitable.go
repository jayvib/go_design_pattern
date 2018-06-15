package arithmetic

type Visitable interface {
	Accept(Multiplier)
}

type Multiplier interface {
	Multiply(*numberContainer)
}

type numberContainer struct {
	number int
}

func (n *numberContainer) Accept(v Multiplier) {
	v.Multiply(n)
}

type multiplier struct {
	multNum int
}

func (m *multiplier) Multiply(c *numberContainer) {
	c.number *= m.multNum
}
