package callbacks

import (
	"errors"
)

const (
// Place for keyboards data
)

type CallbackHandler interface {
	SendMessage(text string, userID int64) error
	ShowAlert(text string, messageID string) error
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

	}

	return errors.New("Callback handler for data '" + data.Data + "' was not found.")
}
