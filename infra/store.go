package infra

import (
	"github.com/MarcGrol/zeeslag/core"
)

type basicEventStore struct {
	events []core.GameEventPdu
}

func NewBasicEventStore() core.GameEventStorer {
	return &basicEventStore{
		events: []core.GameEventPdu{},
	}
}

func (s *basicEventStore) GetEventsOnGame(gameId string) ([]core.GameEventPdu, error) {
	return []core.GameEventPdu{}, nil
}

func (s *basicEventStore) AddEventToGame(gameId string, evt core.GameEventPdu) error {
	return nil
}
