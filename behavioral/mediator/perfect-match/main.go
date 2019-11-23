package main

import "fmt"

func amazingDisplayDecorator(d resultDisplayer) resultDisplayer {
	fmt.Println("its amazing you know!!")
	return d
}

func funDisplayDecorator(d resultDisplayer) resultDisplayer {
	fmt.Println("its fun!!")
	return d
}

type resultDisplayer interface {
	display()
}

type Jayson struct{}

func (Jayson) name() string {
	return "Jayson"
}

type Althea struct{}

func (Althea) name() string {
	return "Althea"
}

type person interface {
	name() string
}

func sayWhat(p person, i interface{}) resultDisplayer {
	switch p.(type) {
	case Jayson:
		switch i.(type) {
		case one:
			o := i.(one)
			o.p = p
			return o
		case two:
			t := i.(two)
			t.p = p
			return t
		case int:
			return intDisplayerFunc(func() {
				fmt.Printf("%d %s\n", i, p.name())
			})
		}
	case Althea:
		switch i.(type) {
		case one:
			o := i.(one)
			o.p = p
			return o
		case two:
			t := i.(two)
			t.p = p
			return t
		case int:
			return intDisplayerFunc(func() {
				fmt.Printf("%d %s\n", i, p.name())
			})
		}
	}
	return nil
}

type one struct {
	p person
}

func (o one) display() {
	fmt.Printf("one '%s'\n", o.p.name())
}

type two struct {
	p person
}

func (t two) display() {
	fmt.Printf("two '%s'\n", t.p.name())
}

type intDisplayerFunc func()

func (fn intDisplayerFunc) display() {
	fn()
}

func main() {
	jayson := Jayson{}
	althea := Althea{}

	sayWhat(jayson, one{}).display()
	sayWhat(althea, one{}).display()
	sayWhat(jayson, 2).display()

	// with decorator
	funDisplayDecorator(
		amazingDisplayDecorator(
			sayWhat(jayson, one{}))).display()

	sayWhat(jayson, two{}).display()
	sayWhat(althea, two{}).display()
	sayWhat(althea, 3).display()
}
