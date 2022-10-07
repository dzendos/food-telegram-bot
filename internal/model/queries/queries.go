package queries

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type QueriesHandler interface {
}

type Model struct {
	tgClient QueriesHandler
	Server   http.Server
}

func New(tgClient QueriesHandler, token string) *Model {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./web/")))
	mux.HandleFunc("/validate", validate(token))
	mux.HandleFunc("/getRestaurant", getRestaurant)
	mux.HandleFunc("/getMenu", getMenu)

	server := http.Server{
		Handler: mux,
		Addr:    "0.0.0.0:8080",
	}

	return &Model{
		tgClient: tgClient,
		Server:   server,
	}
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

// Query getRestaurant -> (name, description, photo, methods of connections)
func getRestaurant(writer http.ResponseWriter, request *http.Request) {
	restaurants := []string{"Dodo", "MakDak", "KFC", "Мишлен"}

	reqBody, err := json.Marshal(map[string][]string{
		"Restaurants": restaurants,
	})

	if err != nil {
		log.Println(err, "queries.getRestaurant")
	}

	log.Println(reqBody)
	writer.Write(reqBody)
}

type dish struct {
	name      string
	price     float64
	dish_type string
}

func getMenu(writer http.ResponseWriter, request *http.Request) {
	a := `[{"name":"хархок","price":500.5,"type":"первое"},{"name":"смузи","price":180.5,"type":"напиток"},{"name":"чай","price":12.5,"type":"напиток"},{"name":"борщ","price":210.5,"type":"первое"}]`
	writer.Write([]byte(a))
}

// Query getMenu restaurant -> (array of dishes(name, description, photo, maybe category))

// Query orderIsReady -> restaurant -> (array of dishes(by id?))
