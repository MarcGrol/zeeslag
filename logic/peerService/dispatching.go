package peerService

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

type msgDispatcher struct {
	msgHandlers []msgGameState
}

type msgDispatcherCallback func(service *PeerService, game model.Game, msg core.GameMsgPdu) ([]core.GameEventPdu, error)

type msgGameState struct {
	description string
	gameState   model.GameStatus
	msgType     core.ReplicatiomMsgType
	callback    msgDispatcherCallback
	nextState   model.GameStatus
}

func newMsgDispatcher(gameStates []msgGameState) *msgDispatcher {
	return &msgDispatcher{
		msgHandlers: gameStates,
	}
}

func (et msgDispatcher) resolve(gameState model.GameStatus, eventType core.ReplicatiomMsgType) (msgDispatcherCallback, model.GameStatus, bool) {
	for _, h := range et.msgHandlers {
		if h.gameState == gameState && h.msgType == eventType {
			return h.callback, h.nextState, true
		}
	}
	return nil, 0, false
}
