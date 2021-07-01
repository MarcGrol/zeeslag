package main

import (
        "github.com/google/uuid"
)

//       A                            B
// ------------------startup by A------------------------
// gameInitiated
// gridPopulated
//                                gameAccepted
//                                gridPopulated
// gameAccepted
//
// ------------------salvo by A--------------------------
//
// saloFired
//                                salvoReceived
// salvoImpactProcessed
//
// ------------------salvo by B--------------------------
//
//                                salvoFired
// salvoReceived
//                                salvoImpactProcessed
//
// ------------------salvo by A--------------------------
// ...
// ------------------salvo by B--------------------------
// ...
// ------------------salvo by A--------------------------
// ...
// ------------------salvo by B--------------------------
// ...
// ------------------B has won---------------------------
//

// commands                   A: events stored       B: events stored
// A -> B: initiateGame       gameInitiated          gameAccepted
// B -> A: acceptGame         gameAccepted
//                            gridPopulated          grodPopulated
//
// A -> B: fireSalvo          salvoFired             salvoReceived
// B -> A: reportSalvoImpact  salvoImpactProcessed   salvoImpactProccessed
//
// B -> A: fireSalvo          salvoReceived          salvoFired
// A -> B: reportSalvoImpact                         salvoImpactProcessed
//
// ...
// ...
// A -> B: markCompleted      gameCompleted          gameCompleted

type Player struct {
        name           string
        events         []EventPdu
        commandChannel chan CommandPdu
        games          []Game
}

func NewPlayer(name string) *Player {
        return &Player{
                name:           name,
                events:         []EventPdu{},
                commandChannel: make(chan CommandPdu),
                games:          []Game{},
        }
}

func (p *Player) Start() {
        go p.run()
}

func (p *Player) Command(cmd CommandPdu) {
        p.commandChannel <- cmd
}

func (p *Player) run() {
        // Listen for commmands from players
        for {
                <-p.commandChannel
        }
        // Listen for events from other ship
}

type CommandType int

const (
        CommandType_Start CommandType = iota
        CommandType_Accept
        CommandType_Fire
        CommandType_ProcessDamage
        CommandType_Quit
)

type CommandPdu struct {
        CommandType   CommandType
        Start         interface{}
        Accept        interface{}
        Fire          interface{}
        ProcessDamage interface{}
        Quit          interface{}
}

type EventType int

const (
        EventType_GameInitiated EventType = iota
        EventType_GameAccepted
        EventType_SalvoFired
        EventType_SalvoReceived
        EventType_SalvoImpactProcessed
        EventType_GameCompleted
)

type EventPdu struct {
        EventType EventType
        Initiated  *GameInitiated
        Accepted  *GameAccepted
        Rejected  *GameRejected
        Fired     *SalvoFired
        Received  *SalvoReceived
        Feedback  *SalvoImpactProcessed
        Aborted   *GameAborted
        Completed *GameCompleted
}

// events
type GameInitiated struct {
        GameId    string
        Initiator string
        Invitee   string
}

type GameRejected struct {
        GameId string
}

type GameAccepted struct {
        GameId   string
        Starting string
}

type GridPopulated struct {
        GameId string
        Winger []Coordinate
        Angle  []Coordinate
        AClass []Coordinate
        BClass []Coordinate
        SClass []Coordinate
}

type SalvoFired struct {
        GameId  string
        FiredBy string
        Targets  []Coordinate
}

type SalvoReceived struct {
        GameId  string
        FiredBy string
        Targets []Coordinate
}

type Coordinate struct {
        Row    int
        Column int
}

type SalvoImpactProcessed struct {
        GameId      string
        FiredBy     string
        TargetStatuses  []TargetStatus
}

type TargetStatus struct {
        Target     Coordinate
        SalvoStatus SalvoStatus
}

type SalvoStatus int

const (
        Miss SalvoStatus = iota
        Hit
        Kill
)

type GameAborted struct {
        GameId    string
        AbortedBy string
}

type GameCompleted struct {
        GameId string
        Winner string
}

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

type Salvo struct {
        Target      Coordinate
        SalvoStatus SalvoStatus
}

type Board struct {
        grid [16][16]CellSalvoStatus
}

func (b *Board) markCellStatus(target Coordinate, cellStatus CellStatus) {
        // TODO check index ranges
        b.grid[target.Row][target.Column].CellStatus = cellStatus
}

func (b *Board) markSalvoStatus(target Coordinate, salvoStatus SalvoStatus) {
        // TODO check index ranges
        b.grid[target.Row][target.Column].SalvoStatus = salvoStatus
}

type CellSalvoStatus struct {
        CellStatus  CellStatus
        SalvoStatus SalvoStatus
}

func (s CellSalvoStatus) Striing() string {
        if s.SalvoStatus == Hit {
                return "x"
        }

        if s.SalvoStatus == Miss {
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

func NewGame() *Game {
        return &Game{}
}

func (g *Game) ApplyAll(evts []EventPdu) {
        for _, e := range evts {
                g.Apply(e)
        }
}

func (g *Game) Apply(evt EventPdu) {
        switch evt.EventType {
        case EventType_GameInitiated:
                g.ApplyGameInitiated(*evt.Started)
        case EventType_GameAccepted:
                g.ApplyGameAccepted(*evt.Accepted)
        case EventType_SalvoFired:
                g.ApplySalvoFired(*evt.Fired)
        case EventType_SalvoReceived:
                g.ApplySalvoReceived(*evt.Received)
        case EventType_SalvoImpactProcessed:
                g.ApplyGameInitiated(*evt.Started)
        case EventType_GameCompleted:
                g.ApplyGameInitiated(*evt.Started)
        default:
                // TODO
        }
}

func (g *Game) ApplyGameInitiated(evt GameInitiated) {
        g.GameId = evt.GameId
        g.Initiator = evt.Initiator
        g.Invitee = evt.Invitee
        g.Status = Requested
}

func (g *Game) ApplyGameAccepted(evt GameAccepted) {
        g.GameId = evt.GameId
}

func (g *Game) ApplyGridPopulated(evt GridPopulated) {
        g.GameId = evt.GameId
}

func (g *Game) ApplyGameRejected(evt GameRejected) {
        g.GameId = evt.GameId
}

func (g *Game) ApplySalvoFired(evt SalvoFired) {
        g.GameId = evt.GameId
}

func (g *Game) ApplySalvoImpactProcessed(evt SalvoImpactProcessed) {
        g.GameId = evt.GameId
}

func (g *Game) ApplyGameAborted(evt GameAborted) {
        g.GameId = evt.GameId
}

func (g *Game) ApplyComplated(evt GameCompleted) {
        g.GameId = evt.GameId
}

func main() {

        gameId := uuid.New().String()

        firstPlayer := NewPlayer("Marc")
        firstPlayer.Start()

        firstPlayerGame := NewGame()
        ///secondPlayerGame := NewGame()

        firstPlayerGame.ApplyGameInitiated(GameInitiated{
                GameId:    gameId,
                Initiator: "marc",
                Invitee:   "Eva",
        })

        firstPlayerGame.ApplyGameAccepted(GameAccepted{
                GameId: gameId,
        })

        firstPlayerGame.ApplySalvoFired(SalvoFired{
                GameId: gameId,
        })
}
