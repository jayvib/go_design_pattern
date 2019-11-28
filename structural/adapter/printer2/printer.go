package printer2

import "fmt"

// LegacyPrinter represents a Legacy Printer
type LegacyPrinter interface {
	Print(s string) string
}

// MyLegacyPrinter implements LegacyPrinter and is a very productive printer.
type MyLegacyPrinter struct{}

func (MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return newMsg
}

// ModernPrinter represents a modern printer very high tech.
type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	return ""
}
