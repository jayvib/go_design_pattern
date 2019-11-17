package email

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageDirector_BuildMessage(t *testing.T) {
	t.Run("JSON Builder", func(t *testing.T){
		msgDirector := &MessageDirector{
			Text: "This is the body of the message",
			Recipient: "Luffy Monkey",
		}
		msg, _ := msgDirector.BuildMessage(&JSONMessageBuilder{})
		wantMsg := `{"recipient":"Luffy Monkey","message":"This is the body of the message"}`
		assertJSONBuilder(t, msg, wantMsg)
	})

	t.Run("XML Builder", func(t *testing.T){
		msgDirector := &MessageDirector{
			Text: "This is the body of the message",
			Recipient: "Luffy Monkey",
		}
		msg, _ := msgDirector.BuildMessage(&XMLMessageBuilder{})
		wantMsg := `<container><recipient>Luffy Monkey</recipient><message>This is the body of the message</message></container>`
		assert.Equal(t, "xml", msg.Format)
		assert.Equal(t, wantMsg, string(msg.Body))
	})
}

func assertJSONBuilder(t *testing.T, msg *Message, msgWant string) {
	t.Helper()
	assert.Equal(t,  msgWant, string(msg.Body))
	assert.Equal(t, "json", msg.Format)
}