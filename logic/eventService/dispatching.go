package eventService

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

type eventDispatcher struct {
	eventHandlers []eventGameState
}

type eventDispatcherCallback func(service *EventService, game model.Game, event core.GameEventPdu) ([]core.GameEventPdu, error)

type eventGameState struct {
	description string
	gameState   model.GameStatus
	eventType   core.EventType
	callback    eventDispatcherCallback
	nextState   model.GameStatus
}

func newEventDispatcher(gameStates []eventGameState) *eventDispatcher {
	return &eventDispatcher{
		eventHandlers: gameStates,
	}
}

func (et eventDispatcher) resolve(gameState model.GameStatus, eventType core.EventType) (eventDispatcherCallback, model.GameStatus, bool) {
	for _, h := range et.eventHandlers {
		if h.gameState == gameState && h.eventType == eventType {
			return h.callback, h.nextState, true
		}
	}
	return nil, 0, false
}
