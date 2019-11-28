package shoes

import "errors"

const (
	white = iota
	black
	blue
)

type Cloner interface {
	Clone(s int) (ItemInfoGetter, error)
}

type ShoeCloner struct{}

func (ShoeCloner) Clone(s int) (ItemInfoGetter, error) {
	switch s {
	case white:
		newItem := *whiteShoePrototype
		return &newItem, nil
	default:
		return nil, errors.New("not yet implemented")
	}
}
