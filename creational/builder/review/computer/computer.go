package main

type BuildProcess interface {
	SetMonitor() BuildProcess
	SetKeyboard() BuildProcess
	SetMouse() BuildProcess
	GetComputer() ComputerProduct
}

// This is the product that will be construct
type ComputerProduct struct {
	monitor  string
	keyboard string
	mouse    string
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
	//d.builder = b
}

func (c *Director) Build() {
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
	return nil
}
func (l *LaptopBuilder) SetKeyboard() BuildProcess {
	return nil
}
func (l *LaptopBuilder) SetMouse() BuildProcess {
	return nil
}
func (l *LaptopBuilder) GetComputer() ComputerProduct {
	return ComputerProduct{}
}
