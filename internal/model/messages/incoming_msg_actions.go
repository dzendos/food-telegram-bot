package messages

import (
	db "github.com/dzendos/dubna/internal/database/temporary"
)

const (
	SendReferenceMessage = "Привет, создаем новый заказ. Выберите ресторан для создания заказа."
)

func (s *Model) toEditState(userID int64) error {
	return s.tgClient.SetTransactionMessage("Введите новое сообщение", userID)
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
