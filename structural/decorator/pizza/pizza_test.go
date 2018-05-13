package pizza

import (
	"testing"
	"strings"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient of the pizza decorator " +
			"it must return the expected text. \n\tExpected: %s\n\tActual: %s",
				expectedText, pizzaResult)
	}

	t.Run("TestOnion_AddIngredient", func(t *testing.T){
		onion := &Onion{Ingredient: pizza}
		onionResult, err := onion.AddIngredient()
		if err != nil {
			t.Error(err)
		}

		if !strings.Contains(onionResult, "onion") {
			t.Error("When calling AddIngredient the result must now contain an 'onion' but has ", onionResult)
		}
	})

	t.Run("TestMeat_AddIngredient", func(t *testing.T) {
		meat := &Meat{Ingredient: pizza}
		meatResult, err := meat.AddIngredient()
		if err != nil {
			t.Error("When calling the meat addIngredient that is empty must return an error.")
		}

		if !strings.Contains(meatResult, "meat") {
			t.Errorf("When calling the add ingredient of the meat decorator, it must" +
				"return the expected text. \n\tExpected: %s\n\tActual: %s",
				"meat", meatResult)
		}
	})
}


func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	expectedText := "Pizza with the following ingredients: meat onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and meat the returned string must"+
			"be expected.\n\tExpected: %s\n\tActual: %s", expectedText, pizzaResult)
	}
}

func TestWithIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}

 	specialPizza, err := WithIngredient(pizza, new(Onion))
 	if err != nil {
 		t.Fatal(err)
	}

 	pizzaResult, err := specialPizza.AddIngredient()
 	if err != nil {
 		t.Error(err)
	}

	if !strings.Contains(pizzaResult, "onion") {
		t.Errorf("pizza with added onion ingredient to it must have an onion but has: %s\n", pizzaResult)
	}

}