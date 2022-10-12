package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/dzendos/dubna/internal/clients/tg"
	"github.com/dzendos/dubna/internal/config"
	"github.com/dzendos/dubna/internal/model/callbacks"
	"github.com/dzendos/dubna/internal/model/messages"
	"github.com/dzendos/dubna/internal/model/queries"
	"github.com/dzendos/dubna/internal/model/state"
	"github.com/dzendos/dubna/scrapper"
)

func pasteUrl(path string, outputPath string, replaceMent string, url string) {
	input, err := os.ReadFile(path)

	if err != nil {
		log.Fatal("pasting url to web failed: ", err)
	}

	output := bytes.Replace(input, []byte(replaceMent), []byte(url), 1)
	if err = os.WriteFile(outputPath, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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

	state.RestaurantReference = config.Url()
	state.MenuReference = config.Url() + "/mainPage.html"
	//{{.url}}
	pasteUrl("web/mainPage_template.js", "web/mainPage.js", "{{.url}}", config.Url())
	pasteUrl("web/restaurantPage_template.js", "web/restaurantPage.js", "{{.url}}", config.Url())

	log.Println(state.RestaurantReference)

	msgModel := messages.New(tgClient)
	callbackModel := callbacks.New(tgClient)
	serverModel := queries.New(tgClient, config.Token())

	tgClient.ListenUpdates(msgModel, callbackModel)

	if err := serverModel.Server.ListenAndServe(); err != nil {
		panic("failed to listen and serve: " + err.Error())
	}
}
