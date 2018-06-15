package online_shopping

import "fmt"

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
	Names string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.Names = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}
