package online_shopping

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Visitor interface {
	Visit(retriever ProductInfoRetriever)
}

type Visitable interface {
	Accept(Visitor)
}
