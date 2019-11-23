package shoes

import "testing"

func TestShoesPrototype(t *testing.T) {
	var cloner Cloner
	shoeCloner := ShoeCloner{}
	cloner = shoeCloner

	whiteShoe, err := cloner.Clone(white)
	if err != nil {
		t.Fatal(err.Error())
	}

	if whiteShoe == nil {
		t.Fatal("whiteShoe must not be empty")
	}

	if whiteShoe.GetPrice() != 15 {
		t.Fatal(err.Error())
	}

	blackShoe, err := cloner.Clone(black)
	if err != nil {
		t.Fatal(err.Error())
	}

	if blackShoe.GetPrice() != 16 {
		t.Fatal(err.Error())
	}

	if blackShoe == nil {
		t.Fatal("blackShoe must not be empty")
	}

	blueShoe, err := cloner.Clone(blue)
	if err != nil {
		t.Fatal(err.Error())
	}

	if blueShoe == nil {
		t.Fatal(err.Error())
	}

	if blueShoe.GetPrice() != 17 {
		t.Errorf("Expecting price %d but got %d\n", blackShoe.GetPrice())
	}
}
