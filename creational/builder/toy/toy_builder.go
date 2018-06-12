package toy

type Builder interface {
	SetHead() Builder
	SetBody() Builder
	SetArms() Builder
	SetLeg() Builder
	GetProduct() string
}

type BuildComplex struct {
	builder Builder
}

func NewBuildComplex(b Builder) *BuildComplex {
	//return &BuildComplex{ builder:b }
	return nil
}

func (b *BuildComplex) SetNewBuilder(builder Builder) {
	//b.builder = builder
}

func (b *BuildComplex) Construct() {

}

type Toy struct {
	Head, Body, Arms, Legs string
}

type Android struct {
	Toy
}

func (a *Android) SetHead() {

}

func (a *Android) SetBody() {

}

func (a *Android) SetArms() {

}

func (a *Android) SetLegs() {

}

func (a *Android) GetProduct() string {
	return ""
}
