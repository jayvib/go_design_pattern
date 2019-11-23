package json

import (
	"encoding/json"
	"github.com/jayvib/go_design_pattern/creational/builder/message-to-santa"
)

type JSONMessageBuilder struct {
	Recipient string
	Msg       string
}

func (j *JSONMessageBuilder) SetRecipient(recipient string) message_to_santa.BuilderDirector {
	j.Recipient = recipient
	return j
}

func (j *JSONMessageBuilder) SetText(msg string) message_to_santa.BuilderDirector {
	j.Msg = msg
	return j
}

func (j *JSONMessageBuilder) GetMessage() (*message_to_santa.Message, error) {
	b, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	m := &message_to_santa.Message{
		Data:   b,
		Format: "JSON",
	}
	return m, nil
}
