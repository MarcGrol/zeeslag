package peerService

import (
	"fmt"
	"log"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/logic/repo"
)

type PeerService struct {
	repo            *repo.GameRepository
	eventDispatcher *msgDispatcher
}

func NewPeerService(repo *repo.GameRepository) *PeerService {
	return &PeerService{
		repo:            repo,
		eventDispatcher: newMsgDispatcher(msgStateDispatching),
	}
}

func (s *PeerService) OnPeerEvent(event core.GameMsgPdu) error {
	game, err := s.repo.GetGameOnId(event.GameId)
	if err != nil {
		log.Printf("Error fetching game for event %+v: %+v", event, err)
		return err
	}

	log.Printf("Got event %s (%+v) for game: %s (%+v)", event.MsgType, event, game.Status, game)

	// Lookup if this state-event can be handled
	dispatchFunc, expectedNextStatus, found := s.eventDispatcher.resolve(game.Status, event.MsgType)
	if !found {
		return fmt.Errorf("Event %+v could not be resolved for state %s", event.MsgType, game.Status)
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
