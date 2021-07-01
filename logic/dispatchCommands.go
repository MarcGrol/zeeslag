package logic

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

type commandDispatcher struct {
	commandHandlers []commandGameState
}

type commandDispatcherCallback func(service *GameLogicService, game model.Game, command core.GameCommandPdu) error

type commandGameState struct {
	gameState   model.GameStatus
	commandType core.CommandType
	callback    commandDispatcherCallback
	nextState model.GameStatus
}


func newCommandDispatcher(gameStates []commandGameState) *commandDispatcher {
	return &commandDispatcher{
		commandHandlers: gameStates,
	}
}

func (et commandDispatcher) resolveEvent(gameState model.GameStatus, commandType core.CommandType) (commandDispatcherCallback, bool) {
	for _, h := range et.commandHandlers {
		if h.gameState == gameState && h.commandType == commandType {
			return h.callback, true
		}
	}
	return nil, false
}

