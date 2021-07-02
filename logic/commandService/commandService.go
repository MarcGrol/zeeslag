package commandService

import (
	"fmt"
	repo2 "github.com/MarcGrol/zeeslag/logic/repo"
	"log"

	"github.com/MarcGrol/zeeslag/core"
)

type CommandService struct {
	repo              *repo2.GameRepository
	commandDispatcher *commandDispatcher
}

func NewCommandService(repo *repo2.GameRepository) *CommandService {
	return &CommandService{
		repo:              repo,
		commandDispatcher: newCommandDispatcher(commandStateDisppatching),
	}
}

func (s *CommandService) OnCommand(command core.GameCommandPdu) error {

	game, err := s.repo.GetGameOnId(command.GameId)
	if err != nil {
		log.Printf("Error fetching game for command %+v: %+v", command, err)
		return err
	}

	log.Printf("Got command %s (%+v) for game: %s 	(%+v)", command.CommandType, command, game.Status, game)

	// Lookup if this state-command can be handled
	dispatchFunc, expectedNextStatus, found := s.commandDispatcher.resolve(game.Status, command.CommandType)
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
		log.Printf("Unexpected next state %s for command: %+v, expected status: %s", game.Status, command, expectedNextStatus)
		return err
	}

	// Store resulting events
	err = s.repo.StoreEvents(events)
	if err != nil {
		log.Printf("Error storing event %+v: %s", events, err)
		return err
	}

	return nil
}
