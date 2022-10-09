package messages

import (
	db "github.com/dzendos/dubna/internal/database/temporary"
)

const (
	SendReferenceMessage = "Привет, создаем новый заказ. Выберите ремторан для создания заказа."
)

func (s *Model) toEditState(userID int64) error {
	currentTransactionMessage := db.GetUserTransaction(userID)
	return s.tgClient.SetTransactionMessage(currentTransactionMessage, userID)
}

func (s *Model) transactionEntered(msg *Message) error {
	return db.SetUserTransaction(msg.Text, msg.UserID)
}

func (s *Model) newOrder(msg *Message) error {
	return s.tgClient.SendReference(SendReferenceMessage, msg.UserID)
}

func (s *Model) getReport(msg *Message) error {

	return nil
}
