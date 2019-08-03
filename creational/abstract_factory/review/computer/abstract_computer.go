package computer

import "errors"

type Computer interface {
	Display() string
	Mouse() string
	Keyboard() string
	CPU() string
}

type ComputerFactory interface {
	NewComputer(c computer) Computer
}

type computer int

const (
	HPLaptop computer = iota
	DellLaptop

	HPDesktop
	DellDesktop
)

type laptopFactory struct{}

func (l laptopFactory) NewComputer(c computer) Computer {
	switch c {
	case HPLaptop:
		return hpLaptop{}
	case DellLaptop:
		return dellLaptop{}
	default:
		return nil
	}
}

type desktopFactory struct{}

func (d desktopFactory) NewComputer(c computer) Computer {
	switch c {
	case HPDesktop:
		return hpDesktop{}
	case DellDesktop:
		return dellDesktop{}
	default:
		return nil
	}
}

type factory int

const (
	DesktopFactory factory = iota
	LaptopFactory
)

// The abstract factory
func NewComputerFactory(f factory) (ComputerFactory, error) {
	switch f {
	case DesktopFactory:
		return desktopFactory{}, nil
	case LaptopFactory:
		return laptopFactory{}, nil
	}
	return nil, errors.New("not implemented yet")
}
