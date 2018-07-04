// an example of template design pattern where the PrinterTemplate implementation object is embedded in the
// object that who will use the template.
package printer

import "fmt"

type PrinterTemplate interface {
	open() string
	print() string
	close() string
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

func (a *AbstractPrinter) Display() string {
	return fmt.Sprintf("%s -> %s -> %s", a.PrinterTemplate.open(), a.PrinterTemplate.print(), a.PrinterTemplate.close())
}
