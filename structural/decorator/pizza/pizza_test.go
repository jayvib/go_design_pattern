package pizza

import (
	"strings"
	"testing"
	"fmt"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient of the pizza decorator "+
			"it must return the expected text. \n\tExpected: %s\n\tActual: %s",
			expectedText, pizzaResult)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(onionResult, "onion") {
		t.Error("When calling AddIngredient the result must now contain an 'onion' but has ", onionResult)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.AddIngredient()
	if err != nil {
		t.Error("When calling the meat addIngredient that is empty must return an error.")
	}

	if !strings.Contains(meatResult, "meat") {
		t.Errorf("When calling the add ingredient of the meat decorator, it must"+
			"return the expected text. \n\tExpected: %s\n\tActual: %s",
			"meat", meatResult)
	}

	t.Run("Onion", func(t *testing.T){
		deco := AddOnionIngredientDecorator(meat)
		res, err := deco.AddIngredient()
		if err != nil {
			t.Fatal(err.Error())
		}
		if !strings.Contains(res, "onion") {
			t.Error("item must already have an onion")
		}
		fmt.Println(res)

		t.Run("Marshmallow", func(t *testing.T){
			deco = AddMarshmallowIngredientDecorator(deco)
			res, err = deco.AddIngredient()
			if err != nil {
				t.Fatal(err.Error())
			}
			if !strings.Contains(res, "marshmallow") {
				t.Error("item must already have a marshmallow")
			}
			fmt.Println(res)
		})

	})

}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err == nil {
		t.Error(err)
	}

	expectedText := "Pizza with the following ingredients: meat, onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and meat the returned string must"+
			"be expected.\n\tExpected: %s\n\tActual: %s", expectedText, pizzaResult)
	}
}
