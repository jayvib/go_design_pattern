package person_test

import (
	"fmt"
	"go_design_pattern/behavioral/visitor/person"
	"testing"
)

type TestData struct {
	visitable person.Visitable
	expected  string
}

func TestPerson_Accept(t *testing.T) {
	pfmt := "My name is %s, %s and %d years old."
	visitables := make([]TestData, 0)

	p1 := &person.Person{
		Firstname: "Jayson",
		Lastname:  "Vibandor",
		Age:       27,
		Gender:    "Male",
	}
	tdata1 := TestData{
		visitable: p1,
		expected:  fmt.Sprintf(pfmt, p1.Firstname+" "+p1.Lastname, p1.Gender, p1.Age),
	}

	visitables = append(visitables, tdata1)

	p2 := &person.Person{
		Firstname: "Althea",
		Lastname:  "Bongalon",
		Age:       25,
		Gender:    "Female",
	}

	tdata2 := TestData{
		visitable: p2,
		expected:  fmt.Sprintf(pfmt, p2.Firstname+" "+p2.Lastname, p2.Gender, p2.Age),
	}

	visitables = append(visitables, tdata2)

	visitor := person.NewVisitor()

	for _, p := range visitables {
		p.visitable.Accept(visitor)
		if p.expected != p.visitable.String() {
			t.Errorf("expected profile wasn't match")
		}
	}
}
