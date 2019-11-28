package vehicle

type SportsCar struct {
	vehicle Vehicle
}

func (c *SportsCar) SetDoors() VehicleBuildProcess {
	c.vehicle.NumDoors = 4
	return c
}

func (c *SportsCar) SetWheels() VehicleBuildProcess {
	c.vehicle.NumWheels = 2
	return c
}

func (c *SportsCar) GetVehicle() Vehicle {
	return c.vehicle
}

type FamilyCar struct {
	vehicle Vehicle
}

func (f *FamilyCar) SetDoors() VehicleBuildProcess {
	f.vehicle.NumDoors = 4
	return f
}

func (f *FamilyCar) SetWheels() VehicleBuildProcess {
	f.vehicle.NumWheels = 4
	return f
}

func (f *FamilyCar) GetVehicle() Vehicle {
	return f.vehicle
}
