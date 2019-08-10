package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockWriter struct {
	msg string
}

func (m *mockWriter) Write(b []byte) (int, error) {
	m.msg = string(b)
	return len(b), nil
}

func TestChainLogger(t *testing.T) {
	// Console logger embedded to the writer logger

	consoleLog := new(consoleLogger)

	w := new(mockWriter)
	writerLog := &writeLogger{
		chain: consoleLog,
		w:     w,
	}
	msg := "Hello World"
	writerLog.Next(msg)

	// Check the writer
	assert.Equal(t, msg, w.msg, "message is not the same as expected")
}
