package pizza

import (
	"errors"
	"fmt"
)

var (
	NoSetObject = errors.New("pizza: an Ingredient field don't have set object.")
)


type IngredientAdd interface {
	AddIngredient() (string, error)
}

func WithIngredient(ingredient ...IngredientAdd) (IngredientAdd, error) {
	switch len(ingredient) {
	case 0:
		return nil, errors.New("pizza: main pizza must be provided")
	case 1:
		return nil, errors.New("pizza: must provide an added ingredient to the pizza")
	}

	return nil, errors.New("not implemented yet")
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", NoSetObject
	}

	p, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", p, "meat"), nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "onion"), nil
}