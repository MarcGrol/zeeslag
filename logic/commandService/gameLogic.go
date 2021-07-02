package commandService

import (
	"fmt"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

var commandStateDisppatching = []commandGameState{
	{
		description: "",
		gameState:   model.Idle,
		commandType: core.CommandType_InviteForGame,
		callback:    inviteForGame,
		nextState:   model.InvitationPending,
	},
	{
		description: "",
		gameState:   model.InvitationPending,
		commandType: core.CommandType_Quit,
		callback:    quitGame,
		nextState:   model.Quited,
	},
	{
		description: "",
		gameState:   model.InvitationPending,
		commandType: core.CommandType_Accept,
		callback:    acceptGame,
		nextState:   model.Active,
	},
	{
		description: "",
		gameState:   model.InvitationPending,
		commandType: core.CommandType_Reject,
		callback:    rejectGame,
		nextState:   model.Rejected,
	},
	{
		description: "",
		gameState:   model.Active,
		commandType: core.CommandType_Quit,
		callback:    quitGame,
		nextState:   model.Quited,
	},
	{
		description: "",
		gameState:   model.Active,
		commandType: core.CommandType_Fire,
		callback:    fireSalvo,
		nextState:   model.WaitforAssessment,
	},
}

func inviteForGame(s *CommandService, game model.Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
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

func acceptGame(s *CommandService, game model.Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
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

func rejectGame(s *CommandService, game model.Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
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

func fireSalvo(s *CommandService, game model.Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
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

func quitGame(s *CommandService, game model.Game, pdu core.GameCommandPdu) ([]core.GameEventPdu, error) {
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

