package infra

import (
	"github.com/MarcGrol/zeeslag/api"
	"github.com/MarcGrol/zeeslag/core"
)

type basicEventStore struct {
	events []core.GameEventPdu
}

func NewBasicEventStore() api.GameEventStorer {
	return &basicEventStore{
		events: []core.GameEventPdu{},
	}
}

func (s *basicEventStore) GetEventsOnGame(gameId string) ([]core.GameEventPdu, error) {
	found := []core.GameEventPdu{}

	for _, e := range s.events {
		if e.GameId == gameId {
			found = append(found, e)
		}
	}

	return found, nil
}

func (s *basicEventStore) AddEventToGame(evt core.GameEventPdu) error {
	s.events = append(s.events, evt)
	return nil
}
