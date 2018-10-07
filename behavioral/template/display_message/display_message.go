package display_message

import "fmt"

type TemplateDisplayer interface {
	open() string
	close() string
	Display(m MessageRetriever)
}

type MessageRetriever interface {
	Message() string
}

type templateDisplayer struct {}

func (t templateDisplayer) open() string {
	return ">>> Opening"
}

func (t templateDisplayer) close() string {
	return ">>> Closing"
}

func (t templateDisplayer) Display(m MessageRetriever) {
	fmt.Println(t.open(), m.Message(), t.close())
}

type helloMessageFunc func() string

func (h helloMessageFunc) Message() string {
	return h()
}