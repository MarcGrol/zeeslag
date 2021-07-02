package core

type CommandServicer interface {
	OnCommand(cmd GameCommandPdu) error
}

type CommandType int

const (
	CommandType_Unknown CommandType = iota
	CommandType_InviteForGame
	CommandType_Accept
	CommandType_Reject
	CommandType_Fire
	CommandType_Quit
)

func (ct CommandType) String() string {
	switch ct {
	case CommandType_InviteForGame:
		return "invite"
	case CommandType_Accept:
		return "accept"
	case CommandType_Reject:
		return "reject"
	case CommandType_Fire:
		return "fire"
	case CommandType_Quit:
		return "quit"
	default:
		return "unknown"
	}
}

type GameCommandPdu struct {
	GameId      string
	CommandType CommandType

	Invite *InviteForGame
	Accept *AcceptGame
	Reject *RejectGame
	Fire   *FireSalvo
	Quit   *QuitGame
}

type InviteForGame struct {
	GameId    string
	Initiator string
	Invitee   string
}

func (c InviteForGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId:      c.GameId,
		CommandType: CommandType_InviteForGame,
		Invite:      &c,
	}
}

type AcceptGame struct {
	GameId  string
	Starter string
}

func (c AcceptGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId:      c.GameId,
		CommandType: CommandType_Accept,
		Accept:      &c,
	}
}

type RejectGame struct {
	GameId string
}

func (c RejectGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId:      c.GameId,
		CommandType: CommandType_Reject,
		Reject:      &c,
	}
}

type FireSalvo struct {
	GameId  string
	FiredBy string
	Targets []Coordinate
}

func (c FireSalvo) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId:      c.GameId,
		CommandType: CommandType_Fire,
		Fire:        &c,
	}
}

type QuitGame struct {
	GameId string
}

func (c QuitGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		GameId:      c.GameId,
		CommandType: CommandType_Quit,
		Quit:        &c,
	}
}
