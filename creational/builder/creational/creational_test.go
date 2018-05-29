package creational

import "testing"

func TestCarBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := new(CarBuilder)
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Error("Car wheels must be 4 but got", car.Wheels)
	}

	if car.Seats != 5 {
		t.Error("Car seats must be 5 but got", car.Seats)
	}

	if car.Structure != "Car" {
		t.Error("Structure must be Car but got", car.Structure)
	}
}

func TestMotorbikeBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	bikeBuilder := new(MotorbikeBuilder)
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Error("Motorbike must have 2 wheels but got", bike.Wheels)
	}

	if bike.Seats != 2 {
		t.Error("Motorbike must have 2 seats but got", bike.Seats)
	}

	if bike.Structure != "Motorbike" {
		t.Error("Structure must be motorbike but got", bike.Structure)
	}
}

func TestPedicabBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	pedicabBuilder := new(PedicabBuilder)
	manufacturingComplex.SetBuilder(pedicabBuilder)
	manufacturingComplex.Construct()
	pedicab := pedicabBuilder.GetVehicle()

	if pedicab.Wheels != 3 {
		t.Error("Pedicab must have 3 wheels but got", pedicab.Wheels)
	}

	if pedicab.Seats != 3 {
		t.Error("Pedicab must have 3 seats but got", pedicab.Seats)
	}

	if pedicab.Structure != "Pedicab" {
		t.Error("Structure must be Pedicab but got", pedicab.Structure)
	}

}
