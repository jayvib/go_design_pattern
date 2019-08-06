package serializer

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperPrinter(t *testing.T) {
	// I'll test three implementations
	// The console print
	// The print to writer
	// The print the writer with serialized data

	t.Run("Console Print", func(t *testing.T) {
		// Initialize object for the PrinterAPI
		// Im gonna use the PrinterAPIFunc
		// to implement the PrinterAPI interface

		// How will I going to catch the data???
		// gonna redirect the value of the
		// stdout: use os.Pipe
		stdOutOld := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		outC := make(chan string)
		go func() {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			outC <- buf.String()
		}()

		var printer PrinterAPI
		printer = PrinterAPIFunc(ConsolePrint)
		superPrinter := NewSuperPrinter(printer)
		msg := "Hello World!"
		err := superPrinter.Print(msg)

		w.Close()
		os.Stdout = stdOutOld
		out := <-outC // Get the value
		t.Log("Got:", out)
		assert.False(t, err != nil, "Error found", err)
		assert.Contains(t, out, "Hello", "Not Hello found")
		assert.Contains(t, out, "Console", "No Console found")
	})

	t.Run("Print to Writer", func(t *testing.T) {
		t.SkipNow()
	})

	t.Run("Print to Writer with Serialization", func(t *testing.T) {
		t.SkipNow()

	})

}
