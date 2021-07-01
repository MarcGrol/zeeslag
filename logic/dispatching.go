package logic

import (
	"github.com/MarcGrol/zeeslag/core"
)

type commandDispatcher struct {
	commandHandlers []commandGameState
}

type commandDispatcherCallback func(service *GameLogicService, game Game, command core.GameCommandPdu) ([]core.GameEventPdu, error)

type commandGameState struct {
	gameState   GameStatus
	commandType core.CommandType
	callback    commandDispatcherCallback
	nextState   GameStatus
}

func newCommandDispatcher(gameStates []commandGameState) *commandDispatcher {
	return &commandDispatcher{
		commandHandlers: gameStates,
	}
}

func (et commandDispatcher) resolveEvent(gameState GameStatus, commandType core.CommandType) (commandDispatcherCallback, GameStatus,bool) {
	for _, h := range et.commandHandlers {
		if h.gameState == gameState && h.commandType == commandType {
			return h.callback, h.nextState, true
		}
	}
	return nil, 0, false
}

type eventDispatcher struct {
	eventHandlers  []eventGameState
}

type eventDispatcherCallback func(service *GameLogicService, game Game, event core.GameEventPdu) ([]core.GameEventPdu,error)

type eventGameState struct {
	gameState GameStatus
	eventType core.EventType
	callback  eventDispatcherCallback
	nextState GameStatus
}

func newEventDispatcher(gameStates []eventGameState) *eventDispatcher {
	return &eventDispatcher{
		eventHandlers: gameStates,
	}
}

func (et eventDispatcher) resolveCommand(gameState GameStatus, eventType core.EventType) (eventDispatcherCallback, GameStatus,bool) {
	for _, h := range et.eventHandlers {
		if h.gameState == gameState && h.eventType == eventType {
			return h.callback, h.nextState, true
		}
	}
	return nil, 0, false
}

