package person

import "fmt"

// The Visitor will be the constructor of the Person's Profile

type Person struct {
	Firstname, Lastname string
	Age                 int
	Gender              string
	Profile             string
}

func (p *Person) Accept(v Visitor) {
	v.Visit(p)
}

func (p *Person) String() string {
	return p.Profile
}

type Visitor interface {
	Visit(*Person)
}

type visitor struct {}

func (visitor) Visit(p *Person) {
	profileFmt := fmt.Sprintf("My name is %s, %s and %d years old.",
		p.Firstname+ " " + p.Lastname, p.Gender, p.Age,
	)

	p.Profile = profileFmt
}

type Visitable interface {
	Accept(Visitor)
	String() string
}

func NewVisitor() Visitor {
	return visitor{}
}