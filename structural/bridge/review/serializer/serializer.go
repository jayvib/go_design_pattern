package serializer

import (
	"errors"
	"io"
)

type SerializerAPI interface {
	Serialize(arg interface{}) ([]byte, error)
}

type JSONSerializer struct {}

func (JSONSerializer) Serialize(arg interface{}) ([]byte, error) {
	return nil, errors.New("not implemented yet")
}

type XMLSerializer struct {}

func(XMLSerializer) Serialize(arg interface{}) ([]byte, error) {
	return nil, errors.New("not implemented yet")
}

// Printer describes a general printer object
type Printer interface{
	Print() error
}

type SuperSerializedPrinter struct {
	printer Printer
}

func (s *SuperSerializedPrinter) SetPrinter(p Printer) {
	s.printer = p
}

func (s *SuperSerializedPrinter) Do() error {
	return errors.New("not implemented yet")
}

type ConsolePrinter struct {}

func (ConsolePrinter) Print() error {
	return errors.New("not implemented yet")
}

type WriterPrinter struct {
	Out io.Writer
}

func (w *WriterPrinter) Print() error {
	return errors.New("not implemented yet")
}
