package computer

import "testing"

func TestDirector(t *testing.T) {
	laptopBuilder := NewLaptopBuilder()

	director := NewDirector(laptopBuilder)
	director.Build()

	computerProduct := laptopBuilder.GetComputer()
	if computerProduct.monitor != "Dell" {
		t.Errorf("expecting Dell monitor but got %s", computerProduct.monitor)
	}

	if computerProduct.keyboard != "HP" {
		t.Errorf("expecting HP keyboard but got %s", computerProduct.keyboard)
	}

	if computerProduct.mouse != "Lenovo" {
		t.Errorf("expecting Lenovo mouse but got %s", computerProduct.mouse)
	}
}
