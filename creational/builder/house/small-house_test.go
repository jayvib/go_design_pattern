package house

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHouseBuilder(t *testing.T) {
	smallHouseStructure := &SmallHouse{
		house: &House{},
	}
	builder := NewBuildComplex(smallHouseStructure)
	builder.Build()

	houseResult := smallHouseStructure.GetHouse()
	assert.Equal(t, 1, houseResult.door)
	assert.Equal(t, 4, houseResult.post)
	assert.Equal(t, 1, houseResult.window)
	assert.Equal(t, 1, houseResult.roof)
}
