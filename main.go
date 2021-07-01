package main

import (
	"flag"
	"github.com/MarcGrol/zeeslag/store"

	"github.com/google/uuid"

	"github.com/MarcGrol/zeeslag/cmd"
	"github.com/MarcGrol/zeeslag/service"
)



var playerName string
func init() {
	flag.StringVar(&playerName, "playernamw", "me", "Player name to use for the game")
}

func main() {
	flag.Parse()

	evtStore := store.NewEventStore()
	coreService := service.NewZeeslagService(evtStore)

	playerService := service.NewPlayerService(playerName, coreService)
	playerService.ListenInBackground()

	gameId := uuid.New().String()

	playerService.Command(
		cmd.StartGame{GameId:gameId}.ToPdu())

	// waitfor ^C
	// readKeyboardInput
	// convert keyboard input into command


}
