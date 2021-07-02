package eventService

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

var eventStateDispatching = []eventGameState{
	{
		description: "",
		gameState:   model.Idle,
		eventType:   core.EventType_GridPopulated,
		callback:    onGridPopulated,
		nextState:   model.Created,
	},
	{
		description: "",
		gameState:   model.Created,
		eventType:   core.EventType_InvitedForGame,
		callback:    onInvitedForGame,
		nextState:   model.InvitationPending,
	},
	{
		description: "",
		gameState:   model.InvitationPending,
		eventType:   core.EventType_GameAccepted,
		callback:    onInvitationAccepted,
		nextState:   model.Active,
	},
	{
		description: "",
		gameState:   model.InvitationPending,
		eventType:   core.EventType_GameRejected,
		callback:    onInvitationRejected,
		nextState:   model.Rejected,
	},
	{
		description: "",
		gameState:   model.Active,
		eventType:   core.EventType_GameQuited,
		callback:    onGameQuited,
		nextState:   model.Quited,
	},
	{
		description: "",
		gameState:   model.Active,
		eventType:   core.EventType_SalvoFired,
		callback:    onSalvoFired,
		nextState:   model.Active,
	},
	{
		description: "",
		gameState:   model.Active,
		eventType:   core.EventType_SalvoImpactAssessed,
		callback:    onSalvoImpactAssessed,
		nextState:   model.Active,
	},
	{
		description: "",
		gameState:   model.Active,
		eventType:   core.EventType_GameCompleted,
		callback:    onGameCompleted,
		nextState:   model.Completed,
	},
}

func onGridPopulated(s *EventService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GridPopulated) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{pdu}

		return events, nil
	}(*pdu.Populated)
}

func onInvitedForGame(s *EventService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.InvitedForGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Invited)
}

func onInvitationAccepted(s *EventService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameAccepted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Accepted)
}

func onInvitationRejected(s *EventService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameRejected) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Rejected)
}

func onGameQuited(s *EventService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameQuited) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Quited)
}

func onSalvoFired(s *EventService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.SalvoFired) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Fired)
}

func onSalvoImpactAssessed(s *EventService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.SalvoImpactAssessed) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Assessed)
}

func onGameCompleted(s *EventService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameCompleted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Completed)
}
