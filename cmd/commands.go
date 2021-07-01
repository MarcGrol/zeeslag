package cmd


type CommandService interface {
	OnStart(cmd StartGame) error
	OnAccept(cmd AcceptGame) error
	OnFire(cmd Salvo) error
}

type StartGame struct {
	GameId string
}

func (c StartGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		CommandType: CommandType_Start,
		Start:       &c,
	}
}

type AcceptGame struct {
	GameId string
}

func (c AcceptGame) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		CommandType: CommandType_Accept,
		Accept:      &c,
	}
}

type Salvo struct {
	GameId string
}

func (c Salvo) ToPdu() GameCommandPdu {
	return GameCommandPdu{
		CommandType: CommandType_Fire,
		Fire:        &c,
	}
}

type CommandType int

const (
	CommandType_Start CommandType = iota
	CommandType_Accept
	CommandType_Fire
	CommandType_ProcessDamage
	CommandType_Quit
)

type GameCommandPdu struct {
	GameId string
	CommandType CommandType
	Start       *StartGame
	Accept      *AcceptGame
	Fire        *Salvo
	//ProcessDamage interface{}
	//Quit          interface{}
}
