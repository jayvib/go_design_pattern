package shoes

const (
	Modern ShoeStyle = iota
	Vintage
	Colorful
)

type SneakerShoesFactory struct {}
func (s SneakerShoesFactory) CreateShoes(style ShoeStyle) Shoes {
	switch style {
	case Modern:
	case Vintage:
	case Colorful:
	}
	return nil
}
func (s SneakerShoesFactory) AvailableStyles() []ShoeStyle {
	return nil
}

type modernSneaker struct{}
func (m modernSneaker) Lace() string        { return "" }
func (m modernSneaker) Sole() string        { return "" }
func (m modernSneaker) Upper() string       { return "" }
func (m modernSneaker) Vamp() string        { return "" }
func (m modernSneaker) HeelCounter() string { return "" }

type vintageSneaker struct {}
func (v vintageSneaker) Lace() string        { return "" }
func (v vintageSneaker) Sole() string        { return "" }
func (v vintageSneaker) Upper() string       { return "" }
func (v vintageSneaker) Vamp() string        { return "" }
func (v vintageSneaker) HeelCounter() string { return "" }

type colorfulSneaker struct {}
func (c colorfulSneaker) Lace() string        { return "" }
func (c colorfulSneaker) Sole() string        { return "" }
func (c colorfulSneaker) Upper() string       { return "" }
func (c colorfulSneaker) Vamp() string        { return "" }
func (c colorfulSneaker) HeelCounter() string { return "" }

