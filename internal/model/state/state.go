// Package state contains FSA for all users.
package state

import (
	"github.com/dzendos/dubna/internal/model/restaurant"
	"github.com/pkg/errors"
)

var userState = make(map[int64]UserStateType)

type UserStateType struct {
	CurrentRestaurant *restaurant.Restaurant
	CurrentOrder      []OrderPosition
}

type OrderPosition struct {
	Name   string `json:"PositionName"`
	Amount int    `json:"PositionAmount"`
}

type Order struct {
	UserID int64           `json:"UserID"`
	Order  []OrderPosition `json:"Order"`
}

func SetUserRestaurant(userID int64, restaurantName string) error {
	restaurant, err := GetRestaurantByName(restaurantName)
	if err != nil {
		return errors.Wrap(err, "state.SetUserRestaurant")
	}

	state, ok := userState[userID]

	if !ok {
		state = UserStateType{}
	}

	state.CurrentRestaurant = restaurant

	return nil
}

func SetUserOrder(order *Order) {
	state, ok := userState[order.UserID]

	if !ok {
		state = UserStateType{}
	}

	state.CurrentOrder = order.Order
}

func GetRestaurantByName(restaurantName string) (*restaurant.Restaurant, error) {
	for _, restaurant := range serverState.Restaurants {
		if restaurant.Name == restaurantName {
			return restaurant, nil
		}
	}

	return nil, errors.New("restaurant was not found")
}

func GetUserOrder(userID int64) []OrderPosition {
	return userState[userID].CurrentOrder
}

func GetPositionPrice(userID int64, positionName string) float64 {
	restaurant := userState[userID].CurrentRestaurant
	for _, position := range restaurant.Menu.Positions {
		if position.Name == positionName {
			return position.Price
		}
	}

	return -1
}

var serverState = ServerStateType{}

type ServerStateType struct {
	Restaurants []*restaurant.Restaurant
}
