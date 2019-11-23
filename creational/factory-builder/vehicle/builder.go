package vehicle

// Builder Interface
type VehicleBuildProcess interface {
	SetDoors() VehicleBuildProcess
	SetWheels() VehicleBuildProcess
	GetVehicle() Vehicle
}

func NewManufacturingVehicleDirector(v car) *ManufacturingVehicleDirector {
	return &ManufacturingVehicleDirector{
		process: NewVehicleBuilder(v),
	}
}

type ManufacturingVehicleDirector struct {
	process VehicleBuildProcess
}

func (m *ManufacturingVehicleDirector) SetVehicleProcess(v VehicleBuildProcess) {
	m.process = v
}

func (m *ManufacturingVehicleDirector) Build() Vehicle {
	return m.process.SetDoors().SetWheels().GetVehicle()
}
