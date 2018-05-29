package main

import (
	"fmt"
)

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	swim := func() { fmt.Println("Swimming!") }

	swimmer := CompositeSwimmerA{
		MySwim: swim,
	}

	swimmer.MySwim()
	swimmer.MyAthlete.Train()

	compoSwimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	compoSwimmer.Train()
	compoSwimmer.Swim()
}
