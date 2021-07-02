package userService

import (
	"fmt"
	"github.com/MarcGrol/zeeslag/model"
	"log"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/logic/repo"
)

type UserService struct {
	repo              *repo.GameRepository
	commandDispatcher *commandDispatcher
}

func NewUserService(repo *repo.GameRepository) *UserService {
	return &UserService{
		repo:              repo,
		commandDispatcher: newCommandDispatcher(commandStateDisppatching),
	}
}

func (s *UserService) OnQuery(gameId string) (*model.Game, bool, error) {
	return s.repo.GetGameOnId(gameId)
}

func (s *UserService) OnCommand(command core.GameCommandPdu) error {

	game, exist, err := s.repo.GetGameOnId(command.GameId)
	if err != nil {
		log.Printf("Error fetching game for command %+v: %+v", command, err)
		return err
	}

	if !exist {
		return fmt.Errorf("Game not found")
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
