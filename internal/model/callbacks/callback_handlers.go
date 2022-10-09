package callbacks

import (
	"errors"
)

const (
	EditTransaction = "EditTransaction"
)

type CallbackHandler interface {
	SendReference(text string, userID int64) error
	ShowNotification(text string, userID int64, callbackID string) error
}

type Model struct {
	tgClient CallbackHandler
}

func New(tgClient CallbackHandler) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type CallbackData struct {
	FromID     int64
	MessageID  int
	Data       string
	CallbackID string
}

func (s *Model) IncomingCallback(data *CallbackData) error {
	switch data.Data {
	case EditTransaction:
		s.toEditTransactionState(data)
	}

	return errors.New("Callback handler for data '" + data.Data + "' was not found.")
}
