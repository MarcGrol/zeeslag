package logic

import (
	"fmt"
	"log"

	"github.com/MarcGrol/zeeslag/core"
)

type GameLogicService struct {
	repo              *GameRepository
	commandDispatcher *commandDispatcher
	eventDispatcher   *eventDispatcher
}

func NewGameLogicService(repo *GameRepository) core.Service {
	return &GameLogicService{
		repo:              repo,
		commandDispatcher: newCommandDispatcher(commandStateDisppatching),
		eventDispatcher:   newEventDispatcher(eventStateDispatching),
	}
}

func (s *GameLogicService) OnCommand(command core.GameCommandPdu) error {

	game, err := s.repo.GetGameOnId(command.GameId)
	if err != nil {
		log.Printf("Error fetching game for command %+v: %+v", command, err)
		return err
	}

	log.Printf("Got command %s (%+v) for game: %s 	(%+v)", command.CommandType, command, game.Status, game)

	// Lookup if this state-command can be handled
	dispatchFunc, expectedNextStatus, found := s.commandDispatcher.resolveEvent(game.Status, command.CommandType)
	if !found {
		return fmt.Errorf("Command %+v could not be resolved for state %+v", command.CommandType, game.Status)
	}

	log.Printf("Handler found for status: %s and command: %s", game.Status, command.CommandType)

	// Call state-command specific logic
	events, err := dispatchFunc(s, *game, command)
	if err != nil {
		log.Printf("Error handling command %+v: %s", command, err)
		return err
	}

	// Safeguard that next state is set correctly
	game.ApplyAll(events)
	if game.Status != expectedNextStatus {
		log.Printf("Unexpected next state %s for command: %+v, expected status: %s", game.Status, command, 	expectedNextStatus)
		return err
	}

	// Store resulting events
	for _, e := range events {
		err = s.repo.StoreEvent(e)
		if err != nil {
			log.Printf("Error storing event %+v: %s", e, err)
			return err
		}
	}

	return nil
}

func (s *GameLogicService) OnEvent(event core.GameEventPdu) error {
	game, err := s.repo.GetGameOnId(event.GameId)
	if err != nil {
		log.Printf("Error fetching game for event %+v: %+v", event, err)
		return err
	}

	log.Printf("Got event %s (%+v) for game: %s (%+v)", event.EventType, event, game.Status, game)

	// Lookup if this state-event can be handled
	dispatchFunc, expectedNextStatus, found := s.eventDispatcher.resolveCommand(game.Status, event.EventType)
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
	for _, e := range events {
		err = s.repo.StoreEvent(e)
		if err != nil {
			log.Printf("Error storing event %+v: %s", e, err)
			return err
		}
	}


	return nil
}
