package logic

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

type eventDispatcher struct {
	eventHandlers  []eventGameState
}

type eventDispatcherCallback func(service *GameLogicService, game model.Game, event core.GameEventPdu) error

type eventGameState struct {
	gameState model.GameStatus
	eventType core.EventType
	callback  eventDispatcherCallback
	nextState model.GameStatus
}

func newEventDispatcher(gameStates []eventGameState) *eventDispatcher {
	return &eventDispatcher{
		eventHandlers: gameStates,
	}
}

func (et eventDispatcher) resolveCommand(gameState model.GameStatus, eventType core.EventType) (eventDispatcherCallback, bool) {
	for _, h := range et.eventHandlers {
		if h.gameState == gameState && h.eventType == eventType {
			return h.callback, true
		}
	}
	return nil, false
}

