// an example of template design where the one who use the template is embedded in the struct object.
package printer_2

import "fmt"

type PrinterTemplate interface {
	open() string
	print() string
	close() string
}

type AbstractPrinter struct{}

func (AbstractPrinter) Display(printer PrinterTemplate) string {
	return fmt.Sprintf("%s %s %s", printer.open(), printer.print(), printer.close())
}

type charDisplay struct {
	AbstractPrinter
	Char rune
}

func (c *charDisplay) open() string {
	return "<<"
}

func (c *charDisplay) print() string {
	return string(c.Char)
}

func (c *charDisplay) close() string {
	return ">>"
}
