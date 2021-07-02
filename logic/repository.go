package logic

import (
	"github.com/MarcGrol/zeeslag/api"
	"github.com/MarcGrol/zeeslag/core"
)

type GameRepository struct {
	store     api.GameEventStorer
	publisher api.Publisher
}

func NewGameRepository(store api.GameEventStorer, publisher api.Publisher) *GameRepository {
	return &GameRepository{
		store:     store,
		publisher: publisher,
	}
}

func (s *GameRepository) StoreEvents(events []core.GameEventPdu) error {
	return s.store.AddEventsToGame(events)
}

func (s *GameRepository) GetGameOnId(gameId string) (*Game, error) {
	events, err := s.store.GetEventsOnGame(gameId)
	if err != nil {
		return nil, err
	}

	game := NewGame(events)

	return game, nil
}
