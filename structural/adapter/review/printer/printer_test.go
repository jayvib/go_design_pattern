package printer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := &PrinterAdapter{
		OldPrinter: new(MyLegacyPrinter),
		Msg:        msg,
	}

	expectedMsg := "Legacy Printer: Adapter: Hello World!\n"
	msgResult := adapter.PrintStored()
	assert.Equal(t, expectedMsg, msgResult, "Message doesn't match")
}
