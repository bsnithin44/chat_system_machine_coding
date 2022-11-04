package msgs

type Message struct {
	Id     string
	Text   string
	Sender string
}

func NewMessage(text string) Message {
	return Message{
		Text: text,
	}
}

func (m *Message) setSender(senderID string) {
	m.Sender = senderID
}
