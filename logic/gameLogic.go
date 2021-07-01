package logic

import (
	"github.com/MarcGrol/zeeslag/core"
)

var commandStateDisppatching = []commandGameState{
	{
		gameState:   Idle,
		commandType: core.CommandType_InviteForGame,
		callback:    inviteForGame,
		nextState:   InvitationPending,
	},
	{
		gameState:   InvitationPending,
		commandType: core.CommandType_Quit,
		callback:    quitGame,
		nextState:   Quited,
	},
	{
		gameState:   InvitationPending,
		commandType: core.CommandType_Accept,
		callback:    acceptGame,
		nextState:   Active,
	},
	{
		gameState:   InvitationPending,
		commandType: core.CommandType_Reject,
		callback:    rejectGame,
		nextState:   Rejected,
	},
	{
		gameState:   Active,
		commandType: core.CommandType_Quit,
		callback:    quitGame,
		nextState:   Quited,
	},
	{
		gameState:   Active,
		commandType: core.CommandType_Fire,
		callback:    fireSalvo,
		nextState:   Active,
	},
}

var eventStateDispatching = []eventGameState{
	{
		gameState: Idle,
		eventType: core.EventType_GridPopulated,
		callback:  onGridPopulated	,
		nextState: Created,
	},
	{
		gameState: Created,
		eventType: core.EventType_InvitedForGame,
		callback:  onInvitedForGame,
		nextState: InvitationPending,
	},
	{
		gameState: InvitationPending,
		eventType: core.EventType_GameAccepted,
		callback:  onInvitationAccepted,
		nextState: Active,
	},
	{
		gameState: InvitationPending,
		eventType: core.EventType_GameRejected,
		callback:  onInvitationRejected,
		nextState: Rejected,
	},
	{
		gameState: Active,
		eventType: core.EventType_GameQuited,
		callback:  onGameQuited,
		nextState: Quited,
	},
	{
		gameState: Active,
		eventType: core.EventType_SalvoFired,
		callback:  onSalvoFired,
		nextState: Active,
	},
	{
		gameState: Active,
		eventType: core.EventType_SalvoImpactAssessed,
		callback:  onSalvoImpactAssessed,
		nextState: Active,
	},
	{
		gameState: Active,
		eventType: core.EventType_GameCompleted,
		callback:  onGameCompleted,
		nextState: Completed,
	},
}


func onGridPopulated(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GridPopulated) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{pdu}

		return events, nil
	}(*pdu.Populated)
}

func inviteForGame(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.InviteForGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// Validate: fields

		// Validate: allowed for current state

		// Compose and repo events

		return events, nil
	}(*pdu.Invite)
}

func acceptGame(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.AcceptGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// Validate: fields

		// Validate: allowed for current state

		// Compose and repo events

		return events, nil
	}(*pdu.Accept)
}

func rejectGame(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.RejectGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// Validate: fields

		// Validate: allowed for current state

		// Compose and repo events

		return events, nil
	}(*pdu.Reject)
}

func fireSalvo(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.Salvo) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// Validate: fields

		// Validate: allowed for current state

		// Compose and repo events

		return events, nil
	}(*pdu.Fire)
}

func quitGame(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.QuitGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// Validate: fields

		// Validate: allowed for current state

		// Compose and repo events

		return events, nil
	}(*pdu.Quit)
}

func onInvitedForGame(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.InvitedForGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Invited)
}

func onInvitationAccepted(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameAccepted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Accepted)
}

func onInvitationRejected(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameRejected) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Rejected)
}

func onGameQuited(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameQuited) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Quited)
}

func onSalvoFired(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.SalvoFired) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Fired)
}

func onSalvoImpactAssessed(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.SalvoImpactAssessed) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Assessed)
}

func onGameCompleted(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GameCompleted) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		return events, nil
	}(*pdu.Completed)
}
