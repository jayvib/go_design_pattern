package short

import (
	"errors"
	"fmt"
)

const (
	Black = iota
	Blue
	Red
)

type Short struct {
	Price float64
	SKU string
	Color string
}

func (s *Short) Info() string {
	return fmt.Sprintf("SKU: %s Price: %f Color: %s", s.SKU, s.Price, s.Color)
}

var BlackShortPrototype = &Short{
	Price: 10,
	SKU: "empty",
	Color: "Black",
}

var BlueShortPrototype = &Short{
	Price: 20,
	SKU: "empty",
	Color: "Blue",
}

var RedShortPrototype = &Short{
	Price: 30,
	SKU: "empty",
	Color: "Red",
}

type ItemInfo interface {
	Info() string
}

type ShortCloner struct{}

func (ShortCloner) GetClone(c int) (ItemInfo, error) {
	switch c {
	case Black:
		out := *BlackShortPrototype
		return &out, nil
	case Blue:
		out := *BlueShortPrototype
		return &out, nil
	case Red:
		out := *RedShortPrototype
		return &out, nil
	default:
		return nil, errors.New("short prototype not recognized")
	}
}
