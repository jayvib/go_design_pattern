package multi_logger

import (
	"strings"
	"testing"
)

func TestCreateDefaultChain(t *testing.T) {
	myWriter := myTestWriter{}

	writerLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it founds "+
		"the word 'hello', third writes to some variable if second found 'hello'",
		func(t *testing.T) {
			chain.Next("message that breaks the chain\n")

			if myWriter.receiveMessage != nil {
				t.Fatal("Last link should not receive any message")
			}

			chain.Next("Hello\n")

			if !strings.Contains(*myWriter.receiveMessage, "Hello") {
				t.Fatal("Last link didn't received expected message")
			}
		})

	t.Run("2 loggers, second uses the closure implementations", func(t *testing.T) {

		myWriter = myTestWriter{}
		closureLogger := &ClosureChain{Closure: func(s string) {
			myWriter.receiveMessage = &s
		}}

		writerLogger.NextChain = closureLogger

		chain.Next("Hello closure logger")

		if *myWriter.receiveMessage != "Hello closure logger" {
			t.Fatal("Expected message wasn't received in myWriter")
		}

	})
}
