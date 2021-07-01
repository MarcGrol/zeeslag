package logic

import "github.com/MarcGrol/zeeslag/core"

type GameEventStorer interface {
	GetEventsOnGame(gameId string) ([]core.GameEventPdu, error)
	AddEventToGame(gameId string, event core.GameEventPdu) error
}
