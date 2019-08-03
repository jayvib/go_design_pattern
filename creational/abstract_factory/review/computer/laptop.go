package computer

// This will implements the Laptop Computer
type hpLaptop struct{}

func (hpLaptop) Display() string {
	return "hp laptop display"
}

func (hpLaptop) Mouse() string {
	return "hp laptop mouse"
}

func (hpLaptop) Keyboard() string {
	return "hp laptop keyboard"
}

func (hpLaptop) CPU() string {
	return "hp quadcore cpu"
}

type dellLaptop struct{}

func (dellLaptop) Display() string {
	return "dell laptop display"
}

func (dellLaptop) Mouse() string {
	return "dell laptop mouse"
}

func (dellLaptop) Keyboard() string {
	return "dell laptop keyboard"
}

func (dellLaptop) CPU() string {
	return "dell quadcore cpu"
}
