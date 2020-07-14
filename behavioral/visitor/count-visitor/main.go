package main

import "fmt"

type Accepter interface {
	Accept(v Visitor)
}

type String string

func (s String) Accept(v Visitor) {
	v.VisitString(s)
}

type Int int

func (s Int) Accept(v Visitor) {
	v.VisitInt(s)
}

type Visitor interface {
	VisitString(s String)
	VisitInt(i Int)
}

// Visitor Implementation/Concrete
type visitor struct {
	Strings []string
	Ints    []int
}

func (v *visitor) VisitString(s String) {
	v.Strings = append(v.Strings, string(s))
}
func (v *visitor) VisitInt(i Int) {
	v.Ints = append(v.Ints, int(i))
}

func (v *visitor) String() string {
	return fmt.Sprintf("String Len: %d, Int Len: %d\n", len(v.Strings), len(v.Ints))
}

type App struct {
	acceptors []Accepter
}

func (a *App) Run(v Visitor) {
	for _, a := range a.acceptors {
		a.Accept(v)
	}
}

func main() {
	v := new(visitor)

	app := &App{
		acceptors: []Accepter{
			String("Luffy"),
			String("Sanji"),
			Int(1),
			Int(2),
			Int(3),
		},
	}

	app.Run(v)
	fmt.Println(v)
}
