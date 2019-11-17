package email

import "encoding/xml"

type XMLMessageBuilder struct {
	messageRecipient string
	messageText      string
}

func (b *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.messageRecipient = recipient
	return b
}

func (b *XMLMessageBuilder) SetMessage(msg string) MessageBuilder {
	b.messageText = msg
	return b
}

func (b *XMLMessageBuilder) Message() (*Message, error) {
	c := &container{
		b.messageRecipient,
		b.messageText,
	}

	payload, err := xml.Marshal(c)
	if err != nil {
		return nil, err
	}

	return &Message{
		Body: payload,
		Format: "xml",
	}, nil
}
