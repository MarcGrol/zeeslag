package repo

import (
	"github.com/MarcGrol/zeeslag/api"
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
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
	err := s.store.AddEventsToGame(events)
	if err != nil {
		return err
	}

	for _, e := range events {
		err := s.publisher.Publish("game", e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *GameRepository) GetGameOnId(gameId string) (*model.Game, error) {
	events, err := s.store.GetEventsOnGame(gameId)
	if err != nil {
		return nil, err
	}

	game := model.NewGame(events)

	return game, nil
}
