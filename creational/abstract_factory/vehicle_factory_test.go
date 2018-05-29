package abstract_factory

import "testing"

func TestSportMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	sportMotorbike, err := motorbikeFactory.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	if sportMotorbike.NumWheels() != 2 {
		t.Error("sport motorbike wheels must be 2 but has", sportMotorbike.NumWheels())
	}

	if sportMotorbike.NumSeats() != 1 {
		t.Error("sport motorbike seat must be 1 but has", sportMotorbike.NumSeats())
	}
}

func TestCruiseMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	cruiseMotorbike, err := motorbikeFactory.NewVehicle(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	if cruiseMotorbike.NumWheels() != 2 {
		t.Error("cruise motorbike wheels must be 2 but has", cruiseMotorbike.NumWheels())
	}

	if cruiseMotorbike.NumSeats() != 2 {
		t.Error("sport motorbike seat must be 1 but has", cruiseMotorbike.NumSeats())
	}
}

func TestLuxuryCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	luxuryCar, err := carFactory.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	if luxuryCar.NumWheels() != 4 {
		t.Error("luxury car must have 4 wheels but has", luxuryCar.NumWheels())
	}

	if luxuryCar.NumSeats() != 5 {
		t.Error("luxury car must have 4 seats but has", luxuryCar.NumSeats())
	}
}
