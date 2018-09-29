package message_to_santa

// Message represents the item that will be constructed
type Message struct {
	Data []byte
	Format string
}

type message struct {
	recipient string
	text string
}

type BuilderMessage interface {
	SetRecipient(recipient string) BuilderMessage
	SetText(msg string) BuilderMessage
	GetMessage() (*Message, error)
}

// Director...
type BuilderDirector struct {
	BuilderMsg BuilderMessage
	Msg        message
}


func (b *BuilderDirector) Build() (*Message, error){
	return b.BuilderMsg.SetRecipient(b.Msg.recipient).SetText(b.Msg.text).GetMessage()
}

// Define the Object Implementes the BuilderMessage





