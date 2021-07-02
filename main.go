package main

import (
	"flag"
	"github.com/MarcGrol/zeeslag/logic/peerService"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MarcGrol/zeeslag/infra"
	"github.com/MarcGrol/zeeslag/logic/eventService"
	"github.com/MarcGrol/zeeslag/logic/repo"
	"github.com/MarcGrol/zeeslag/logic/userService"
)

var (
	localPort     string
	remotePeerUrl string
	playerName    string
)

func init() {
	flag.StringVar(&localPort, "localPort", ":8080", "Local listen port")
	flag.StringVar(&remotePeerUrl, "remotePeerUrl", "http://localhost:8081", "Hostname of remote peer")
	flag.StringVar(&playerName, "playerName", "me", "Player name to use for the game")
}

func main() {
	flag.Parse()
	flag.PrintDefaults()

	pubsub := infra.NewBasicPubsub()
	repo := repo.NewGameRepository(infra.NewBasicEventStore(), pubsub)
	peer := infra.NewBasicPeer(remotePeerUrl)

	router := mux.NewRouter()

	eventService.RegisterHTTPEndpoint(router, eventService.NewEventService(repo, peer))
	peerService.RegisterHTTPEndpoint(router, peerService.NewPeerService(repo))
	userService.RegisterHTTPEndpoint(router, userService.NewUserService(repo))

	// Start listening for http requests in foreground
	log.Fatal(http.ListenAndServe(localPort, router))
}
