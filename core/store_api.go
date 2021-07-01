package core

import (
	"github.com/MarcGrol/zeeslag/evt"
)

type GameEventStorer interface {
	GetEventsOnGame(gameId string) ([]evt.GameEventPdu, error)
	AddEventToGame(gameId string, event evt.GameEventPdu) error
}

