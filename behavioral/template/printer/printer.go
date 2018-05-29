package printer

import "fmt"

type PrinterTemplate interface {
	open() string
	print() string
	close() string
	Display() string
}

func NewAbstractPrinter(p PrinterTemplate) *AbstractPrinter {
	return &AbstractPrinter{
		PrinterTemplate: p,
	}
}

type AbstractPrinter struct {
	PrinterTemplate
}

type UserPrinter struct {
	content string
}

func (u *UserPrinter) open() string {
	return "opening"
}

func (u *UserPrinter) print() string {
	return u.content
}

func (u *UserPrinter) close() string {
	return "closing"
}

func (u *UserPrinter) Display() string {
	return fmt.Sprintf("%s -> %s -> %s", u.open(), u.print(), u.close())
}
