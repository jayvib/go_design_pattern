package house

type HouseComponentStructurer interface {
	SetPost() HouseComponentStructurer
	SetDoor() HouseComponentStructurer
	SetWindow() HouseComponentStructurer
	SetRoof() HouseComponentStructurer
	GetHouse() *House
}

type BuildComplex struct {
	structurer HouseComponentStructurer
}

func NewBuildComplex(s HouseComponentStructurer) *BuildComplex {
	return &BuildComplex{
		structurer: s,
	}
}

func (b *BuildComplex) Build() {
	b.structurer.SetPost().SetRoof().SetDoor().SetWindow()
}
