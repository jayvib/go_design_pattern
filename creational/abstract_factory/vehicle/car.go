package vehicle

type Car interface {
	NumDoors() int
}

type LuxuryCar struct{}

func (l *LuxuryCar) NumSeats() int {
	return 0
}

func (l *LuxuryCar) NumWheels() int {
	return 0
}

func (l *LuxuryCar) NumDoors() int {
	return 0
}

type FamilyCar struct{}

func (f *FamilyCar) NumSeats() int {
	return 0
}

func (f *FamilyCar) NumWheels() int {
	return 0
}

func (f *FamilyCar) NumDoors() int {
	return 0
}
