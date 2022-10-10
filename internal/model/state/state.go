// Package state contains FSA for all users.
package state

import (
	"fmt"

	"github.com/dzendos/dubna/internal/model/restaurant"
	"github.com/pkg/errors"
)

var UserState = make(map[int64]UserStateType)

type State int

const (
	RestaurantReference = "https://c259-188-130-155-154.eu.ngrok.io/restaurantPage.html"
	MenuReference       = "https://c259-188-130-155-154.eu.ngrok.io/mainPage.html"
)

const (
	EditTransaction State = iota + 1
)

type UserStateType struct {
	CurrentRestaurant *restaurant.Restaurant
	CurrentOrder      []OrderPosition
	EditState         State
	EditMessageID     int64
	OrderOrganizerID  int64
}

func NewUserState(r *restaurant.Restaurant, o []OrderPosition, id int64) UserStateType {
	return UserStateType{
		CurrentRestaurant: r,
		CurrentOrder:      o,
		OrderOrganizerID:  id,
	}
}

type OrderPosition struct {
	Name   string `json:"PositionName"`
	Amount int    `json:"PositionAmount"`
}

type Order struct {
	UserID int64           `json:"UserID"`
	Order  []OrderPosition `json:"Order"`
}

func GetUserState(userID int64) (UserStateType, bool) {
	state, ok := UserState[userID]
	return state, ok
}

func SetOrderOrganizer(userID int64, organizerID int64) error {
	state, ok := UserState[userID]

	if !ok {
		state = UserStateType{}
	}

	state.OrderOrganizerID = organizerID

	UserState[userID] = state

	return nil
}

func OrderToString(userID int64) string {
	var result string = "Ваш заказ готов:\n\n"

	state := UserState[userID]

	for _, pos := range state.CurrentOrder {
		result += "'" + pos.Name + "' " + fmt.Sprint(pos.Amount) + " шт\n"
	}

	return result
}

func SetMessageID(userID, messageID int64) int64 {
	state, ok := UserState[userID]
	if !ok {
		state = UserStateType{}
	}
	state.EditMessageID = messageID
	UserState[userID] = state

	return UserState[userID].EditMessageID
}

func GetFullOrder(userID int64) string {
	var result string = ""

	for id, state := range UserState {
		if state.OrderOrganizerID != userID {
			continue
		}

		result += fmt.Sprint(id) + ":\n"
		for _, pos := range state.CurrentOrder {
			result += "'" + pos.Name + "' " + fmt.Sprint(pos.Amount) + " шт\n"
		}
		result += "\n"
	}

	return result
}

func ResetUsers(userID int64) {
	for id, state := range UserState {
		if state.OrderOrganizerID != userID {
			continue
		}

		delete(UserState, id)
	}
}

func GetDebts(userID int64) (result map[int64]float64) {
	result = make(map[int64]float64)

	for id, state := range UserState {
		if state.OrderOrganizerID != userID {
			continue
		}
		var res float64 = 0
		for _, position := range state.CurrentOrder {
			res += float64(position.Amount) * positionCost(position.Name, state)
		}

		result[id] = res
	}

	return
}

func positionCost(name string, state UserStateType) float64 {
	for _, pos := range state.CurrentRestaurant.Menu.Positions {
		if pos.Name == name {
			return pos.Price
		}
	}

	return 0
}

func GetOrderOwner(userID int64) int64 {
	state, ok := UserState[userID]
	if !ok {
		state = UserStateType{
			OrderOrganizerID: userID,
		}
	}
	UserState[userID] = state

	return UserState[userID].OrderOrganizerID
}

func SetUserRestaurant(userID int64, restaurantName string) error {
	restaurant, err := GetRestaurantByName(restaurantName)
	if err != nil {
		return errors.Wrap(err, "state.SetUserRestaurant")
	}

	state, ok := UserState[userID]

	if !ok {
		state = UserStateType{}
	}

	state.CurrentRestaurant = restaurant
	state.OrderOrganizerID = userID

	UserState[userID] = state

	return nil
}

func GetUserRestaurant(userID int64) *restaurant.Restaurant {
	return UserState[userID].CurrentRestaurant
}

func SetState(userID int64, st State) {
	state, ok := UserState[userID]

	if !ok {
		state = UserStateType{}
	}

	state.EditState = st
	UserState[userID] = state
}

func SetUserOrder(order *Order) {
	state, ok := UserState[order.UserID]

	if !ok {
		state = UserStateType{}
	}

	state.CurrentOrder = order.Order
	UserState[order.UserID] = state
}

func GetRestaurantByName(restaurantName string) (*restaurant.Restaurant, error) {
	for _, restaurant := range ServerState.Restaurants {
		if restaurant.Name == restaurantName {
			return restaurant, nil
		}
	}

	return nil, errors.New("restaurant was not found")
}

func GetUserOrder(userID int64) []OrderPosition {
	return UserState[userID].CurrentOrder
}

func GetPositionPrice(userID int64, positionName string) float64 {
	restaurant := UserState[userID].CurrentRestaurant
	for _, position := range restaurant.Menu.Positions {
		if position.Name == positionName {
			return position.Price
		}
	}

	return -1
}

var ServerState = ServerStateType{}

type ServerStateType struct {
	Restaurants []*restaurant.Restaurant
}
