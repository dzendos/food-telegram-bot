package callbacks

import (
	"log"

	"github.com/dzendos/dubna/internal/model/state"
)

// Here we need to store methods that are executed
// when one of the callback handlers have been activated.

func (s *Model) toEditTransactionState(data *CallbackData) error {
	state.SetState(data.FromID, state.EditTransaction)
	log.Println(state.GetUserState(data.FromID))
	return s.tgClient.ShowNotification("Введите новое сообщение", data.FromID, data.CallbackID)
}
