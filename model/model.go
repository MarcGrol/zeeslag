package model

import (
	"github.com/MarcGrol/zeeslag/core"
	"log"
)

// Aggregate
type Game struct {
	GameId         string
	Initiator      string
	Invitee        string
	Starter        string
	Status         GameStatus
	SalvosFired    []Salvo
	SalvosReceived []Salvo
	MyBoard        Board
	HisBoard       Board
}

func NewGame(events []core.GameEventPdu) *Game {
	game := &Game{}
	game.ApplyAll(events)
	return game
}

type Salvo struct {
	Target      core.Coordinate
	SalvoStatus core.SalvoStatus
}

type Board struct {
	grid [16][16]CellSalvoStatus
}

func (b *Board) markCellStatus(target core.Coordinate, cellStatus CellStatus) {
	// TODO check index ranges
	b.grid[target.Row][target.Column].CellStatus = cellStatus
}

func (b *Board) markSalvoStatus(target core.Coordinate, salvoStatus core.SalvoStatus) {
	// TODO check index ranges
	b.grid[target.Row][target.Column].SalvoStatus = salvoStatus
}

type CellSalvoStatus struct {
	CellStatus  CellStatus
	SalvoStatus core.SalvoStatus
}

func (s CellSalvoStatus) String() string {
	if s.SalvoStatus == core.Hit {
		return "x"
	}

	if s.SalvoStatus == core.Miss {
		return "-"
	}

	if s.CellStatus == Filled {
		return "*"
	}

	return "."
}

func NewRandom() *Board {
	b := &Board{}

	return b
}

func (b Board) String() string {
	return ""
	//	`................
	//................
	//................
	//................
	//................
	//................
	//................
	//................
	//................
	//................
	//................
	//................
	//................
	//................
	//................
	//................`
}

type CellStatus int

const (
	Empty CellStatus = iota
	Filled
)

type Winger struct{}
type Angle struct{}
type AClass struct{}
type BClass struct{}
type SClass struct{}

type GameStatus int

const (
	Idle GameStatus = iota
	Created
	InvitationPending
	Invited
	Rejected
	Active
	WaitforAssessment
	Quited
	Completed
)

func (s GameStatus) String() string {
	switch s {
	case Idle:
		return "status-idle"
	case Created:
		return "status-created"
	case InvitationPending:
		return "status-pending"
	case Invited:
		return "status-invited"
	case Rejected:
		return "status-rejeected"
	case Active:
		return "status-active"
	case Quited:
		return "status-quited"
	case Completed:
		return "status-completed"
	default:
		return "status-unknown"
	}
}

func (g *Game) ApplyAll(evts []core.GameEventPdu) {
	for _, e := range evts {
		g.Apply(e)
	}
}

func (g *Game) Apply(event core.GameEventPdu) {
	switch event.EventType {
	case core.EventType_GridPopulated:
		g.ApplyGridPopulated(*event.Populated)
	case core.EventType_GameInvitationSent:
		g.ApplyInvitationForGameSent(*event.InvitationSent)
	case core.EventType_GameInvitationReceived:
		g.ApplyInvitedForGame(*event.Invited)
	case core.EventType_GameInvitationAccepted:
		g.ApplyGameAccepted(*event.Accepted)
	case core.EventType_GameInvitationRejected:
		g.ApplyGameRejected(*event.Rejected)
	case core.EventType_SalvoFired:
		g.ApplySalvoFired(*event.Fired)
	case core.EventType_SalvoImpactAssessed:
		g.ApplySalvoImpactAssessed(*event.Assessed)
	case core.EventType_GameQuited:
		g.ApplyGameAborted(*event.Quited)
	case core.EventType_GameCompleted:
		g.ApplyGameCompleted(*event.Completed)
	default:
		log.Fatalf("Unrecogized event: %+v", event.EventType)
	}
}

func (g *Game) ApplyGridPopulated(evt core.GridPopulated) {
	g.GameId = evt.GameId
	g.Status = Created
}

func (g *Game) ApplyInvitationForGameSent(evt core.GameInvitationSent) {
	g.GameId = evt.GameId
	g.Initiator = evt.Initiator
	g.Invitee = evt.Invitee
	g.Status = InvitationPending
}

func (g *Game) ApplyInvitedForGame(evt core.GameInvitionReceived) {
	g.GameId = evt.GameId
	g.Initiator = evt.Initiator
	g.Invitee = evt.Invitee
	g.Status = Invited
}

func (g *Game) ApplyGameAccepted(evt core.GameAccepted) {
	g.GameId = evt.GameId
	g.Starter = evt.Starter
	g.Status = Active
}

func (g *Game) ApplyGameRejected(evt core.GameRejected) {
	g.GameId = evt.GameId
	g.Status = Rejected
}

func (g *Game) ApplySalvoFired(evt core.SalvoFired) {
	g.GameId = evt.GameId
	g.Status = WaitforAssessment
}

func (g *Game) ApplySalvoImpactAssessed(evt core.SalvoImpactAssessed) {
	g.GameId = evt.GameId
	g.Status = Active
}

func (g *Game) ApplyGameAborted(evt core.GameQuited) {
	g.GameId = evt.GameId
	g.Status = Quited
}

func (g *Game) ApplyGameCompleted(evt core.GameCompleted) {
	g.GameId = evt.GameId
	g.Status = Completed
}
