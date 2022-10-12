package tg

import (
	"log"

	tgbotapi "github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/inlinequery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
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
	log.Println(ctx.EffectiveMessage)
	if ctx.CallbackQuery != nil {
		tgClient.callbackModel.IncomingCallback(&callbacks.CallbackData{
			FromID:     ctx.CallbackQuery.From.Id,
			Data:       ctx.CallbackQuery.Data,
			CallbackID: ctx.CallbackQuery.Id,
		})
	} else if ctx.Message != nil {
		tgClient.msgModel.IncomingMessage(&messages.Message{
			Text:      ctx.Message.Text,
			UserID:    ctx.Message.From.Id,
			MessageID: ctx.Message.MessageId,
		})
	}

	return nil
}

func (c *Client) DeleteMessage(userID int64, messageID int64) error {
	_, err := c.bot.DeleteMessage(userID, messageID, &tgbotapi.DeleteMessageOpts{})

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) EditMessage(text string, userID int64, messageID int64) error {
	log.Println(userID, messageID)
	_, b, err := c.bot.EditMessageText(text, &tgbotapi.EditMessageTextOpts{
		ChatId:    userID,
		MessageId: messageID,
	})

	log.Println(b)

	if err != nil {
		return err
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

func (c *Client) SendMessage(text string, userID int64) error {
	_, err := c.bot.SendMessage(userID, text, &tgbotapi.SendMessageOpts{
		ParseMode: "Markdown",
	})

	if err != nil {
		return errors.Wrap(err, "client.SendMes")
	}
	return nil
}

func (c *Client) SendReference(text string, userID int64) error {
	log.Println("SenD=Ref")
	_, err := c.bot.SendMessage(userID, text, &tgbotapi.SendMessageOpts{
		ParseMode:   "HTML",
		ReplyMarkup: chooseRestaurantKeyboard,
	})

	if err != nil {
		return errors.Wrap(err, "client.SendRef")
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

func (c *Client) SendRestaurantMenu(userID int64) error {
	_, err := c.bot.SendMessage(userID, "Меню готово! Перешлите сообщение для того чтобы поделиться заказом с друзьями", &tgbotapi.SendMessageOpts{
		ParseMode:   "HTML",
		ReplyMarkup: getMenuKeyboard,
	})

	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) SendOrderMenu(text string, userID int64) error {
	_, err := c.bot.SendMessage(userID, text, &tgbotapi.SendMessageOpts{
		ParseMode:   "HTML",
		ReplyMarkup: getMenuKeyboard,
	})

	if err != nil {
		return errors.Wrap(err, "client.SendOrderMenu")
	}
	return nil
}

func answerInlineQuery(bot *tgbotapi.Bot, ctx *ext.Context) error {
	log.Println("caught")

	markup := getShareOrderKeyboard(ctx.EffectiveUser.Id)
	log.Println(ctx.EffectiveUser.Id)
	// TODO add check - whether we have order or not
	ShareMyOrder := tgbotapi.InlineQueryResultArticle{
		Id:                  "ShareMyOrder",
		Title:               "Поделиться моим заказом",
		Description:         "Отправить ссылку для доступа к моему заказу",
		ReplyMarkup:         &markup,
		InputMessageContent: tgbotapi.InputTextMessageContent{MessageText: "Приввет, вот ссылка для подключения к моему заказу"},
	}

	bot.AnswerInlineQuery(ctx.InlineQuery.Id, []tgbotapi.InlineQueryResult{
		ShareMyOrder,
	}, &tgbotapi.AnswerInlineQueryOpts{CacheTime: 1})

	return nil
}

func (c *Client) GetUserByID(userID int64) string {
	user, _ := c.bot.GetChat(userID, &tgbotapi.GetChatOpts{})
	return user.FirstName + " " + user.LastName
}

func (c *Client) ListenUpdates(msgModel *messages.Model, callbackModel *callbacks.Model) {
	c.msgModel = msgModel
	c.callbackModel = callbackModel

	c.dispatcher.AddHandler(handlers.NewCommand("start", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCommand("help", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCommand("my_order", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCommand("full_order", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCommand("confirm_order", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCommand("new_order", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCommand("get_report", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCommand("set_transaction_message", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCommand("cancel_order", incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewMessage(message.All, incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewCallback(callbackquery.All, incomingUpdate))
	c.dispatcher.AddHandler(handlers.NewInlineQuery(inlinequery.All, answerInlineQuery))

	//c.dispatcher.AddHandler(handlers.NewChosenInlineResult(choseninlineresult.All, incomingUpdate))

	err := c.updater.StartPolling(c.bot, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
}
