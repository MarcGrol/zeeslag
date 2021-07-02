package userService

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

type commandDispatcher struct {
	commandHandlers []commandGameState
}

type commandDispatcherCallback func(service *UserService, game model.Game, command core.GameCommandPdu) ([]core.GameEventPdu, error)

type commandGameState struct {
	description string
	gameState   model.GameStatus
	commandType core.CommandType
	callback    commandDispatcherCallback
	nextState   model.GameStatus
}

func newCommandDispatcher(gameStates []commandGameState) *commandDispatcher {
	return &commandDispatcher{
		commandHandlers: gameStates,
	}
}

func (et commandDispatcher) resolve(gameState model.GameStatus, commandType core.CommandType) (commandDispatcherCallback, model.GameStatus, bool) {
	for _, h := range et.commandHandlers {
		if h.gameState == gameState && h.commandType == commandType {
			return h.callback, h.nextState, true
		}
	}
	return nil, 0, false
}
