package tg

import (
	"log"

	tgbotapi "github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/dzendos/dubna/internal/model/callbacks"
	"github.com/dzendos/dubna/internal/model/messages"
	"github.com/pkg/errors"
)

type tokenGetter interface {
	Token() string
}

var tgClient *Client

type Client struct {
	bot           *tgbotapi.Bot
	dispatcher    *ext.Dispatcher
	updater       *ext.Updater
	msgModel      *messages.Model
	callbackModel *callbacks.Model
}

func New(tokenGetter tokenGetter) (*Client, error) {
	bot, err := tgbotapi.NewBot(tokenGetter.Token(), &tgbotapi.BotOpts{
		// Client: http.Client{},
		DefaultRequestOpts: &tgbotapi.RequestOpts{
			Timeout: tgbotapi.DefaultTimeout,
			APIURL:  tgbotapi.DefaultAPIURL,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot create NewBotAPI")
	}

	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			ErrorLog: nil,
			Error: func(b *tgbotapi.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println(("an error occurred while handling update:" + err.Error()))
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		},
	})
	dispatcher := updater.Dispatcher

	tgClient = &Client{
		bot:        bot,
		dispatcher: dispatcher,
		updater:    &updater,
	}

	return tgClient, nil
}

func incomingUpdate(bot *tgbotapi.Bot, ctx *ext.Context) error {
	log.Printf("update caught: %s", ctx.Message.Text)

	if ctx.CallbackQuery != nil {

	} else if ctx.Message != nil {
		tgClient.msgModel.IncomingMessage(&messages.Message{
			Text:      ctx.Message.Text,
			UserID:    ctx.Message.From.Id,
			MessageID: int(ctx.Message.MessageId),
		})
	}

	return nil
}

func (c *Client) ShowNotification(text string, userID int64, callbackID string) error {
	_, err := c.bot.AnswerCallbackQuery(callbackID, &tgbotapi.AnswerCallbackQueryOpts{
		Text: text,
	})

	if err != nil {
		return errors.Wrap(err, "client.ShowNotification")
	}
	return nil
}

func (c *Client) SendReference(text string, userID int64) error {
	_, err := c.bot.SendMessage(userID, text, &tgbotapi.SendMessageOpts{
		ParseMode:   "HTML",
		ReplyMarkup: chooseRestaurantKeyboard,
	})

	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) SetTransactionMessage(text string, userID int64) error {
	_, err := c.bot.SendMessage(userID, text, &tgbotapi.SendMessageOpts{
		ReplyMarkup: editTransactionKeyboard,
	})

	if err != nil {
		return errors.Wrap(err, "client.SetTransactionMessage")
	}
	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model, callbackModel *callbacks.Model) {
	c.msgModel = msgModel
	c.callbackModel = callbackModel

	c.dispatcher.AddHandler(handlers.NewCommand("start", incomingUpdate))

	err := c.updater.StartPolling(c.bot, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
}
