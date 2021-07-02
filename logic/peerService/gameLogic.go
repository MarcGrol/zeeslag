package peerService

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

var msgStateDispatching = []msgGameState{
	{
		description: "",
		gameState:   model.Idle,
		msgType:     core.ReplicatiomMsgType_PeerHasInvitedYouForGame,
		callback:    onInvited,
		nextState:   model.Invited,
	}, {
		description: "",
		gameState:   model.InvitationPending,
		msgType:     core.ReplicatiomMsgType_PeerHasAcceptedInvitation,
		callback:    onInvitationAccepted,
		nextState:   model.Active,
	},
	{
		description: "",
		gameState:   model.InvitationPending,
		msgType:     core.ReplicatiomMsgType_PeerHasRejectedInvitation,
		callback:    onInvitationRejected,
		nextState:   model.Rejected,
	},
	{
		description: "",
		gameState:   model.Active,
		msgType:     core.ReplicatiomMsgType_PeerHasQuited,
		callback:    onGameQuited,
		nextState:   model.Quited,
	},
	{
		description: "",
		gameState:   model.Active,
		msgType:     core.ReplicatiomMsgType_PeerHasFiredSalvo,
		callback:    onSalvoFired,
		nextState:   model.Active,
	},
	{
		description: "",
		gameState:   model.Active,
		msgType:     core.ReplicatiomMsgType_PeerHasAssessedImpactOfSalvo,
		callback:    onSalvoImpactAssessed,
		nextState:   model.Active,
	},
	{
		description: "",
		gameState:   model.Active,
		msgType:     core.ReplicatiomMsgType_PeerHasCompleted,
		callback:    onGameCompleted,
		nextState:   model.Completed,
	},
}

func onInvited(s *PeerService, game model.Game, pdu core.GameMsgPdu) ([]core.GameEventPdu, error) {
	return func(msg core.PeerHasInvitedYouForGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		recvd := core.GameInvitionReceived{
			GameId:    msg.GameId,
			Initiator: msg.Initiator,
			Invitee:   msg.Invitee,
		}
		events = append(events, recvd.ToPdu())

		return events, nil
	}(*pdu.Invited)
}

func onInvitationAccepted(s *PeerService, game model.Game, pdu core.GameMsgPdu) ([]core.GameEventPdu, error) {
	return func(msg core.PeerHasAcceptedGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// TODO

		return events, nil
	}(*pdu.Accepted)
}

func onInvitationRejected(s *PeerService, game model.Game, pdu core.GameMsgPdu) ([]core.GameEventPdu, error) {
	return func(msg core.PeerHasRejectedInvitation) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// TODO

		return events, nil
	}(*pdu.Rejected)
}

func onGameQuited(s *PeerService, game model.Game, pdu core.GameMsgPdu) ([]core.GameEventPdu, error) {
	return func(msg core.PeerHasQuited) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// TODO

		return events, nil
	}(*pdu.Quited)
}

func onSalvoFired(s *PeerService, game model.Game, pdu core.GameMsgPdu) ([]core.GameEventPdu, error) {
	return func(msg core.PeerHasFiredSalvo) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// TODO

		return events, nil
	}(*pdu.Fired)
}

func onSalvoImpactAssessed(s *PeerService, game model.Game, pdu core.GameMsgPdu) ([]core.GameEventPdu, error) {
	return func(msg core.PeerHasAssessedImpactOfSalvo) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// TODO

		return events, nil
	}(*pdu.Assessed)
}

func onGameCompleted(s *PeerService, game model.Game, pdu core.GameMsgPdu) ([]core.GameEventPdu, error) {
	return func(msg core.PeerHasCompleted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// TODO

		return events, nil
	}(*pdu.Completed)
}
