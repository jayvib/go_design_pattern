package log

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

func TestWriteLogger(t *testing.T) {
	w := &mockWriter{}
	l := WriteLogger{
		w: w,
	}
	msg := "Hello World"
	visitor := &visitorImpl{Msg: msg}
	l.Accept(visitor)

	l.Log()

	// Check the output
	assert.Equal(t, msg, w.msg, "message not equal")
}
