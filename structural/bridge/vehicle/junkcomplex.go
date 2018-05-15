package vehicle

import "fmt"

type JunkComplex struct {
	parts PartGetter
}

func (jc *JunkComplex) WheelsIGot() string {
	return fmt.Sprintf("I got %d wheels\n", jc.parts.GetWheels())
}

func (jc *JunkComplex) DoorsIGot() string {
	return fmt.Sprintf("I got %d doors\n", jc.parts.GetDoors())
}

func (jc *JunkComplex) SeatsIGot() string {
	return fmt.Sprintf("I got %d seats\n", jc.parts.GetSeats())
}

func (jc *JunkComplex) IGot() string {
	return jc.parts.WhatIveGot()
}