package printer

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct {
}

func (l *MyLegacyPrinter) Print(s string) string {
	msg := fmt.Sprintf("Legacy Printer: %s\n", s)
	log.Println(msg)
	return msg
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		newMsg := p.OldPrinter.Print(fmt.Sprintf("Adapter: %s", p.Msg))
		return newMsg
	}

	return p.Msg
}
