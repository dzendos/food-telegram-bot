package messages

import (
	"fmt"
	"log"

	db "github.com/dzendos/dubna/internal/database/temporary"
	"github.com/dzendos/dubna/internal/model/state"
)

type MessageSender interface {
	SendReference(text string, userID int64) error
	SetTransactionMessage(text string, userID int64) error
	SendMessage(text string, userID int64) error
	DeleteMessage(userID int64, messageID int64) error
	EditMessage(text string, userID int64, messageID int64) error
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
	MessageID int64
}

func (s *Model) IncomingMessage(msg *Message) error {
	log.Println(msg)
	// Trying to recognize the command.
	switch msg.Text {
	case "/start":
		return nil
	case "/my_order":
		text := state.OrderToString(msg.UserID)
		return s.tgClient.SendMessage(text, msg.UserID)
	case "/full_order":
		text := state.GetFullOrder(msg.UserID)
		return s.tgClient.SendMessage(text, msg.UserID)
	case "/confirm_order": // TODO add restaurant number
		text := "Готово! С данным рестораном можно связаться по телефону: 8800\nВы заказали:" + state.GetFullOrder(msg.UserID)
		debts := state.GetDebts(msg.UserID)

		message := db.GetUserTransaction(msg.UserID)
		for id, debt := range debts {
			s.tgClient.SendMessage(message+"\nОбщая сумма: "+fmt.Sprint(debt), id)
		}

		state.ResetUsers(msg.UserID)
		return s.tgClient.SendMessage(text, msg.UserID)
	case "/new_order":
		return s.newOrder(msg)
	case "/get_report":
		return s.getReport(msg)
	case "/set_transaction_message":
		return s.toEditState(msg.UserID, msg.MessageID)
	case "/cancel_order":
		return nil
	}

	if userState, ok := state.GetUserState(msg.UserID); ok {
		log.Println("userState")
		switch userState.EditState {
		case state.EditTransaction:
			return s.transactionEntered(msg)
		}
	}

	return nil
}
