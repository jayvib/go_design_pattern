package email

type MessageBuilder interface {
	SetRecipient(recipient string) MessageBuilder
	SetMessage(msg string) MessageBuilder
	Message() (*Message, error)
}

type Message struct {
	Body   []byte
	Format string
}
