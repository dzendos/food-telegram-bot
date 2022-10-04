package tg

import (
	"log"

	"github.com/dzendos/dubna/internal/model/callbacks"
	"github.com/dzendos/dubna/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type tokenGetter interface {
	Token() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tokenGetter tokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())
	if err != nil {
		return nil, errors.Wrap(err, "cannot create NewBotAPI")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userID int64) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userID, text))
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) ShowAlert(text string, messageID string) error {
	alert := tgbotapi.NewCallback(messageID, text)

	_, err := c.client.Send(alert)

	if err != nil {
		return errors.Wrap(err, "client.ShowAlert")
	}

	return nil
}

func (c *Client) DeleteMessage(userID int64, messageID int) error {
	deleteMessage := tgbotapi.NewDeleteMessage(userID, messageID)
	_, err := c.client.Send(deleteMessage)

	if err != nil {
		return errors.Wrap(err, "client.EditMessage")
	}

	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model, callbackModel *callbacks.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := msgModel.IncomingMessage(&messages.Message{
				Text:      update.Message.Text,
				UserID:    update.Message.From.ID,
				MessageID: update.Message.MessageID,
			})
			if err != nil {
				log.Println("error processing message: ", err)
			}
		} else if update.CallbackQuery != nil {
			log.Printf("[%s] data: %s",
				update.CallbackQuery.Message.From.UserName,
				update.CallbackQuery.Data,
			)

			err := callbackModel.IncomingCallback(&callbacks.CallbackData{
				FromID:     update.CallbackQuery.From.ID,
				MessageID:  update.CallbackQuery.Message.MessageID,
				Data:       update.CallbackData(),
				CallbackID: update.CallbackQuery.ID,
			})

			if err != nil {
				log.Println("error processing data: ", err)
			}
		}
	}
}
