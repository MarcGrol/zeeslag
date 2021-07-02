package logic

import (
	"fmt"

	"github.com/MarcGrol/zeeslag/core"
)

var commandStateDisppatching = []commandGameState{
	{
		description: "",
		gameState:   Idle,
		commandType: core.CommandType_InviteForGame,
		callback:    inviteForGame,
		nextState:   InvitationPending,
	},
	{
		description: "",
		gameState:   InvitationPending,
		commandType: core.CommandType_Quit,
		callback:    quitGame,
		nextState:   Quited,
	},
	{
		description: "",
		gameState:   InvitationPending,
		commandType: core.CommandType_Accept,
		callback:    acceptGame,
		nextState:   Active,
	},
	{
		description: "",
		gameState:   InvitationPending,
		commandType: core.CommandType_Reject,
		callback:    rejectGame,
		nextState:   Rejected,
	},
	{
		description: "",
		gameState:   Active,
		commandType: core.CommandType_Quit,
		callback:    quitGame,
		nextState:   Quited,
	},
	{
		description: "",
		gameState:   Active,
		commandType: core.CommandType_Fire,
		callback:    fireSalvo,
		nextState:   WaitforAssessment,
	},
}

var eventStateDispatching = []eventGameState{
	{
		description: "",
		gameState:   Idle,
		eventType:   core.EventType_GridPopulated,
		callback:    onGridPopulated,
		nextState:   Created,
	},
	{
		description: "",
		gameState:   Created,
		eventType:   core.EventType_InvitedForGame,
		callback:    onInvitedForGame,
		nextState:   InvitationPending,
	},
	{
		description: "",
		gameState:   InvitationPending,
		eventType:   core.EventType_GameAccepted,
		callback:    onInvitationAccepted,
		nextState:   Active,
	},
	{
		description: "",
		gameState:   InvitationPending,
		eventType:   core.EventType_GameRejected,
		callback:    onInvitationRejected,
		nextState:   Rejected,
	},
	{
		description: "",
		gameState:   Active,
		eventType:   core.EventType_GameQuited,
		callback:    onGameQuited,
		nextState:   Quited,
	},
	{
		description: "",
		gameState:   Active,
		eventType:   core.EventType_SalvoFired,
		callback:    onSalvoFired,
		nextState:   Active,
	},
	{
		description: "",
		gameState:   Active,
		eventType:   core.EventType_SalvoImpactAssessed,
		callback:    onSalvoImpactAssessed,
		nextState:   Active,
	},
	{
		description: "",
		gameState:   Active,
		eventType:   core.EventType_GameCompleted,
		callback:    onGameCompleted,
		nextState:   Completed,
	},
}

func inviteForGame(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.InviteForGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// 	Validate: fields
		if cmd.GameId == "" || cmd.Initiator == "" || cmd.Invitee == "" {
			return events, fmt.Errorf("Invalid input for command %+v", cmd)
		}

		// Compose events
		pdu := core.InvitedForGame{
			GameId:    cmd.GameId,
			Initiator: cmd.Initiator,
			Invitee:   cmd.Invitee,
		}.ToPdu()
		events = append(events, pdu)

		return events, nil
	}(*pdu.Invite)
}

func acceptGame(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.AcceptGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// 	Validate: fields
		if cmd.GameId == "" {
			return events, fmt.Errorf("Invalid input for command %+v", cmd)
		}

		// Compose events
		pdu := core.GameAccepted{
			GameId:  cmd.GameId,
			Starter: game.Initiator,
		}.ToPdu()
		events = append(events, pdu)

		return events, nil
	}(*pdu.Accept)
}

func rejectGame(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.RejectGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// 	Validate: fields
		if cmd.GameId == "" {
			return events, fmt.Errorf("Invalid input for command %+v", cmd)
		}

		// Compose events
		pdu := core.GameRejected{
			GameId: cmd.GameId,
		}.ToPdu()
		events = append(events, pdu)

		return events, nil
	}(*pdu.Reject)
}

func fireSalvo(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.FireSalvo) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// 	Validate: fields
		if cmd.GameId == "" || cmd.FiredBy == "" || len(cmd.Targets) == 0 {
			return events, fmt.Errorf("Invalid input for command %+v", cmd)
		}

		// Compose events
		pdu := core.SalvoFired{
			GameId:  cmd.GameId,
			FiredBy: cmd.FiredBy,
			Targets: cmd.Targets,
		}.ToPdu()
		events = append(events, pdu)

		return events, nil
	}(*pdu.Fire)
}

func quitGame(s *GameLogicService, game Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
	return func(cmd core.QuitGame) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{}

		// 	Validate: fields
		if cmd.GameId == "" {
			return events, fmt.Errorf("Invalid input for command %+v", cmd)
		}

		// Compose events
		pdu := core.GameQuited{
			GameId: cmd.GameId,
		}.ToPdu()
		events = append(events, pdu)

		return events, nil
	}(*pdu.Quit)
}

func onGridPopulated(s *GameLogicService, game Game, pdu core.GameEventPdu) ([]core.GameEventPdu, error) {
	return func(evt core.GridPopulated) ([]core.GameEventPdu, error) {
		events := []core.GameEventPdu{pdu}

		return events, nil
	}(*pdu.Populated)
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
