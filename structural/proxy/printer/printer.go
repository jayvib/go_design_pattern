// This is an attempt to implement the proxy design pattern
package printer

type printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(str string) string
}

// printer acts as a real heavy object
type printer struct {
	name string
}

func (self *printer) SetPrinterName(name string) {

}

func (self *printer) GetPrinterName() string {
	return ""
}

func (self *printer) Print(str string) string {
	return ""
}

type PrinterProxy struct {
	Name string
	printer printable
}

func (self *PrinterProxy) SetPrinterName(name string) {

}

func (self *PrinterProxy) GetPrinterName() string {
	return ""
}

func (self *PrinterProxy) Print(str string) string {
	return ""
}

func (self *PrinterProxy) Init() {

}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{
		Name: name,
		printer: &printer{
			name: name,
		},
	}
}