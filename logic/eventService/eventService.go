package eventService

import (
	"fmt"
	"github.com/MarcGrol/zeeslag/api"
	"github.com/MarcGrol/zeeslag/model"
	"log"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/logic/repo"
)

type EventService struct {
	repo            *repo.GameRepository
	eventDispatcher *eventDispatcher
	peerer          api.Peerer
	pubsub          api.PubSub
}

func NewEventService(repo *repo.GameRepository, pubsub api.PubSub, peerer api.Peerer) *EventService {

	es := &EventService{
		repo:            repo,
		eventDispatcher: newEventDispatcher(eventStateDispatching),
		peerer:          peerer,
	}

	pubsub.Subscribe("game", es)

	return es
}

func (eh *EventService) OnEventPublished(topic string, event core.GameEventPdu) error {
	return eh.OnEvent(event)
}

func (s *EventService) OnEvent(event core.GameEventPdu) error {
	game, exists, err := s.repo.GetGameOnId(event.GameId)
	if err != nil {
		log.Printf("Error fetching game for event %+v: %+v", event, err)
		return err
	}

	if !exists {
		game = &model.Game{}
	}

	log.Printf("Got event %s (%+v) for game: %s (%+v)", event.EventType, event, game.Status, game)

	// Lookup if this state-event can be handled
	dispatchFunc, expectedNextStatus, found := s.eventDispatcher.resolve(game.Status, event.EventType)
	if !found {
		return fmt.Errorf("Event %+v could not be resolved for state %s", event.EventType, game.Status)
	}

	// Call state-event specific logic
	events, err := dispatchFunc(s, *game, event)
	if err != nil {
		log.Printf("Error handling event %+v: %s", event, err)
		return err
	}

	// Safeguard that next state is set correctly
	game.ApplyAll(events)
	if game.Status != expectedNextStatus {
		log.Printf("Unexpected next state %v for event: %+v, expected status: %+v", game.Status, game, expectedNextStatus)
		return err
	}

	// Store resulting events
	err = s.repo.StoreEvents(events)
	if err != nil {
		log.Printf("Error storing events %+v: %s", events, err)
		return err
	}

	return nil
}
