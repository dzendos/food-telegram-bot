package queries

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/dzendos/dubna/internal/model/position"
	"github.com/dzendos/dubna/internal/model/state"
)

type QueriesHandler interface {
	SendRestaurantMenu(userID int64) error
}

type Model struct {
	tgClient QueriesHandler
	Server   http.Server
}

func New(tgClient QueriesHandler, token string) *Model {
	var s Model

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./web/")))
	mux.HandleFunc("/validate", validate(token))
	mux.HandleFunc("/getRestaurants", s.getRestaurant)
	mux.HandleFunc("/sendRestaurant", s.sendRestaurant)
	mux.HandleFunc("/getMenu", s.getMenu)
	mux.HandleFunc("/sendOrder", s.orderIsReady)
	mux.HandleFunc("/getOrder", s.getOrder)

	server := http.Server{
		Handler: mux,
		Addr:    "0.0.0.0:8080",
	}

	s.tgClient = tgClient
	s.Server = server

	return &s
}

func validate(token string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		ok, err := ext.ValidateWebAppQuery(request.URL.Query(), token)
		if err != nil {
			writer.Write([]byte("validation failed; error: " + err.Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if ok {
			writer.Write([]byte("validation success; user is authenticated."))
		} else {
			writer.Write([]byte("validation failed; data cannot be trusted."))
		}
	}
}

type restaurantQuery struct {
	Name string `json:"Name"`
	Url  string `json:"Url"`
}

type restaurantsQuery struct {
	Restaurants []restaurantQuery `json:"Restaurants"`
}

// Query getRestaurant -> (name, description, photo, methods of connections)
func (s *Model) getRestaurant(writer http.ResponseWriter, request *http.Request) {
	restaurants := restaurantsQuery{}

	for _, restaurant := range state.ServerState.Restaurants {
		restaurants.Restaurants = append(restaurants.Restaurants, restaurantQuery{
			Name: restaurant.Name,
			Url:  restaurant.ImageUrl,
		})
	}

	reqBody, err := json.Marshal(restaurants)

	if err != nil {
		log.Println(err, "queries.getRestaurant")
	}

	log.Println(reqBody)
	writer.Write(reqBody)
}

type getMenuQuery struct {
	UserID     int64  `json:"UserID"`
	Restaurant string `json:"Restaurant"`
}

func (s *Model) sendRestaurant(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
	}

	userID, restaurantName := encodeGetMenuQuery(body)

	state.SetUserRestaurant(userID, restaurantName)

	// Send user reference to the menu of the restaurant
	s.tgClient.SendRestaurantMenu(userID)
}

type returnMenuQuery struct {
	RestaurantName string               `json:"Restaurant"`
	Menu           []*position.Position `json:"Menu"`
}

// Query getMenu (userID, restaurant) -> (array of dishes(name, description, photo, maybe category))
func (s *Model) getMenu(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
	}

	userID := encodeUserID(body)
	restaurant := state.GetUserRestaurant(userID)

	if restaurant == nil {
		time.Sleep(3 * time.Millisecond)
		restaurant = state.GetUserRestaurant(userID)
	}

	var answer = returnMenuQuery{
		RestaurantName: restaurant.Name,
		Menu:           restaurant.Menu.Positions,
	}

	reqBody, err := json.Marshal(answer)

	if err != nil {
		log.Println(err, "queries.getMenu")
	}

	log.Println(string(reqBody))
	writer.Write(reqBody)
}

func encodeGetMenuQuery(body []byte) (int64, string) {
	restaurant := getMenuQuery{}

	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&restaurant)

	if err != nil {
		return -1, ""
	}

	return restaurant.UserID, restaurant.Restaurant
}

// Query orderIsReady -> restaurant -> (array of dishes(by id?))
func (s *Model) orderIsReady(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)

	if err != nil {
		log.Println(err)
	}

	order := encodeOrderIsReady(body)

	state.SetUserOrder(order)
}

func encodeOrderIsReady(body []byte) *state.Order {
	order := state.Order{}

	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&order)

	if err != nil {
		return nil
	}

	return &order
}

type RetOrderPosition struct {
	Name   string  `json:"PositionName"`
	Amount int     `json:"PositionAmount"`
	Price  float64 `json:"PositionPrice"`
}

type orderToReturn struct {
	Order []RetOrderPosition `json:"Order"`
}

func (s *Model) getOrder(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)

	if err != nil {
		log.Println(err)
	}

	userID := encodeUserID(body)

	usersOrder := state.GetUserOrder(userID)
	if usersOrder == nil {
		log.Println("this user does not have an order")
		return
	}

	order := orderToReturn{}
	for _, position := range usersOrder {
		order.Order = append(order.Order, RetOrderPosition{
			Name:   position.Name,
			Amount: position.Amount,
			Price:  state.GetPositionPrice(userID, position.Name),
		})
	}

	reqBody, err := json.Marshal(order)

	if err != nil {
		log.Println(err, "queries.getMenu")
	}

	log.Println(string(reqBody))
	writer.Write(reqBody)
}

type userIDType struct {
	UserID int64 `json:"UserID"`
}

func encodeUserID(body []byte) int64 {
	userID := userIDType{}

	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&userID)

	if err != nil {
		return -1
	}

	return userID.UserID
}
