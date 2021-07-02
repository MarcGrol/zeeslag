package core

type ReplicationServicer interface {
	OnPeerEvent(msg GameMsgPdu) error
}

type ReplicatiomMsgType int

const (
	ReplicatiomMsgType_Unknown = iota
	ReplicatiomMsgType_PeerHasInvitedYouForGame
	ReplicatiomMsgType_PeerHasAcceptedInvitation
	ReplicatiomMsgType_PeerHasRejectedInvitation
	ReplicatiomMsgType_PeerHasFiredSalvo
	ReplicatiomMsgType_PeerHasAssessedImpactOfSalvo
	ReplicatiomMsgType_PeerHasCompleted
	ReplicatiomMsgType_PeerHasQuited
)

func (et ReplicatiomMsgType) String() string {
	switch et {
	case ReplicatiomMsgType_PeerHasInvitedYouForGame:
		return "peer-invited"
	case ReplicatiomMsgType_PeerHasAcceptedInvitation:
		return "peer-accepted"
	case ReplicatiomMsgType_PeerHasRejectedInvitation:
		return "peer-rejected"
	case ReplicatiomMsgType_PeerHasFiredSalvo:
		return "peer-fired"
	case ReplicatiomMsgType_PeerHasAssessedImpactOfSalvo:
		return "peer-assessed"
	case ReplicatiomMsgType_PeerHasCompleted:
		return "peer-completed"
	case ReplicatiomMsgType_PeerHasQuited:
		return "peer-quited"
	default:
		return "unknown"
	}
}

type GameMsgPdu struct {
	GameId  string
	MsgType ReplicatiomMsgType

	Invited   *PeerHasInvitedYouForGame
	Accepted  *PeerHasAcceptedGame
	Rejected  *PeerHasRejectedInvitation
	Fired     *PeerHasFiredSalvo
	Assessed  *PeerHasAssessedImpactOfSalvo
	Completed *PeerHasCompleted
	Quited    *PeerHasQuited
}

type PeerHasAssessedImpactOfSalvo struct {
	GameId string
}

func (e PeerHasAssessedImpactOfSalvo) ToPdu() GameMsgPdu {
	return GameMsgPdu{
		GameId:   e.GameId,
		MsgType:  ReplicatiomMsgType_PeerHasInvitedYouForGame,
		Assessed: &e,
	}
}

type PeerHasFiredSalvo struct {
	GameId string
}

func (e PeerHasFiredSalvo) ToPdu() GameMsgPdu {
	return GameMsgPdu{
		GameId:  e.GameId,
		MsgType: ReplicatiomMsgType_PeerHasFiredSalvo,
		Fired:   &e,
	}
}

type PeerHasRejectedInvitation struct {
	GameId string
}

func (e PeerHasRejectedInvitation) ToPdu() GameMsgPdu {
	return GameMsgPdu{
		GameId:   e.GameId,
		MsgType:  ReplicatiomMsgType_PeerHasRejectedInvitation,
		Rejected: &e,
	}
}

type PeerHasAcceptedGame struct {
	GameId string
}

func (e PeerHasAcceptedGame) ToPdu() GameMsgPdu {
	return GameMsgPdu{
		GameId:   e.GameId,
		MsgType:  ReplicatiomMsgType_PeerHasAcceptedInvitation,
		Accepted: &e,
	}
}

type PeerHasInvitedYouForGame struct {
	GameId    string
	Initiator string
	Invitee   string
}

func (e PeerHasInvitedYouForGame) ToPdu() GameMsgPdu {
	return GameMsgPdu{
		GameId:  e.GameId,
		MsgType: ReplicatiomMsgType_PeerHasInvitedYouForGame,
		Invited: &e,
	}
}

type PeerHasCompleted struct {
	GameId string
}

func (e PeerHasCompleted) ToPdu() GameMsgPdu {
	return GameMsgPdu{
		GameId:    e.GameId,
		MsgType:   ReplicatiomMsgType_PeerHasCompleted,
		Completed: &e,
	}
}

type PeerHasQuited struct {
	GameId string
}

func (e PeerHasQuited) ToPdu() GameMsgPdu {
	return GameMsgPdu{
		GameId:  e.GameId,
		MsgType: ReplicatiomMsgType_PeerHasQuited,
		Quited:  &e,
	}
}
