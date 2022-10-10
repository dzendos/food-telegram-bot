package temporary

import "log"

var DefaultTransactionMessage string = "Привет! Заказ успешно оформлен, напишите мне в личное сообщение чтобы перевести за него"

var UserSettings = make(map[int64]UserSettingsType)

type UserSettingsType struct {
	TransactionMessage string
}

func NewUserSettings() *UserSettingsType {
	return &UserSettingsType{
		TransactionMessage: DefaultTransactionMessage,
	}
}

func GetUserTransaction(userID int64) string {
	settings, ok := UserSettings[userID]

	if !ok {
		settings = *NewUserSettings()
	}
	UserSettings[userID] = settings
	return UserSettings[userID].TransactionMessage
}

func SetUserTransaction(text string, userID int64) error {
	log.Println(text)
	settings, ok := UserSettings[userID]

	if !ok {
		settings = *NewUserSettings()
	}

	settings.TransactionMessage = text
	UserSettings[userID] = settings

	return nil
}
