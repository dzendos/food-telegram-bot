package callbacks

import (
	db "github.com/dzendos/dubna/internal/database/temporary"
	"github.com/dzendos/dubna/internal/model/state"
)

// Here we need to store methods that are executed
// when one of the callback handlers have been activated.

func (s *Model) toEditTransactionState(data *CallbackData) error {
	currentTransactionMessage := db.GetUserTransaction(data.FromID)
	state.SetState(data.FromID, state.EditTransaction)
	return s.tgClient.ShowNotification(currentTransactionMessage, data.FromID, data.CallbackID)
}

func SendRestaurantMenu(userID int64) {

}
