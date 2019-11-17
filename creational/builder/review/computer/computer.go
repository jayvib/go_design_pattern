package computer

// WISDOM:
// Try to avoid the builder pattern when you are not completely sure that the algorithm
// is going to be more or less stable because any small change in this interface
// will affect all you builders and it could be awkward if you add a new method 
// that some of your builders need and others Builders do not.

type BuildProcess interface {
	SetMonitor() BuildProcess
	SetKeyboard() BuildProcess
	SetMouse() BuildProcess
	SetType() BuildProcess
	GetComputer() ComputerProduct
}

// This is the product that will be construct
type ComputerProduct struct {
	monitor  string
	keyboard string
	mouse    string
	t string // computer type
}

// Director is an object that is responsible for
// building the Process

func NewDirector(b BuildProcess) *Director {
	return &Director{
		builder: b,
	}
}

type Director struct {
	builder BuildProcess
}

func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}

func (c *Director) Build() {
	c.builder.SetMonitor().SetKeyboard().SetMouse().SetType()
}

// Builders are the objects that responsible how the
// moniter will be built

func NewLaptopBuilder() BuildProcess {
	return &LaptopBuilder{}
}

type LaptopBuilder struct {
	product      ComputerProduct
	computerType string
}

func (l *LaptopBuilder) SetMonitor() BuildProcess {
	l.product.monitor = "Dell"
	return l
}
func (l *LaptopBuilder) SetKeyboard() BuildProcess {
	l.product.keyboard = "HP"
	return l
}
func (l *LaptopBuilder) SetMouse() BuildProcess {
	l.product.mouse = "Lenovo"
	return l
}
func (l *LaptopBuilder) SetType() BuildProcess{
	l.product.t = "Laptop"
}
func (l *LaptopBuilder) GetComputer() ComputerProduct {
	return l.product
}
