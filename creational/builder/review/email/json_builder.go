package email

import "encoding/json"

// Trying to solidify my understanding of Builder Design Pattern
//
// The Builder Pattern is comprised of four components:
// - builder
// - concrete builder
// - director
// - product

type JSONMessageBuilder struct {
	recipient string
	message   string
}

func (b *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.recipient = recipient
	return b
}

func (b *JSONMessageBuilder) SetMessage(msg string) MessageBuilder {
	b.message = msg
	return b
}

type container struct {
	Recipient string `json:"recipient" xml:"recipient"`
	Message   string `json:"message" xml:"message"`
}

func (b *JSONMessageBuilder) Message() (*Message, error) {
	c := &container{
		Recipient: b.recipient,
		Message:   b.message,
	}
	payload, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return &Message{
		Body:   payload,
		Format: "json",
	}, nil
}
