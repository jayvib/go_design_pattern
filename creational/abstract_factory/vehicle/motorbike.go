package vehicle

type Motorbike interface {
	GetMotorbikeType() int
}

type SportMotorbike struct{}

func (s *SportMotorbike) NumSeats() int {
	return 0
}

func (s *SportMotorbike) NumWheels() int {
	return 0
}

func (s *SportMotorbike) GetMotorbikeType() int {
	return 0
}

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) NumSeats() int {
	return 0
}

func (c *CruiseMotorbike) NumWheels() int {
	return 0
}

func (c *CruiseMotorbike) GetMotorbikeType() int {
	return 0
}
