package service

import (
	"github.com/MarcGrol/zeeslag/cmd"
		"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
	"log"
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

type ZeeslagService struct {
	store core.GameEventStorer
}

func NewZeeslagService(store core.GameEventStorer) cmd.CommandService {
	return &ZeeslagService{
		store:store,
	}
}

func (s *ZeeslagService) OnStart(cmd cmd.StartGame) error {
	game, err := s.gameForEvents(cmd.GameId)
	if err != nil {
		return err
	}

	log.Printf("game: %+v",game)

	// Validate: fields

	// Validate: allowed for current state

	// Compose and store events

	return nil
}

func (s *ZeeslagService) OnAccept(cmd cmd.AcceptGame) error {
	game,err := s.gameForEvents(cmd.GameId)
	if err != nil {
		return err
	}

	log.Printf("game: %+v",game)

	// Validate: fields

	// Validate: allowed for current state

	// Compose and store events

	return nil
}

func (s *ZeeslagService) OnFire(cmd cmd.Salvo) error {
	game, err := s.gameForEvents(cmd.GameId)
	if err != nil {
		return err
	}

	log.Printf("game: %+v",game)

	// Validate: fields

	// Validate: allowed for current state

	// Compose and store events

	return nil
}

func (s *ZeeslagService)gameForEvents(gameId string) (*model.Game, error) {
	events, err := s.store.GetEventsOnGame(gameId)
	if err != nil {
		return nil, err
	}

	game := model.NewGame(events)

	return game, nil
}

//
//
/////secondPlayerGame := NewGame()
//
//firstPlayerGame.ApplyGameInitiated(GameInitiated{
//	GameId:    gameId,
//	Initiator: "marc",
//	Invitee:   "Eva",
//})
//
//firstPlayerGame.ApplyGameAccepted(GameAccepted{
//	GameId: gameId,
//})
//
//firstPlayerGame.ApplySalvoFired(SalvoFired{
//	GameId: gameId,
//})
//}
