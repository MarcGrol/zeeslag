package core

type EventServicer interface {
	OnEvent(cmd GameEventPdu) error
}

type EventType int

const (
	EventType_Unknown = iota
	EventType_GridPopulated
	EventType_GameInvitationSent
	EventType_GameInvitationReceived
	EventType_GameInvitationAccepted
	EventType_GameInvitationRejected
	EventType_SalvoFired
	EventType_SalvoImpactAssessed
	EventType_GameQuited
	EventType_GameCompleted
)

func (et EventType) String() string {
	switch et {
	case EventType_GridPopulated:
		return "populated"
	case EventType_GameInvitationSent:
		return "invitation-sent"
	case EventType_GameInvitationReceived:
		return "invitation-received"
	case EventType_GameInvitationAccepted:
		return "accepted"
	case EventType_GameInvitationRejected:
		return "rejected"
	case EventType_SalvoImpactAssessed:
		return "assessed"
	case EventType_SalvoFired:
		return "fired"
	case EventType_GameCompleted:
		return "completed"
	default:
		return "unknown"
	}
}

type GameEventPdu struct {
	GameId    string
	EventType EventType

	Populated      *GridPopulated
	InvitationSent *GameInvitationSent
	Invited        *GameInvitionReceived
	Accepted       *GameAccepted
	Rejected       *GameRejected
	Fired          *SalvoFired
	Assessed       *SalvoImpactAssessed
	Quited         *GameQuited
	Completed      *GameCompleted
}

type GridPopulated struct {
	GameId string
	Winger []Coordinate
	Angle  []Coordinate
	AClass []Coordinate
	BClass []Coordinate
	SClass []Coordinate
}

func (e GridPopulated) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:    e.GameId,
		EventType: EventType_GridPopulated,
		Populated: &e,
	}
}

// events
type GameInvitationSent struct {
	GameId    string
	Initiator string
	Invitee   string
}

func (e GameInvitationSent) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:         e.GameId,
		EventType:      EventType_GameInvitationSent,
		InvitationSent: &e,
	}
}

// events
type GameInvitionReceived struct {
	GameId    string
	Initiator string
	Invitee   string
}

func (e GameInvitionReceived) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:    e.GameId,
		EventType: EventType_GameInvitationReceived,
		Invited:   &e,
	}
}

type GameRejected struct {
	GameId string
}

func (e GameRejected) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:    e.GameId,
		EventType: EventType_GameInvitationRejected,
		Rejected:  &e,
	}
}

type GameAccepted struct {
	GameId  string
	Starter string
}

func (e GameAccepted) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:    e.GameId,
		EventType: EventType_GameInvitationAccepted,
		Accepted:  &e,
	}
}

type SalvoFired struct {
	GameId  string
	SalvoId string
	FiredBy string
	Targets []Coordinate
}

func (e SalvoFired) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:    e.GameId,
		EventType: EventType_SalvoFired,
		Fired:     &e,
	}
}

type Coordinate struct {
	Row    int
	Column int
}

type SalvoImpactAssessed struct {
	GameId         string
	FiredBy        string
	TargetStatuses []TargetStatus
}

type TargetStatus struct {
	Target      Coordinate
	SalvoStatus SalvoStatus
}

type SalvoStatus int

const (
	Miss SalvoStatus = iota
	Hit
	Kill
)

func (e SalvoImpactAssessed) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:    e.GameId,
		EventType: EventType_SalvoImpactAssessed,
		Assessed:  &e,
	}
}

type GameQuited struct {
	GameId    string
	AbortedBy string
}

func (e GameQuited) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:    e.GameId,
		EventType: EventType_GameQuited,
		Quited:    &e,
	}
}

type GameCompleted struct {
	GameId string
	Winner string
}

func NewCompletedFromPdu(pdu GameEventPdu) (*GameCompleted, bool) {
	return pdu.Completed, pdu.EventType == EventType_GameCompleted
}

func (e GameCompleted) ToPdu() GameEventPdu {
	return GameEventPdu{
		GameId:    e.GameId,
		EventType: EventType_GameCompleted,
		Completed: &e,
	}
}
