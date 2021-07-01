package logic

import (
	"github.com/MarcGrol/zeeslag/api"
	"github.com/MarcGrol/zeeslag/core"
)

type GameRepository struct {
	store api.GameEventStorer
}

func NewGameRepository(store api.GameEventStorer) *GameRepository {
	return &GameRepository{
		store: store,
	}
}

func (s *GameRepository) StoreEvent(event core.GameEventPdu) error {
	return s.store.AddEventToGame(event)
}

func (s *GameRepository) GetGameOnId(gameId string) (*Game, error) {
	events, err := s.store.GetEventsOnGame(gameId)
	if err != nil {
		return nil, err
	}

	game := NewGame(events)

	return game, nil
}
