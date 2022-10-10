package callbacks

import (
	"log"
	"strconv"

	"github.com/dzendos/dubna/internal/model/state"
)

const (
	EditTransaction = "EditTransaction"
)

type CallbackHandler interface {
	SendReference(text string, userID int64) error
	ShowNotification(text string, userID int64, callbackID string) error
	SendOrderMenu(text string, userID int64) error
}

type Model struct {
	tgClient CallbackHandler
}

func New(tgClient CallbackHandler) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type CallbackData struct {
	FromID     int64
	Data       string
	CallbackID string
	OwnerID    int64
}

func (s *Model) IncomingCallback(data *CallbackData) error {
	switch data.Data {
	case EditTransaction:
		return s.toEditTransactionState(data)
	default:
		wasCallback, err := s.checkUserCallback(data)

		if !wasCallback {
			return err
		}

		return s.tgClient.SendOrderMenu("Привет, заказ готов", data.FromID)
	}
}

func (s *Model) checkUserCallback(data *CallbackData) (bool, error) {
	senderID, _ := strconv.ParseInt(data.Data, 10, 64)
	log.Println(senderID)
	st, ok := state.GetUserState(senderID)

	if !ok {
		return false, nil
	} else {
		log.Println("this is the user")
	}
	// TODO check if there is no race condition

	state.UserState[data.FromID] = state.NewUserState(st.CurrentRestaurant, st.CurrentOrder, st.OrderOrganizerID)

	return true, nil
}
