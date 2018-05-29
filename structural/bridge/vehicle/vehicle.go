// This is an attempt to implement a bridge design pattern
package vehicle

import "fmt"

type PartGetter interface {
	GetWheels() int
	GetSeats() int
	GetDoors() int
	WhatIveGot() string
}

type JunkVehicle struct {
	Wheel, Seat, Door int
}

func (jv *JunkVehicle) GetWheels() int {
	return jv.Wheel
}

func (jv *JunkVehicle) GetSeats() int {
	return jv.Seat
}

func (jv *JunkVehicle) GetDoors() int {
	return jv.Door
}

func (jv *JunkVehicle) WhatIveGot() string {
	return fmt.Sprintf("got %d wheels %d seats %d doors",
		jv.Wheel, jv.Seat, jv.Door)
}

type PartAddedJunkVehicle struct {
	jv *JunkVehicle
}

func (self *PartAddedJunkVehicle) addone(n int) int {
	return n + 1
}

func (self *PartAddedJunkVehicle) GetWheels() int {
	return self.addone(self.jv.GetWheels())
}

func (self *PartAddedJunkVehicle) GetSeats() int {
	return self.addone(self.jv.GetSeats())
}

func (self *PartAddedJunkVehicle) GetDoors() int {
	return self.addone(self.jv.GetDoors())
}

func (self *PartAddedJunkVehicle) WhatIveGot() string {
	gotFrom := self.jv.WhatIveGot()
	WithAddedGot := fmt.Sprintf(gotFrom + " with additional 1 each")
	return WithAddedGot
}
