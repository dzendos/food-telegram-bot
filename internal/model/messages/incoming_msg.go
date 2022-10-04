package messages

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text      string
	UserID    int64
	MessageID int
}

func (s *Model) IncomingMessage(msg *Message) error {
	// Trying to recognize the command.
	switch msg.Text {
	case "/start":
		return s.tgClient.SendMessage("hello", msg.UserID)
	}

	// It is not a known command - maybe it is message to change the state.

	return s.tgClient.SendMessage("не знаю эту команду", msg.UserID)
}
