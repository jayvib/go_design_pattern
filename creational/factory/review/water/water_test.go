package water

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetWater(t *testing.T) {

	t.Run("Faucet", func(t *testing.T) {
		wg := GetWater(Faucet)
		w := wg.Get()
		assert.Contains(t, w, "faucet", "expecting to contain faucet but nothing found. got: %s", w)
	})

	t.Run("Well", func(t *testing.T) {
		wg := GetWater(Well)
		w := wg.Get()
		assert.Contains(t, w, "well", "expecting to contain well but found nothing. got: %s", w)
	})
}
