package evt

type EventType int

const (
	EventType_GameInitiated EventType = iota
	EventType_GameAccepted
	EventType_GameRejected
	EventType_GridPopulated
	EventType_SalvoFired
	EventType_SalvoReceived
	EventType_SalvoImpactProcessed
	EventType_GameAborted
	EventType_GameCompleted
)

type GameEventPdu struct {
	EventType EventType
	Initiated *GameInitiated
	Accepted  *GameAccepted
	Rejected  *GameRejected
	Populated *GridPopulated
	Fired     *SalvoFired
	Received  *SalvoReceived
	Processed *SalvoImpactProcessed
	Aborted   *GameAborted
	Completed *GameCompleted
}

// events
type GameInitiated struct {
	GameId    string
	Initiator string
	Invitee   string
}

func NewGameInitiatedFromPdu(pdu GameEventPdu) (*GameInitiated, bool) {
	return pdu.Initiated, pdu.EventType == EventType_GameInitiated
}

func (e GameInitiated) ToPdu() GameEventPdu {
	return GameEventPdu{
		EventType: EventType_GameInitiated,
		Initiated: &e,
	}
}

type GameRejected struct {
	GameId string
}

func NewGameRejectedFromPdu(pdu GameEventPdu) (*GameRejected, bool) {
	return pdu.Rejected, pdu.EventType == EventType_GameRejected
}

func (e GameRejected) ToPdu() GameEventPdu {
	return GameEventPdu{
		EventType: EventType_GameRejected,
		Rejected:  &e,
	}
}

type GameAccepted struct {
	GameId   string
	Starting string
}

func NewGameAcceptedFromPdu(pdu GameEventPdu) (*GameAccepted, bool) {
	return pdu.Accepted, pdu.EventType == EventType_GameAccepted
}

func (e GameAccepted) ToPdu() GameEventPdu {
	return GameEventPdu{
		EventType: EventType_GameAccepted,
		Accepted:  &e,
	}
}

type GridPopulated struct {
	GameId string
	Winger []Coordinate
	Angle  []Coordinate
	AClass []Coordinate
	BClass []Coordinate
	SClass []Coordinate
}

func NewGridPopulatedFromPdu(pdu GameEventPdu) (*GridPopulated, bool) {
	return pdu.Populated, pdu.EventType == EventType_GridPopulated
}

func (e GridPopulated) ToPdu() GameEventPdu {
	return GameEventPdu{
		EventType: EventType_GridPopulated,
		Populated: &e,
	}
}

type SalvoFired struct {
	GameId  string
	FiredBy string
	Targets []Coordinate
}

func NewSalvoFiredFromPdu(pdu GameEventPdu) (*SalvoFired, bool) {
	return pdu.Fired, pdu.EventType == EventType_SalvoFired
}

func (e SalvoFired) ToPdu() GameEventPdu {
	return GameEventPdu{
		EventType: EventType_SalvoFired,
		Fired:     &e,
	}
}

type SalvoReceived struct {
	GameId  string
	FiredBy string
	Targets []Coordinate
}

func NewSalvoReceivedFromPdu(pdu GameEventPdu) (*SalvoReceived, bool) {
	return pdu.Received, pdu.EventType == EventType_SalvoReceived
}

func (e SalvoReceived) ToPdu() GameEventPdu {
	return GameEventPdu{
		EventType: EventType_SalvoReceived,
		Received:  &e,
	}
}

type Coordinate struct {
	Row    int
	Column int
}

type SalvoImpactProcessed struct {
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

func NewImpactProcessedFromPdu(pdu GameEventPdu) (*SalvoImpactProcessed, bool) {
	return pdu.Processed, pdu.EventType == EventType_SalvoImpactProcessed
}

func (e SalvoImpactProcessed) ToPdu() GameEventPdu {
	return GameEventPdu{
		EventType: EventType_SalvoImpactProcessed,
		Processed: &e,
	}
}

type GameAborted struct {
	GameId    string
	AbortedBy string
}

func NewAbortedFromPdu(pdu GameEventPdu) (*GameAborted, bool) {
	return pdu.Aborted, pdu.EventType == EventType_GameAborted
}

func (e GameAborted) ToPdu() GameEventPdu {
	return GameEventPdu{
		EventType: EventType_GameAborted,
		Aborted:   &e,
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
		EventType: EventType_GameCompleted,
		Completed: &e,
	}
}
