package main

import (
	"flag"
	"github.com/MarcGrol/zeeslag/infra"
	"github.com/MarcGrol/zeeslag/ui"

	"github.com/MarcGrol/zeeslag/logic"
)

var playerName string

func init() {
	flag.StringVar(&playerName, "playernamw", "me", "Player name to use for the game")
}

func main() {
	flag.Parse()

	pubsub := infra.NewBasicPubsub()

	// channels shareed by both service as userinterface
	channelsToSelf := infra.NewChannelsToSelf()

	{
		// Start own service in background
		coreLogic := logic.NewGameLogicService(logic.NewGameRepository(infra.NewBasicEventStore(), pubsub))
		playerService := infra.NewPlayerService(playerName, channelsToSelf, coreLogic)
		go playerService.Listen()
	}

	// Start service in foreground
	ui.NewUserInterface(channelsToSelf)
	ui.Listen()
}
