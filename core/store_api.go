package core

type GameEventStorer interface {
	GetEventsOnGame(gameId string) ([]GameEventPdu, error)
	AddEventToGame(gameId string, event GameEventPdu) error
}
