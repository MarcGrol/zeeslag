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
		eventType:   core.EventType_GameInvitationReceived,
		callback:    onInvitedForGame,
		nextState:   model.InvitationPending,
	},
	{
		description: "",
		gameState:   model.InvitationPending,
		eventType:   core.EventType_GameInvitationAccepted,
		callback:    onInvitationAccepted,
		nextState:   model.Active,
	},
	{
		description: "",
		gameState:   model.InvitationPending,
		eventType:   core.EventType_GameInvitationRejected,
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

func onGridPopulated(s *UserService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GridPopulated) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{pdu}

		return events, nil
	}(*pdu.Populated)
}

func onInvitedForGame(s *UserService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameInvitionReceived) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// inform peer that he is invited
		s.peerer.InformPeer(pdu)

		return events, nil
	}(*pdu.Invited)
}

func onInvitationAccepted(s *UserService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameAccepted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// inform initiator that invitation is accepted
		s.peerer.InformPeer(pdu)

		return events, nil
	}(*pdu.Accepted)
}

func onInvitationRejected(s *UserService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameRejected) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// inform initiator that invitation is rejected
		s.peerer.InformPeer(pdu)

		return events, nil
	}(*pdu.Rejected)
}

func onGameQuited(s *UserService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameQuited) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// inform peer that other side has quit
		s.peerer.InformPeer(pdu)

		return events, nil
	}(*pdu.Quited)
}

func onSalvoFired(s *UserService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.SalvoFired) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// inform peer has fired solvo
		s.peerer.InformPeer(pdu)

		return events, nil
	}(*pdu.Fired)
}

func onSalvoImpactAssessed(s *UserService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.SalvoImpactAssessed) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// inform peer has impact of solvo has been assessment
		s.peerer.InformPeer(pdu)

		return events, nil
	}(*pdu.Assessed)
}

func onGameCompleted(s *UserService, game model.Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameCompleted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// inform peer that he has won
		s.peerer.InformPeer(pdu)

		return events, nil
	}(*pdu.Completed)
}
