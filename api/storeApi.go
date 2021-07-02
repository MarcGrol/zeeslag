package api

import "github.com/MarcGrol/zeeslag/core"

type GameEventStorer interface {
	GetEventsOnGame(gameId string) ([]core.GameEventPdu, error)
	AddEventsToGame(event []core.GameEventPdu) error
}
