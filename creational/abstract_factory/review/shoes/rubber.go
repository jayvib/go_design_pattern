package shoes

const (
	Running ShoeStyle = iota
	BasketBall
)

type RubberShoesFactory struct{}

func (r RubberShoesFactory) CreateShoes(s ShoeStyle) Shoes {
	switch s {
	case Running:
	case BasketBall:
	}
	return nil
}

func (r RubberShoesFactory) AvailableStyles() []ShoeStyle {
	return nil
}

type runningShoes struct{}

func (r runningShoes) Lace() string        { return "" }
func (r runningShoes) Sole() string        { return "" }
func (r runningShoes) Upper() string       { return "" }
func (r runningShoes) Vamp() string        { return "" }
func (r runningShoes) HeelCounter() string { return "" }

type basketballShoes struct{}

func (b basketballShoes) Lace() string        { return "" }
func (b basketballShoes) Sole() string        { return "" }
func (b basketballShoes) Upper() string       { return "" }
func (b basketballShoes) Vamp() string        { return "" }
func (b basketballShoes) HeelCounter() string { return "" }
