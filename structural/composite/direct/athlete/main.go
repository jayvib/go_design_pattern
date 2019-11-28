package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func swim() {
	fmt.Println("Swimming")
}

type Animal struct{}

func (Animal) Eat() {
	fmt.Println("Eating")
}

type Fish struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl string

func (s SwimmerImpl) Swim() {
	fmt.Printf("%s is swimming\n", s)
}

type TrainerImpl string

func (t TrainerImpl) Train() {
	fmt.Printf("%s is training\n", t)
}

// Embedding composition
type EmbeddingImpl struct {
	Swimmer
	Trainer
}

func main() {
	compositeSwimmerA := CompositeSwimmerA{
		MySwim: swim,
	}

	compositeSwimmerA.MyAthlete.Train()
	compositeSwimmerA.MySwim()

	shark := Fish{
		Swim: swim,
	}

	shark.Eat()
	shark.Swim()

	jayson := EmbeddingImpl{
		Swimmer: SwimmerImpl("jayson"),
		Trainer: TrainerImpl("jayson"),
	}

	jayson.Train()
	jayson.Swim()
}
