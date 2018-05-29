package creational

type PedicabBuilder struct {
	VehicleProduct
}

func (p *PedicabBuilder) SetWheels() BuildProcess {
	p.VehicleProduct.Wheels = 3
	return p
}

func (p *PedicabBuilder) SetSeats() BuildProcess {
	p.VehicleProduct.Seats = 3
	return p
}

func (p *PedicabBuilder) SetStructure() BuildProcess {
	p.VehicleProduct.Structure = "Pedicab"
	return p
}

func (p *PedicabBuilder) GetVehicle() VehicleProduct {
	return p.VehicleProduct
}
