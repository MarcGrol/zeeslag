package api

import "github.com/MarcGrol/zeeslag/core"

type GameEventStorer interface {
	GetEventsOnGame(gameId string) ([]core.GameEventPdu, error)
	AddEventToGame(event core.GameEventPdu) error
}
