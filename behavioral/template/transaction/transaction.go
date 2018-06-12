package transaction

import (
	"fmt"
	"strings"
)

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return "hello"
}

func (t *TemplateImpl) third() string {
	return "template"
}

func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type TestStruct struct {
	Template
}

func (m *TestStruct) Message() string {
	return "world"
}

type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(m MessageRetriever) string {
	return fmt.Sprintf("%s %s %s", a.first(), m.Message(), a.third())
}

type TemplateAdapterFunc func() string

func (t TemplateAdapterFunc) Message() string {
	return t()
}

func NewMessageRetrieverAdapter(f func() string) MessageRetriever {
	return TemplateAdapterFunc(f)
}
