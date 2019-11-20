package shirt

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetShirtCloner(t *testing.T) {
	t.Run("when getting the shirt clone it should not be nil", func(t *testing.T){
		shirtCloner := GetShirtCloner()
		require.NotNil(t, shirtCloner)
	})

	t.Run("when cloning two instances there memory address should not the same", func(t *testing.T){
		shirtCloner := GetShirtCloner()
		require.NotNil(t, shirtCloner)

		shirt := shirtCloner.GetClone(White)
		require.NotNil(t, shirt)

		shirt2 := shirtCloner.GetClone(White)
		assert.False(t, shirt==shirt2)

		item := shirt.(*Shirt)
		item.SKU = "W-SHIRT1"

		item2 := shirt2.(*Shirt)
		item2.SKU = "W-SHIRT2"

		assert.NotEqual(t, item.SKU, item2.SKU)
	})
}
