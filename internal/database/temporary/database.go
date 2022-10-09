package temporary

var DefaultTransactionMessage string = "Привет! Заказ успешно оформлен, здесь ты можешь ознакомиться с составом заказа:"

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

	return settings.TransactionMessage
}

func SetUserTransaction(text string, userID int64) error {
	settings, ok := UserSettings[userID]

	if !ok {
		settings = *NewUserSettings()
	}

	settings.TransactionMessage = text
	UserSettings[userID] = settings

	return nil
}
