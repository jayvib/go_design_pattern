// This is an attempt to implement the proxy design pattern
// Instead of actual printer will be use by the client it will be
// the proxy printer.
package printer

type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	GetMainPrinterName() string
	Print(str string) string
}

// printer acts as a real heavy object
// make it private so that the client don't have a direct access to the actual printer
type printer struct {
	name string
}

func (self *printer) SetPrinterName(name string) {
	self.name = "main:" + name
}

func (self *printer) GetPrinterName() string {
	return self.name
}

func (self *printer) Print(str string) string {
	return self.name + ":" + str
}

type PrinterProxy struct {
	Name string
	*printer
}

func (self *PrinterProxy) SetPrinterName(name string) {
	self.Name = name
	self.printer.SetPrinterName(name)
}

func (self *PrinterProxy) GetPrinterName() string {
	return self.Name
}

func (self *PrinterProxy) Print(str string) string {
	return self.printer.Print(str)
}

func (self *PrinterProxy) GetMainPrinterName() string {
	return self.printer.GetPrinterName()
}

func NewPrinterProxy(name string) Printable {
	pp := &PrinterProxy{
		Name:    name,
		printer: new(printer),
	}

	pp.printer.SetPrinterName(name)

	return pp
}
