package animal

type Mover interface {
	Move()
}

type Animal struct {
	Mover
}

type Person struct {
	Animal // in bridge design pattern the base object should inherit the implementation object
}

type Bird struct {
	Animal
}

type Fish struct {
	Animal
}

// Bridge Design Pattern
// Implementation Interface
// Implementation Object
// Base Client Object
