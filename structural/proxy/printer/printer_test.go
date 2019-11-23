package printer_test

import (
	"github.com/jayvib/go_design_pattern/structural/proxy/printer"
	"log"
	"testing"
)

func TestPrinterProxy(t *testing.T) {

	printerProxy := printer.NewPrinterProxy("epson")

	if printerProxy.GetPrinterName() != "epson" {
		t.Errorf("expecting printer name 'epson' but got %s", printerProxy.GetPrinterName())
	}

	if printerProxy.GetMainPrinterName() != "main:epson" {
		t.Errorf("expecting main printer name 'main:epson' but got %s", printerProxy.GetPrinterName())
	}

	printerProxy.SetPrinterName("HP")

	if printerProxy.GetMainPrinterName() != "main:HP" {
		t.Errorf("expecting main printer name 'main:HP' but got %s", printerProxy.GetPrinterName())
	}

	if printerProxy.Print("Printing press") != "main:HP:Printing press" {
		t.Errorf("expecting print 'main:Printing press but receive %s", printerProxy.Print)
	}

	log.Printf(printerProxy.Print("Test Passed!"))

}
