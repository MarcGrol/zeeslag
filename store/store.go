package store

import (
		"github.com/MarcGrol/zeeslag/evt"
	"github.com/MarcGrol/zeeslag/core"
)

type EventStore struct {
	events []evt.GameEventPdu
}

func NewEventStore() core.GameEventStorer {
	return &EventStore{
		events: []evt.GameEventPdu{},
	}
}

func (s *EventStore) GetEventsOnGame(gameId string) ([]evt.GameEventPdu, error) {
	return []evt.GameEventPdu{}, nil
}

func (s *EventStore) AddEventToGame(gameId string, evt evt.GameEventPdu) error {
	return nil
}

