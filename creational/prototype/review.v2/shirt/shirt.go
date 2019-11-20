package shirt

type ShirtCloner interface {
	GetClone(s int) ItemInfoGetter
}

type ShirtClonerFunc func(s int) ItemInfoGetter

func (fn ShirtClonerFunc) GetClone(s int) ItemInfoGetter {
	return fn(s)
}

type ItemInfoGetter interface {
	GetInfo() string
}

func GetShirtCloner() ShirtCloner {
	return ShirtClonerFunc(GetShirtClone)
}

const (
	White = iota
	Blue
	Black
)

var whiteShirtPrototype = &Shirt{
	Color: White,
	Price: 15,
	SKU: "empty",
}

var blackShirtPrototype = &Shirt{
	Color: Black,
	Price: 16,
	SKU: "empty",
}

var blueShirtPrototype = &Shirt{
	Color: Blue,
	Price: 17,
	SKU: "empty",
}

func GetShirtClone(color int) ItemInfoGetter {
	switch color {
	case White:
		whiteShirtCopy := *whiteShirtPrototype
		return &whiteShirtCopy
	case Black:
		blackShirtCopy := *blackShirtPrototype
		return &blackShirtCopy
	case Blue:
		blueShirtCopy := *blueShirtPrototype
		return &blueShirtCopy
	}
	return nil
}

type Shirt struct {
	Price float32
	SKU   string
	Color int
}

func (s *Shirt) GetInfo() string {
	return "something info here"
}
