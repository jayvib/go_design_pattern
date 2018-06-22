package songer

import "fmt"

type VintageSonger interface {
	Sing(title string) string
}

func NewVintageSonger() VintageSonger {
	return vintageSongerImpl{}
}

type vintageSongerImpl struct {}

func (vintageSongerImpl) Sing(title string) string {
	return fmt.Sprintf("Singing %s", title)
}

type ModernSonger interface {
	SingingStoredSong() string
	PrepareNewSong(song string)
}

func NewModernSonger(songer VintageSonger, song string) ModernSonger {
	return &modernSongerImpl{
		oldSonger: songer,
		song:      song,
	}
}

type modernSongerImpl struct {
	oldSonger VintageSonger
	song      string
}

func (m *modernSongerImpl) SingingStoredSong() string {
	return fmt.Sprintf("%s in upbeat", m.oldSonger.Sing(m.song))
}

func (m *modernSongerImpl) PrepareNewSong(song string) {
	m.song = song
}

