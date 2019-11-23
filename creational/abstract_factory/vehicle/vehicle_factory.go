package vehicle

import "errors"

type AbstractVehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (AbstractVehicleFactory, error) {
	switch f {
	default:
		return nil, errors.New("vehicle factory type is not yet available")
	}
}
