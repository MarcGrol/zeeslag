package infra

import (
	"github.com/MarcGrol/zeeslag/api"
	"github.com/MarcGrol/zeeslag/core"
	"sync"
)

type basicEventStore struct {
	mutex  sync.Mutex
	events []core.GameEventPdu
}

func NewBasicEventStore() api.GameEventStorer {
	return &basicEventStore{
		events: []core.GameEventPdu{},
	}
}

func (s *basicEventStore) GetEventsOnGame(gameId string) ([]core.GameEventPdu, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	found := []core.GameEventPdu{}

	for _, e := range s.events {
		if e.GameId == gameId {
			found = append(found, e)
		}
	}

	return found, nil
}

func (s *basicEventStore) AddEventsToGame(events []core.GameEventPdu) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, e := range events {
		s.events = append(s.events, e)
	}
	return nil
}
