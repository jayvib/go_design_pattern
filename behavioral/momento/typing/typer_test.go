package typing

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestOriginator(t *testing.T) {
	orig := originator{}

	t.Run("Get memento", func(t *testing.T){
		s := state{ "a" }
		orig.memento.state = s

		mem := orig.getmemento()
		assert.Equal(t, mem.state.s, s.s, "They are not equal")
	})

	t.Run("Extract and Store State", func(t *testing.T){
		s := state{ "j" }
		mem := memento{ s }
		orig.extractAndStoreState(mem)
		assert.Equal(t, orig.memento.state.s, mem.state.s, "They are not equal")
	})
}