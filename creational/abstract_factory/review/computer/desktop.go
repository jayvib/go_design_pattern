package computer

// This will implements the Laptop Computer
type hpDesktop struct{}

func (hpDesktop) Display() string {
	return "hp desktop display"
}

func (hpDesktop) Mouse() string {
	return "hp desktop mouse"
}

func (hpDesktop) Keyboard() string {
	return "hp desktop keyboard"
}

func (hpDesktop) CPU() string {
	return "hp quadcore cpu"
}

type dellDesktop struct{}

func (dellDesktop) Display() string {
	return "dell desktop display"
}

func (dellDesktop) Mouse() string {
	return "dell desktop mouse"
}

func (dellDesktop) Keyboard() string {
	return "dell desktop keyboard"
}

func (dellDesktop) CPU() string {
	return "dell quadcore cpu"
}
