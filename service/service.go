package service

import (
	"github.com/MarcGrol/zeeslag/cmd"
		"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

type ZeeslagService struct {
	id    string
	store core.GameEventStorer
}

func (s *ZeeslagService) OnStart(cmd cmd.StartGame) (*model.Game, error) {
	game, err := s.gameForEvents(cmd.GameId)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (s *ZeeslagService) OnAccept(cmd cmd.AcceptGame) (*model.Game, error) {
	game,err := s.gameForEvents(cmd.GameId)
	if err != nil {
		return nil, err
	}

	return game, nil
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
