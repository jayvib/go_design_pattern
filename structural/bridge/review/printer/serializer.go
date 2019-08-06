package serializer

// Description:
// Bridge Design Pattern decouples an abstraction
// from its implementation so that two can vary
// independently.

import (
	"fmt"
	"io"
)

type Serializer interface {
	Serialize(data interface{}) []byte
}

type JSONSerializer struct{}

func (JSONSerializer) Serialize(data interface{}) []byte {
	return nil
}

type XMLSerializer struct{}

func (XMLSerializer) Serialize(data interface{}) []byte {
	return nil
}

type PrinterAPI interface {
	Print(msg string) error
}

type PrinterAPIFunc func(msg string) error

func (p PrinterAPIFunc) Print(msg string) error {
	return p(msg)
}

type WriterPrinterAPI struct {
	out io.Writer
}

func (wp *WriterPrinterAPI) Print(msg string) error {
	return nil
}

type WriterSerializedPrinter struct {
	out        io.Writer
	serializer Serializer
}

// Will use the PrinterAPIFunc
func ConsolePrint(msg string) error {
	// use logrus :)
	fmt.Printf("Console:%s\n", msg)
	return nil
}

func NewSuperPrinter(printer PrinterAPI) SuperPrinter {
	return &SuperPrinterImpl{
		printer: printer,
	}
}

type SuperPrinter interface {
	Print(msg string) error
}

// This struct is the bridge for the implementation of
// the SuperPrinter
type SuperPrinterImpl struct {
	//s Serializer // No! This should be done in the PrinterAPI implementation
	//out io.Writer //No! This should be done in the PrinterAPI implementation
	printer PrinterAPI // assuming that the implementation for this api is subjected to modification
}

func (p *SuperPrinterImpl) Print(msg string) error {
	return p.printer.Print(msg)
}
