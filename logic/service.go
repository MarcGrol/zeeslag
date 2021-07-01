package logic

import (
	"fmt"
	"log"

	"github.com/MarcGrol/zeeslag/core"
)

type GameLogicService struct {
	store             GameEventStorer
	commandDispatcher *commandDispatcher
	eventDispatcher   *eventDispatcher
}

func NewGameService(store GameEventStorer) core.Service {
	return &GameLogicService{
		store:             store,
		commandDispatcher: newCommandDispatcher(commandStateDisppatching),
		eventDispatcher:   newEventDispatcher(eventStateDispatching),
	}
}

func (s *GameLogicService) OnCommand(command core.GameCommandPdu) error {

	game, err := s.getGameOnId(command.GameId)
	if err != nil {
		log.Printf("Error fetching game for command %+v: %+v", command, err)
		return err
	}

	log.Printf("Got command %+v for game: %+v", command, game)

	dispatchFunc, expectedNextStatus, found := s.commandDispatcher.resolveEvent(game.Status, command.CommandType)
	if !found {
		return fmt.Errorf("Command %+v could not be resolved for state %+v", command.CommandType, game.Status)
	}

	events, err := dispatchFunc(s, *game, command)
	if err != nil {
		log.Printf("Error handling command %+v: %+v", command, err)
		return err
	}

	game.ApplyAll(events)

	if game.Status != expectedNextStatus {
		log.Printf("Unexpected next state for command: %+v, expected status: %+v", game.Status, expectedNextStatus)
		return err
	}

	return nil
}

func (s *GameLogicService) OnEvent(event core.GameEventPdu) error {
	game, err := s.getGameOnId(event.GameId)
	if err != nil {
		log.Printf("Error fetching game for event %+v: %+v", event, err)
		return err
	}

	log.Printf("Got event %+v for game: %+v", event, game)

	dispatchFunc, expectedNextStatus, found := s.eventDispatcher.resolveCommand(game.Status, event.EventType)
	if !found {
		return fmt.Errorf("Event %+v could not be resolved for state %+v", event.EventType, game.Status)
	}

	events, err := dispatchFunc(s, *game, event)
	if err != nil {
		log.Printf("Error handling event %+v: %+v", event, err)
		return err
	}

	game.ApplyAll(events)

	if game.Status != expectedNextStatus {
		log.Printf("Unexpected next state for event: %+v, expected status: %+v", game, expectedNextStatus)
		return err
	}

	return nil
}

// TODO Convert into a repository
func (s *GameLogicService) getGameOnId(gameId string) (*Game, error) {
	events, err := s.store.GetEventsOnGame(gameId)
	if err != nil {
		return nil, err
	}

	game := NewGame(events)

	return game, nil
}
