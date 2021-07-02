package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MarcGrol/zeeslag/infra"
	"github.com/MarcGrol/zeeslag/logic/commandService"
	"github.com/MarcGrol/zeeslag/logic/eventService"
	"github.com/MarcGrol/zeeslag/logic/repo"
)

var playerName string

func init() {
	flag.StringVar(&playerName, "playernamw", "me", "Player name to use for the game")
}

func main() {
	flag.Parse()

	pubsub := infra.NewBasicPubsub()
	repo := repo.NewGameRepository(infra.NewBasicEventStore(), pubsub)

	router := mux.NewRouter()

	eventService.RegisterHTTPEndpoint(router, eventService.NewEventService(repo))
	commandService.RegisterHTTPEndpoint(router, commandService.NewCommandService(repo))

	// Start listening for http requests in foreground
	http.ListenAndServe(":8080", router)
}
