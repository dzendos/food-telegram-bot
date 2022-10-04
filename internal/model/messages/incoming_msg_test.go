package messages

import (
	"testing"

	mocks "github.com/dzendos/dubna/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("hello", int64(123))

	err := model.IncomingMessage(&Message{
		Text:   "/start",
		UserID: 123,
	})

	assert.NoError(t, err)
}
