package main

import "fmt"

// An Object that will provide a value then
// set the processed value
type Value interface {
	GetNumber() int
	SetNumber(n int)
}

// Visitable is a generic type that will
// accept v Value interface.
//
// All the implementations of the logic
// will be the responsibility of this
// interface in order to process the value,
type Visitable interface {
	Accept(v Value)
}

type value struct {
	given  int
	result int
}

func (v *value) GetNumber() int {
	return v.given
}

func (v *value) SetNumber(n int) {
	v.result = n
}

type multiplierBy struct {
	multiplier int
}

func (m *multiplierBy) Accept(v Value) {
	v.SetNumber(m.product(v.GetNumber()))
}

func (m *multiplierBy) product(n int) int {
	return m.multiplier * n
}

func (v *value) String() string {
	return fmt.Sprintf("The input %d and the output is %d", v.given, v.result)
}

func main() {
	var visitable Visitable
	var val Value

	visitable = &multiplierBy{multiplier: 10}
	val = &value{given: 5}

	visitable.Accept(val)

	stringer, ok := val.(fmt.Stringer)
	if ok {
		fmt.Println(stringer)
	}
}
