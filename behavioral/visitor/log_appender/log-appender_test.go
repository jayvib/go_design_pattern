package log_appender

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

type TestHelper struct {
	Received string
}

func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

func Test_Overall(t *testing.T) {
	testHelper := new(TestHelper)
	visitor := new(MessageVisitor)

	t.Run("MessageA test", func(t *testing.T){
		msg := &MessageA{
			Msg: "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()

		assert.Equal(t, testHelper.Received, "A: Hello World (Visited A)", "expected result was incorrect")
	})

	t.Run("MessageB test", func(t *testing.T){
		msg := &MessageB{
			Msg: "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()
		assert.Equal(t, testHelper.Received, "B: Hello World (Visited B)", "expected result was incorrect")
	})
}
