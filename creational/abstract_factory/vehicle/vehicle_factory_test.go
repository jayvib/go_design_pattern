package vehicle

import (
	"fmt"
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	sportMotorbike, err := motorbikeF.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Sports motorbike has %s wheels and %s seats\n", sportMotorbike.NumWheels(), sportMotorbike.NumSeats())
}
