package logic

import (
	"log"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/model"
)

var commandStateDisppatching = []commandGameState{
	{
		gameState:   model.Idle,
		commandType: core.CommandType_InviteForGame,
		callback: inviteForGame,
		nextState: model.InvitationPending,
	},
	{
		gameState:   model.InvitationPending,
		commandType: core.CommandType_Quit,
		callback: quitGame,
		nextState: model.Quited,
	},
	{
		gameState:   model.InvitationPending,
		commandType: core.CommandType_Accept,
		callback: acceptGame,
		nextState: model.Active,
	},
	{
		gameState:   model.InvitationPending,
		commandType: core.CommandType_Reject,
		callback: rejectGame,
		nextState: model.Rejected,
	},
	{
		gameState:   model.Active,
		commandType: core.CommandType_Quit,
		callback: quitGame,
		nextState: model.Quited,
	},
	{
		gameState:   model.Active,
		commandType: core.CommandType_Fire,
		callback: fireSalvo,
		nextState: model.Active,
	},
}

var eventStateDispatching = []eventGameState{
	{
		gameState: model.Idle,
		eventType: core.EventType_InvitedForGame,
		callback: onInvitedForGame,
		nextState: model.InvitationPending,
	},
	{
		gameState: model.InvitationPending,
		eventType: core.EventType_GameAccepted,
		callback: onInvitationAccepted,
		nextState: model.Active,
	},
	{
		gameState: model.InvitationPending,
		eventType: core.EventType_GameRejected,
		callback: onInvitationRejected,
		nextState: model.Rejected,
	},
	{
		gameState: model.Active,
		eventType: core.EventType_GameQuited,
		callback: onGameQuited,
		nextState: model.Quited,
	},
	{
		gameState: model.Active,
		eventType: core.EventType_SalvoFired,
		callback: onSalvoFired,
		nextState: model.Active,
	},
	{
		gameState: model.Active,
		eventType: core.EventType_SalvoImpactAssessed,
		callback: onSalvoImpactAssessed,
		nextState: model.Active,
	},
	{
		gameState: model.Active,
		eventType: core.EventType_GameCompleted,
		callback: onGameCompleted,
		nextState: model.Completed,
	},
}

func inviteForGame(s *GameLogicService, game model.Game, pdu core.GameCommandPdu) error {
	return s.inviteForGame(game, *pdu.Invite)
}

func (s *GameLogicService) inviteForGame(game model.Game, cmd core.InviteForGame) error {
	log.Printf("game: %+v, cmd: %+v", game, cmd)

	// Validate: fields

	// Validate: allowed for current state

	// Compose and store events

	return nil
}

func acceptGame(s *GameLogicService, game model.Game, cmd core.GameCommandPdu) error {
	return s.acceptGame(game, *cmd.Accept)
}

func (s *GameLogicService) acceptGame(game model.Game, cmd core.AcceptGame) error {
	log.Printf("game: %+v, cmd: %+v", game, cmd)

	// Validate: fields

	// Validate: allowed for current state

	// Compose and store events

	return nil
}

func rejectGame(s *GameLogicService, game model.Game, pdu core.GameCommandPdu) error {
	return s.rejectGame(game, *pdu.Reject)
}

func (s *GameLogicService) rejectGame(game model.Game, cmd core.RejectGame) error {
	log.Printf("game: %+v, cmd: %+v", game, cmd)

	// Validate: fields

	// Validate: allowed for current state

	// Compose and store events

	return nil
}

func fireSalvo(s *GameLogicService, game model.Game, pdu core.GameCommandPdu) error {
	return s.fireSalvo(game, *pdu.Fire)
}

func (s *GameLogicService) fireSalvo(game model.Game, cmd core.Salvo) error {
	log.Printf("game: %+v, cmd: %+v", game, cmd)

	// Validate: fields

	// Validate: allowed for current state

	// Compose and store events

	return nil
}

func quitGame(s *GameLogicService, game model.Game, pdu core.GameCommandPdu) error {
	return s.quitGame(game, *pdu.Quit)
}

func (s *GameLogicService) quitGame(game model.Game, cmd core.QuitGame) error {
	log.Printf("game: %+v, cmd: %+v", game, cmd)

	// Validate: fields

	// Validate: allowed for current state

	// Compose and store events

	return nil
}

func onInvitedForGame(s *GameLogicService, game model.Game, pdu core.GameEventPdu) error {
	return s.onInvitedForGame(game, *pdu.Invited)
}

func (s *GameLogicService) onInvitedForGame(game model.Game, evt core.InvitedForGame) error {
	log.Printf("game: %+v, evt: %+v", game, evt)

	return nil
}

func onInvitationAccepted(s *GameLogicService, game model.Game, pdu core.GameEventPdu) error {
	return s.onInvitationAccepted(game, *pdu.Accepted)
}

func (s *GameLogicService) onInvitationAccepted(game model.Game, evt core.GameAccepted) error {
	log.Printf("game: %+v, evt: %+v", game, evt)

	return nil
}

func onInvitationRejected(s *GameLogicService, game model.Game, pdu core.GameEventPdu) error {
	return s.onInvitationRejected(game, *pdu.Rejected)
}

func (s *GameLogicService) onInvitationRejected(game model.Game, evt core.GameRejected) error {
	log.Printf("game: %+v, evt: %+v", game, evt)

	return nil
}

func onGameQuited(s *GameLogicService, game model.Game, pdu core.GameEventPdu) error {
	return s.onGameQuited(game, *pdu.Quited)
}

func (s *GameLogicService) onGameQuited(game model.Game, evt core.GameQuited) error {
	log.Printf("game: %+v, evt: %+v", game, evt)

	return nil
}

func onSalvoFired(s *GameLogicService, game model.Game, pdu core.GameEventPdu) error {
	return s.onSalvoFired(game, *pdu.Fired)
}

func (s *GameLogicService) onSalvoFired(game model.Game, evt core.SalvoFired) error {
	log.Printf("game: %+v, evt: %+v", game, evt)

	return nil
}

func onSalvoImpactAssessed(s *GameLogicService, game model.Game, pdu core.GameEventPdu) error {
	return s.onSalvoImpactAssessed(game, *pdu.Assessed)
}

func (s *GameLogicService) onSalvoImpactAssessed(game model.Game, evt core.SalvoImpactAssessed) error {
	log.Printf("game: %+v, evt: %+v", game, evt)

	return nil
}

func onGameCompleted(s *GameLogicService, game model.Game, pdu core.GameEventPdu) error {
	return s.onGameCompleted(game, *pdu.Completed)
}

func (s *GameLogicService) onGameCompleted(game model.Game, evt core.GameCompleted) error {
	log.Printf("game: %+v, evt: %+v", game, evt)

	return nil
}
