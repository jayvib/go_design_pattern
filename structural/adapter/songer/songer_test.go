package songer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestSonger(t *testing.T) {
	song := "My Way"
	expected := fmt.Sprintf("Singing %s", song)
	oldSonger := NewVintageSonger()
	singing := oldSonger.Sing(song)
	assert.Equal(t, expected, singing, "Expected result wasn't matched with the actual result")

	funkySonger := NewModernSonger(oldSonger, song)
	singing = funkySonger.SingingStoredSong()
	expected += "in upbeat"
	assert.Equal(t, "Singing My Way in upbeat", singing, "Expected result wasn't matched with the actual result")

	newsong := "Uptown Funk"
	funkySonger.PrepareNewSong(newsong)
	expected = fmt.Sprintf("Singing %s in upbeat", newsong)
	singing = funkySonger.SingingStoredSong()
	assert.Equal(t, singing, expected, "Expected result wasn't matched with the actual result")
}
