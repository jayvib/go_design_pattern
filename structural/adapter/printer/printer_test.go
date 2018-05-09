package printer

import (
	"testing"
	"strings"
)

func TestModernPrinter(t *testing.T) {

	legacyPrinter := &LegacyPrinter{}
	legacyPrinter.SetMessage("hello printer!")

	modernPrinter := &ModernPrinter{}
	modernPrinter.SetOldPrinter(legacyPrinter)

	if !strings.Contains(modernPrinter.printMessage(), "|modern|") {
		t.Error("the expected word doesn't appear in the message")
	}

	superPrinter := &SuperModernPrinter{}
	superPrinter.SetModernPrinter(modernPrinter)
	superPrinter.Display()
}