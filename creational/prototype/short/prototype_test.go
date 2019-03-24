package short

import (
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestShortCloner(t *testing.T) {
	cloner := ShortCloner{}


	t.Run("Address Testing", func(t *testing.T){
		item1, err := cloner.GetClone(Black)
		checkErr(t, err)
		if item1 == BlackShortPrototype {
			t.Error("Pointer address must not be the same for black short")
		}

		item2, err := cloner.GetClone(Red)
		checkErr(t, err)
		if item2 == RedShortPrototype {
			t.Error("Pointer address must not be the same for red short")
		}

		item3, err := cloner.GetClone(Blue)
		checkErr(t, err)
		if item3 == BlueShortPrototype {
			t.Error("Pointer address must not be the same for blue short")
		}
	})

	t.Run("Data Modification", func(t *testing.T){
		item1, err := cloner.GetClone(Blue)
		checkErr(t, err)

		blueShort, ok := item1.(*Short)
		if !ok {
			t.Fatal("item1 must be a Short object when casted")
		}

		blueShort.SKU = "blahblah"

		item2, err := cloner.GetClone(Blue)
		checkErr(t, err)

		blueShort2, ok := item2.(*Short)
		if !ok {
			t.Fatal("item2 must be a Short object when casted")
		}

		assert.NotEqual(t, blueShort.SKU, blueShort2.SKU,
			"Modifying the SKU in blueshort 1 must not reflect in cloned blueshort 2.")
	})
}
