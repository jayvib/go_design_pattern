package vehicle

type car int

const (
	SPORTS car = 0
	FAMILY car = 1
)

func NewVehicleBuilder(v car) VehicleBuildProcess {
	switch v {
	case SPORTS:
		return new(SportsCar)
	case FAMILY:
		return new(FamilyCar)
	}
	return nil
}
