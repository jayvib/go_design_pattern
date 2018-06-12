package multi_logger

import "fmt"

type myTestWriter struct {
	receiveMessage *string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	if m.receiveMessage == nil {
		m.receiveMessage = new(string)
	}
	tempMessage := fmt.Sprintf("%s%s", *m.receiveMessage, p)
	m.receiveMessage = &tempMessage
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}
