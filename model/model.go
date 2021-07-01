package model

import (
	"log"

	"github.com/MarcGrol/zeeslag/evt"
)

// Aggregate
type Game struct {
	GameId         string
	Initiator      string
	Invitee        string
	Starting       string
	Status         GameStatus
	SalvosFired    []Salvo
	SalvosReceived []Salvo
	MyBoard        Board
	HisBoard       Board
}

func NewGame(events []	evt.GameEventPdu) *Game {
	game := &Game{}
	game.ApplyAll(events)
	return game
}

type Salvo struct {
	Target      evt.Coordinate
	SalvoStatus evt.SalvoStatus
}

type Board struct {
	grid [16][16]CellSalvoStatus
}

func (b *Board) markCellStatus(target evt.Coordinate, cellStatus CellStatus) {
	// TODO check index ranges
	b.grid[target.Row][target.Column].CellStatus = cellStatus
}

func (b *Board) markSalvoStatus(target evt.Coordinate, salvoStatus evt.SalvoStatus) {
	// TODO check index ranges
	b.grid[target.Row][target.Column].SalvoStatus = salvoStatus
}

type CellSalvoStatus struct {
	CellStatus  CellStatus
	SalvoStatus evt.SalvoStatus
}

func (s CellSalvoStatus) Striing() string {
	if s.SalvoStatus == evt.Hit {
		return "x"
	}

	if s.SalvoStatus == evt.Miss {
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
	return `................
................
................
................
................
................
................
................
................
................
................
................
................
................
................
................`
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
	Initial GameStatus = iota
	Requested
	Rejected
	Accepted
	WaitforOpponent
	Aborted
	Done
)

func (g *Game) ApplyAll(evts []evt.GameEventPdu) {
	for _, e := range evts {
		g.Apply(e)
	}
}

func (g *Game) Apply(event evt.GameEventPdu) {
	switch event.EventType {
	case evt.EventType_GameInitiated	:
		g.ApplyGameInitiated(*event.Initiated)
	case evt.EventType_GameAccepted:
		g.ApplyGameAccepted(*event.Accepted)
	case evt.EventType_GridPopulated:
		g.ApplyGridPopulated(*event.Populated)
	case evt.EventType_GameRejected:
		g.ApplyGameRejected(*event.Rejected)
	case evt.EventType_SalvoFired:
		g.ApplySalvoFired(*event.Fired)
	case evt.EventType_SalvoReceived:
		g.ApplySalvoReceived(*event.Received)
	case evt.EventType_SalvoImpactProcessed:
		g.ApplySalvoImpactProcessed(*event.Processed)
	case evt.EventType_GameAborted:
		g.ApplyGameAborted(*event.Aborted)
	case evt.EventType_GameCompleted:
		g.ApplyGameCompleted(*event.Completed)
	default:
		log.Fatalf("Unrecogized event: %+v", event.EventType)
	}
}

func (g *Game) ApplyGameInitiated(evt evt.GameInitiated) {
	g.GameId = evt.GameId
	g.Initiator = evt.Initiator
	g.Invitee = evt.Invitee
	g.Status = Requested
}

func (g *Game) ApplyGameAccepted(evt evt.GameAccepted) {
	g.GameId = evt.GameId
}

func (g *Game) ApplyGridPopulated(evt evt.GridPopulated) {
	g.GameId = evt.GameId
}

func (g *Game) ApplyGameRejected(evt evt.GameRejected) {
	g.GameId = evt.GameId
}

func (g *Game) ApplySalvoFired(evt evt.SalvoFired) {
	g.GameId = evt.GameId
}

func (g *Game) ApplySalvoReceived(evt evt.SalvoReceived) {
	g.GameId = evt.GameId
}

func (g *Game) ApplySalvoImpactProcessed(evt evt.SalvoImpactProcessed) {
	g.GameId = evt.GameId
}

func (g *Game) ApplyGameAborted(evt evt.GameAborted) {
	g.GameId = evt.GameId
}

func (g *Game) ApplyGameCompleted(evt evt.GameCompleted) {
	g.GameId = evt.GameId
}
