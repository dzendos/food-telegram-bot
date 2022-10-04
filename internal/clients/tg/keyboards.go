// Package tg contains definitions for keyboards
// used inside the bot.
package tg

import tgbotapi "github.com/PaulSonOfLars/gotgbot/v2"

var shopKeyboard = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{
		{Text: "Press me", WebApp: &tgbotapi.WebAppInfo{Url: "https://revenkroz.github.io/telegram-web-app-bot-example/index.html"}},
	}},
}
