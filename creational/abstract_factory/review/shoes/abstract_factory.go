package shoes

type FactoryType int

const (
	sneaker FactoryType = iota
	rubber
)

var DefaultAbstractShoeFactory = &AbstractShoeFactory{
	factories: map[FactoryType]ShoesFactory{
		sneaker: &SneakerShoesFactory{},
		rubber:  &RubberShoesFactory{},
	},
}

type AbstractShoeFactory struct {
	factories map[FactoryType]ShoesFactory
}

func (a AbstractShoeFactory) CreateFactory(factory FactoryType) ShoesFactory {
	return a.factories[factory]
}

func (a *AbstractShoeFactory) AddFactory(t FactoryType, factory ShoesFactory) *AbstractShoeFactory {
	a.factories[t] = factory
	return a
}
