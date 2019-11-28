package email

type MessageDirector struct {
	Text      string
	Recipient string
}

func (s *MessageDirector) BuildMessage(builder MessageBuilder) (*Message, error) {
	return builder.SetRecipient(s.Recipient).SetMessage(s.Text).Message()
}
