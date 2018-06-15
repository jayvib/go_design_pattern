package birthday_party

type NameRetriever interface {
	GetName() []string
}

type Visitor interface {
	Visit(retriever NameRetriever)
}


type Visitable interface {
	Accept(Visitor)
}
