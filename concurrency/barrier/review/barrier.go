package review

import (
	"bytes"
	"io"
	"os"
)

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	barrier(endpoints...)
	writer.Close() // close the writer to signal the goroutine that no more input is going to come to it.
	temp := <-out
	return temp
}
