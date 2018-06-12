package printer

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}

	_, err := d.Writer.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("content received on Writer was empty")
	return
}

type PrinterAbstraction interface {
	Print() error
}

type SuperPrinter struct {
	Printer PrinterAbstraction
}

func (s *SuperPrinter) SetPrinter(printer PrinterAbstraction) {
	s.Printer = printer
}

func (s *SuperPrinter) Do() error {
	if s.Printer == nil {
		return errors.New("You must provide a Printer value.")
	}
	err := s.Printer.Print()
	return err
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	err := c.Printer.PrintMessage(c.Msg)
	if err != nil {
		return err
	}
	return nil
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (pp *PacktPrinter) Print() error {
	err := pp.Printer.PrintMessage(fmt.Sprintf("Greeting from PACKT: %s", pp.Msg))
	if err != nil {
		return err
	}
	return nil
}
