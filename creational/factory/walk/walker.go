package walk

import "errors"

const (
	animal = 1
	person = 2
)

type Walker interface {
	Walk()
}

func GetWalker(o int) (Walker, error) {
	switch o {
	default:
		return nil, errors.New("not implemented yet")
	}
}
