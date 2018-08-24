package main

import "fmt"

type Worker interface {
	Work(name string)
}

type ChainWork []Worker

func (c *ChainWork) AddWork(workers ...Worker) {
	for _, w := range workers {
		*c = append(*c, w)
	}
}

func (c *ChainWork) Work(name string) {
	for _, cw := range *c {
		cw.Work(name)
	}
}

type LaundryWork struct {}

func (LaundryWork) Work(name string) {
	fmt.Printf("%s is doing laundry\n", name)
}

type CookingWork struct {}

func (CookingWork) Work(name string) {
	fmt.Printf("%s is cooking a dishes\n", name)
}

type CodingWork struct {}

func (CodingWork) Work(name string) {
	fmt.Printf("%s is coding Go\n", name)
}

func main() {
	var laundry LaundryWork
	var cooking CookingWork
	var coding CodingWork

	cw := new(ChainWork)
	cw.AddWork(laundry, cooking, coding)
	cw.Work("Jayson")
}
