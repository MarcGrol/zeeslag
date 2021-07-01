package logic

import (
	"fmt"
	"log"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
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

func (s *GameLogicService) OnEvent(event core.GameEventPdu) error {
	game, err := s.getGameOnId(event.GetId())
	if err != nil {
		log.Printf("Error fetching game for event %+v: %+v", event, err)
		return err
	}

	dispatchFunc, found := s.eventDispatcher.resolveCommand(game.Status, event.EventType)
	if !found {
		return fmt.Errorf("Event %+v could not be resolved for state %+v", event.EventType, game.Status)
	}

	return dispatchFunc(s, *game, event)
}

func (s *GameLogicService) OnCommand(command core.GameCommandPdu) error {
	game, err := s.getGameOnId(command.GetId())
	if err != nil {
		log.Printf("Error fetching game for command %+v: %+v", command, err)
		return err
	}

	dispatchFunc, found := s.commandDispatcher.resolveEvent(game.Status, command.CommandType)
	if !found {
		return fmt.Errorf("Command %+v could not be resolved for state %+v", command.CommandType, game.Status)
	}

	return dispatchFunc(s, *game, command)
}

// TODO Convert into a repository
func (s *GameLogicService) getGameOnId(gameId string) (*model.Game, error) {
	events, err := s.store.GetEventsOnGame(gameId)
	if err != nil {
		return nil, err
	}

	game := model.NewGame(events)

	return game, nil
}
