package shoes

type ShoeStyle int

type ShoesFactory interface {
	CreateShoes(s ShoeStyle) Shoes
	AvailableStyles() []ShoeStyle
}


