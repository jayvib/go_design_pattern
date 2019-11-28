package counter

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Requirement:
// 1. When getting the instance it should be
//    the same pointer without re-initializing the
//    default counter.
func TestGetInstance(t *testing.T) {
	t.Run("When calling GetDefaultCounter the same instance should be return", func(t *testing.T) {
		i1 := GetDefaultCounter() // first initialization
		require.NotNil(t, i1)

		i2 := GetDefaultCounter()

		// They should be the same
		assert.True(t, i1 == i2)
	})

	t.Run("When calling GetDefaultCounter multiple times it should not re-initialized", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			GetDefaultCounter()
		}
		assert.Equal(t, 1, defaultCounterInitializedCalls)
	})
}

func TestCounter_AddOne(t *testing.T) {
	c1 := GetDefaultCounter()
	c1.AddOne()
	c1.AddOne()
	assert.Equal(t, 2, defaultCounter.count)
}
