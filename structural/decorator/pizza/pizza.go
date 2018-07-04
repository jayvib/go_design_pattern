package pizza

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
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
	result := ""
	err := (error)(nil)
	if m.Ingredient != nil {
		result, err = m.Ingredient.AddIngredient()
		if err != nil {
			return "", err
		}
	}
	result += "meat"
	return result, nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An Ingredient Add is needed in the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}

func AddMeatIngredientDecorator(i IngredientAdd) IngredientAdd {
	return &Meat{
		Ingredient: i,
	}
}

func AddOnionIngredientDecorator(i IngredientAdd) IngredientAdd {
	return &Onion{
		Ingredient: i,
	}
}

type Marshmallow struct {
	Ingredient IngredientAdd
}

func (m *Marshmallow) AddIngredient() (string, error) {
	result := ""
	err := (error)(nil)
	if m.Ingredient != nil {
		result, err = m.Ingredient.AddIngredient()
		if err != nil {
			return "", err
		}
	}
	result += "marshmallow"
	return result, nil
}

func AddMarshmallowIngredientDecorator(i IngredientAdd) IngredientAdd {
	return &Marshmallow{
		Ingredient: i,
	}
}