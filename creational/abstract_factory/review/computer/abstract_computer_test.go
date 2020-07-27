package computer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewComputerFactory(t *testing.T) {
	t.Run("Desktop Factory", func(t *testing.T) {
		f, err := NewComputerFactory(DesktopFactory)
		if err != nil {
			t.Fatal(err)
		}

		// Test for HP Desktop
		t.Run("HP Desktop Computer", func(t *testing.T) {
			comp := f.CreateComputer(HPDesktop)

			// check its method return
			assert.Equal(t, "hp desktop display", comp.Display())
			assert.Equal(t, "hp desktop mouse", comp.Mouse())
			assert.Equal(t, "hp desktop keyboard", comp.Keyboard())
			assert.Equal(t, "hp quadcore cpu", comp.CPU())
		})

		t.Run("Dell Desktop Computer", func(t *testing.T) {
			comp := f.CreateComputer(DellDesktop)
			// check its method return
			assert.Equal(t, "dell desktop display", comp.Display())
			assert.Equal(t, "dell desktop mouse", comp.Mouse())
			assert.Equal(t, "dell desktop keyboard", comp.Keyboard())
			assert.Equal(t, "dell quadcore cpu", comp.CPU())
		})

	})

	t.Run("Laptop Factory", func(t *testing.T) {
		f, err := NewComputerFactory(LaptopFactory)
		if err != nil {
			t.Fatal(err)
		}

		// Test for HP Laptop
		t.Run("HP Laptop Computer", func(t *testing.T) {
			comp := f.CreateComputer(HPLaptop)
			// check its method return
			assert.Equal(t, "hp laptop display", comp.Display())
			assert.Equal(t, "hp laptop mouse", comp.Mouse())
			assert.Equal(t, "hp laptop keyboard", comp.Keyboard())
			assert.Equal(t, "hp quadcore cpu", comp.CPU())
		})

		t.Run("Dell Laptop Computer", func(t *testing.T) {
			comp := f.CreateComputer(DellLaptop)
			// check its method return
			assert.Equal(t, "dell laptop display", comp.Display())
			assert.Equal(t, "dell laptop mouse", comp.Mouse())
			assert.Equal(t, "dell laptop keyboard", comp.Keyboard())
			assert.Equal(t, "dell quadcore cpu", comp.CPU())
		})
	})
}
