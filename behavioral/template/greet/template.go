package greet

import "strings"

type MessageFunc func() string

func (m MessageFunc) Message() string {
	return m()
}

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(retriever MessageRetriever) string
}

type TemplateImp struct{}

func (t *TemplateImp) first() string {
	return "hello"
}

func (t *TemplateImp) third() string {
	return "template"
}

func (t *TemplateImp) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}
