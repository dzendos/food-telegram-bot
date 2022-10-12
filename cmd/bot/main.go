package main

import (
	"log"

	"github.com/dzendos/dubna/internal/clients/tg"
	"github.com/dzendos/dubna/internal/config"
	"github.com/dzendos/dubna/internal/model/callbacks"
	"github.com/dzendos/dubna/internal/model/messages"
	"github.com/dzendos/dubna/internal/model/queries"
	"github.com/dzendos/dubna/internal/model/state"
	"github.com/dzendos/dubna/scrapper"
)

func main() {
	scrapper.InitializeServerState()

	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	state.RestaurantReference = config.Url() + "restaurantPage.html"
	state.MenuReference = config.Url() + "mainPage.html"

	log.Println(state.RestaurantReference)

	msgModel := messages.New(tgClient)
	callbackModel := callbacks.New(tgClient)
	serverModel := queries.New(tgClient, config.Token())

	tgClient.ListenUpdates(msgModel, callbackModel)

	if err := serverModel.Server.ListenAndServe(); err != nil {
		panic("failed to listen and serve: " + err.Error())
	}
}
