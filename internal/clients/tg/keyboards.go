// Package tg contains definitions for keyboards
// used inside the bot.
package tg

import (
	"fmt"

	tgbotapi "github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/dzendos/dubna/internal/model/callbacks"
	"github.com/dzendos/dubna/internal/model/state"
)

var chooseRestaurantKeyboard = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{
		{Text: "Выбрать ресторан", WebApp: &tgbotapi.WebAppInfo{Url: state.RestaurantReference}},
	}},
}

var getMenuKeyboard = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{
		{Text: "Составить заказ", WebApp: &tgbotapi.WebAppInfo{Url: state.MenuReference}},
	}},
}

func getShareOrderKeyboard(userID int64) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{
			{Text: "Присоединиться к заказу", CallbackData: fmt.Sprint(userID)},
		}},
	}
}

var editTransactionKeyboard = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{
		{Text: "Изменить", CallbackData: callbacks.EditTransaction}},
	},
}
