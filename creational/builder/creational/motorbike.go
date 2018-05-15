package creational

type MotorbikeBuilder struct{
	vehicle VehicleProduct
}

func (mb *MotorbikeBuilder) SetWheels() BuildProcess {
	mb.vehicle.Wheels = 2
	return mb
}

func (mb *MotorbikeBuilder) SetSeats() BuildProcess {
	mb.vehicle.Seats = 2
	return mb
}

func (mb *MotorbikeBuilder) SetStructure() BuildProcess {
	mb.vehicle.Structure = "Motorbike"
	return mb
}

func (mb *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return mb.vehicle
}
