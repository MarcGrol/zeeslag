package main

import (
	"flag"
	"github.com/MarcGrol/zeeslag/infra"

	"github.com/google/uuid"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/logic"
)

var playerName string

func init() {
	flag.StringVar(&playerName, "playernamw", "me", "Player name to use for the game")
}

func main() {
	flag.Parse()

	// shared channels:
	channelsToSelf := infra.NewChannelsToSelf()

	channelsToOther := infra.NewChannelsToOther()

	evtStore := infra.NewBasicEventStore()
	coreLogic := logic.NewGameService(evtStore)

	playerService := infra.NewPlayerService(playerName, channelsToSelf, channelsToOther, coreLogic)
	playerService.ListenInBackground()

	gameId := uuid.New().String()

	playerService.Command(
		core.InviteForGame{GameId: gameId}.ToPdu())

	// waitfor ^C or command-line actions
	// readKeyboardInput
	// convert keyboard input into command
	// push command into service

}
