package core

type CommandService interface {
	OnCommand(cmd GameCommandPdu) error
}

type CommandType int

const (
	CommandType_InviteForGame CommandType = iota
	CommandType_Accept
	CommandType_Reject
	CommandType_Fire
	CommandType_Quit
)

type GameCommandPdu struct {
	GameId      string
	CommandType CommandType

	Invite      *InviteForGame
	Accept      *AcceptGame
	Reject      *RejectGame
	Fire        *Salvo
	Quit        *QuitGame
}

func (p GameCommandPdu)GetId() string {
	return ""
}

type InviteForGame struct {
	GameId string
}

func (c InviteForGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId:      c.GameId,
		CommandType: CommandType_InviteForGame,
		Invite:      &c,
	}
}

type AcceptGame struct {
	GameId string
}

func (c AcceptGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId: c.GameId,
		CommandType: CommandType_Accept,
		Accept:      &c,
	}
}


type RejectGame struct {
	GameId string
}

func (c RejectGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId: c.GameId,
		CommandType: CommandType_Reject,
		Reject:      &c,
	}
}

type Salvo struct {
	GameId string
}

func (c Salvo) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId: c.GameId,
		CommandType: CommandType_Fire,
		Fire:        &c,
	}
}

type QuitGame struct {
	GameId string
}

func (c QuitGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId: c.GameId,
		CommandType: CommandType_Quit,
		Quit:        &c,
	}
}

