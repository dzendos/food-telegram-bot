// Package state contains FSA for all users.
package state

import (
	"github.com/dzendos/dubna/internal/model/restaurant"
	"github.com/pkg/errors"
)

var UserState = make(map[int64]UserStateType)

type State int

const (
	RestaurantReference = "https://c7d5-188-130-155-154.eu.ngrok.io/restaurantPage.html"
	MenuReference       = "https://c7d5-188-130-155-154.eu.ngrok.io/mainPage.html"
)

const (
	EditTransaction State = iota + 1
)

type UserStateType struct {
	CurrentRestaurant *restaurant.Restaurant
	CurrentOrder      []OrderPosition
	EditState         State
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
