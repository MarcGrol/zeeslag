package main

import (
	"flag"

	"github.com/google/uuid"

	"github.com/MarcGrol/zeeslag/service"
	"github.com/MarcGrol/zeeslag/cmd"

)

//       A                            B
// ------------------startup by A------------------------
// gameInitiated
// gridPopulated
//                                gameAccepted
//                                gridPopulated
// gameAccepted
//
// ------------------salvo by A--------------------------
//
// saloFired
//                                salvoReceived
// salvoImpactProcessed
//
// ------------------salvo by B--------------------------
//
//                                salvoFired
// salvoReceived
//                                salvoImpactProcessed
//
// ------------------salvo by A--------------------------
// ...
// ------------------salvo by B--------------------------
// ...
// ------------------salvo by A--------------------------
// ...
// ------------------salvo by B--------------------------
// ...
// ------------------B has won---------------------------
//

// commands                   A: events stored       B: events stored
// A -> B: initiateGame       gameInitiated          gameAccepted
// B -> A: acceptGame         gameAccepted
//                            gridPopulated          grodPopulated
//
// A -> B: fireSalvo          salvoFired             salvoReceived
// B -> A: reportSalvoImpact  salvoImpactProcessed   salvoImpactProccessed
//
// B -> A: fireSalvo          salvoReceived          salvoFired
// A -> B: reportSalvoImpact                         salvoImpactProcessed
//
// ...
// ...
// A -> B: markCompleted      gameCompleted          gameCompleted

var playerName string
func init() {
	flag.StringVar(&playerName, "playernamw", "me", "Player name to use for the game")
}

func main() {
	flag.Parse()

	playerService := service.NewPlayerService(playerName)
	playerService.ListenInBackground()

	gameId := uuid.New().String()

	playerService.Command(
		cmd.StartGame{GameId:gameId}.ToPdu())

	// waitfor ^C
	// readKeyboardInput
	// convert keyboard input into command


}
