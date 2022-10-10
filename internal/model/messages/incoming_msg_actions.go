package messages

import (
	db "github.com/dzendos/dubna/internal/database/temporary"
	"github.com/dzendos/dubna/internal/model/state"
)

const (
	SendReferenceMessage = "Привет, создаем новый заказ. Выберите ресторан для создания заказа."
)

func (s *Model) toEditState(userID int64, messageID int64) error {
	currentTransactionMessage := db.GetUserTransaction(userID)
	state.SetMessageID(userID, messageID)
	return s.tgClient.SetTransactionMessage("Ваше текущее сообщение:\n"+currentTransactionMessage, userID)
}

func (s *Model) transactionEntered(msg *Message) error {
	s.tgClient.DeleteMessage(msg.UserID, msg.MessageID)

	messageID := state.UserState[msg.UserID].EditMessageID

	s.tgClient.EditMessage("Ваше текущее сообщение:\n"+msg.Text, msg.UserID, messageID)
	state.SetState(msg.UserID, 0)

	return db.SetUserTransaction(msg.Text, msg.UserID)
}

func (s *Model) newOrder(msg *Message) error {
	return s.tgClient.SendReference(SendReferenceMessage, msg.UserID)
}

func (s *Model) getReport(msg *Message) error {

	return nil
}
