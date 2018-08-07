package shoes

type ItemInfoGetter interface {
	GetPrice() int
	GetColor() string
}

var whiteShoePrototype = &Shoe{
	Price: 15,
	Color: white,
}

var blackShoePrototype = &Shoe{
	Price: 16,
	Color: black,
}

var blueShoePrototype = &Shoe {
	Price: 17,
	Color: blue,
}

type ShirtColor byte

func (c ShirtColor) ParseColor() string {
	switch c {
	case white:
		return "white"
	case black:
		return "black"
	case blue:
		return "blue"
	default:
		return ""
	}
}

type Shoe struct {
	Price int
	Color ShirtColor
}

func (s *Shoe) GetPrice() int {
	return s.Price
}

func (s *Shoe) GetColor() string {
	return s.Color.ParseColor()
}