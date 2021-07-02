package peerService

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
	"log"
)

var eventStateDispatching = []eventGameState{
	{
		description: "",
		gameState:   model.Idle,
		eventType:   core.EventType_InvitedForGame,
		callback:    onInvited,
		nextState:   model.Invited,
	}, {
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

func onInvited(s *PeerService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.InvitedForGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		log.Printf("invited:%+v", pdu)
		events = append(events, pdu)

		return events, nil
	}(*pdu.Invited)
}

func onInvitationAccepted(s *PeerService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameAccepted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Accepted)
}

func onInvitationRejected(s *PeerService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameRejected) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Rejected)
}

func onGameQuited(s *PeerService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameQuited) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Quited)
}

func onSalvoFired(s *PeerService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.SalvoFired) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Fired)
}

func onSalvoImpactAssessed(s *PeerService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.SalvoImpactAssessed) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Assessed)
}

func onGameCompleted(s *PeerService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameCompleted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Completed)
}
