package callbacks

import (
	"errors"
	"strconv"

	"github.com/dzendos/dubna/internal/model/state"
)

const (
	EditTransaction = "EditTransaction"
)

type CallbackHandler interface {
	SendReference(text string, userID int64) error
	ShowNotification(text string, userID int64, callbackID string) error
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
}

func (s *Model) IncomingCallback(data *CallbackData) error {
	switch data.Data {
	case EditTransaction:
		return s.toEditTransactionState(data)
	default:
		wasCallback, err := s.checkUserCallback(data)

		if wasCallback {
			return err
		}
	}

	return errors.New("Callback handler for data '" + data.Data + "' was not found.")
}

func (s *Model) checkUserCallback(data *CallbackData) (bool, error) {
	id, err := strconv.ParseInt(data.Data, 10, 64)
	if err != nil {
		return false, err
	}

	st, ok := state.GetUserState(id)

	if !ok {
		return false, nil
	}
	// TODO check if there is no race condition
	state.UserState[data.FromID] = state.NewUserState(st.CurrentRestaurant, st.CurrentOrder, st.OrderOrganizerID)

	return true, nil
}
