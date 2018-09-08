package log_appender

import (
	"io"
	"fmt"
	"os"
)

type MessageA struct {
	Msg string
	Output io.Writer
}

func(m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

type MessageB struct {
	Msg string
	Output io.Writer
}

func(mb *MessageB) Accept(v Visitor) {
	v.VisitB(mb)
}

func(mb *MessageB) Print() {
	if mb.Output == nil {
		mb.Output = os.Stdout
	}
	fmt.Fprintf(mb.Output, "B: %s", mb.Msg)
}

type Visitor interface {
	VisitA(a *MessageA)
	VisitB(b *MessageB)
}

type Visitable interface {
	Accept(Visitor)
}

type MessageVisitor struct {}

func (mv *MessageVisitor) VisitA(a *MessageA) {
	a.Msg = fmt.Sprintf("%s %s", a.Msg, "(Visited A)")
}

func (mv *MessageVisitor) VisitB(b *MessageB) {
	b.Msg = fmt.Sprintf("%s %s", b.Msg, "(Visited B)")
}