package printer2

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World"
	expected := "Legacy Printer: " + msg

	adapter := &PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg: msg,
	}

	if adapter.PrintStored() != expected {
		t.Errorf("expected '%s' but got '%s'\n", expected, adapter.PrintStored())
	}
}
