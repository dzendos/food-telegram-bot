package queries

import (
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

var webUrl = "https://dce5-188-130-155-166.eu.ngrok.io"

func New(tgClient QueriesHandler, token string) *Model {
	mux := http.NewServeMux()
	mux.HandleFunc("/validate", validate(token))
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	return &Model{
		tgClient: tgClient,
		Server:   server,
	}
}

func validate(token string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%v", request)
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
