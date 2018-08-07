package shirt

import "errors"

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

// GetShirtsCloner returns an interface ShirtCloner
func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

type ShirtsCache struct{}

func (sc *ShirtsCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}
}