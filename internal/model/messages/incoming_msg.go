package messages

import (
	"log"

	"github.com/dzendos/dubna/internal/model/state"
)

type MessageSender interface {
	SendReference(text string, userID int64) error
	SetTransactionMessage(text string, userID int64) error
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
	log.Println(msg)
	// Trying to recognize the command.
	switch msg.Text {
	case "/start":
		return nil
	case "/new_order":
		return s.newOrder(msg)
	case "/get_report":
		return s.getReport(msg)
	case "/set_transaction_message":
		return s.toEditState(msg.UserID)
	case "/cancel_order":
		return nil
	}

	if userState, ok := state.GetUserState(msg.UserID); ok {
		switch userState.EditState {
		case state.EditTransaction:
			return s.transactionEntered(msg)
		}
	}

	// It is not a known command - maybe it is message to change the state.

	return s.tgClient.SendReference("не знаю эту команду", msg.UserID)
}
