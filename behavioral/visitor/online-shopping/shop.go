package online_shopping

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

func (p *Product) GetName() string {
	return p.Name
}

type Rice struct {
	Product
}

type Paste struct {
	Product
}
