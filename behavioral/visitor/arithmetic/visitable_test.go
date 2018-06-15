package arithmetic

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestMultiplier_Multiply(t *testing.T) {
	nc := &numberContainer{ number: 20 }
	multiplierVisitor := &multiplier{ multNum: 2 }

	nc.Accept(multiplierVisitor)

	assert.Equal(t, nc.number, 40, "expected value doesn't match")
}

func Test