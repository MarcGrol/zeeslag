package logic

import (
	"testing"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/infra"
	"github.com/stretchr/testify/assert"
)

func TestPopulated( t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{}
	cmd := core.GridPopulated{
		GameId: "1",
	}

	// when
	game, err := when(preconditions, cmd.GameId, func( sut core.Service) error{
		return sut.OnEvent(cmd.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, cmd.GameId, game.GameId)
}

func TestInvite( t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{}
	event := core.InviteForGame{
		GameId:    "1",
		Initiator: "me",
		Invitee:   "you",
	}

	// when
	game, err := when(preconditions, event.GameId, func( sut core.Service) error{
		return sut.OnCommand(event.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, event.GameId, game.GameId)
	assert.Equal(t, event.Initiator, game.Initiator)
	assert.Equal(t, event.Invitee, game.Invitee)
}

func NoTestInvited( t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{
			core.GridPopulated{
				GameId:    "1",
			}.ToPdu(),
	}
	event := core.InvitedForGame{
		GameId: "1",
		Initiator: "me",
		Invitee:   "you",
	}

	// when
	game, err := when(preconditions, event.GameId, 	func( sut core.Service) error{
		return sut.OnEvent(event.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, event.GameId, game.GameId)
	assert.Equal(t, event.Initiator, game.Initiator)
	assert.Equal(t, event.Invitee, game.Invitee)
}

func TestAccept( t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{
		core.InvitedForGame{
			GameId:    "1",
			Initiator: "me",
			Invitee:   "you",
		}.ToPdu(),
	}
	command := core.AcceptGame{
		GameId:    "1",
	}

	// when
	game, err := when(preconditions, command.GameId, func( sut core.Service) error{
		return sut.OnCommand(command.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, command.GameId, game.GameId)
}

func TestAccepted( t *testing.T) {
	// given

	// when

	// then

}

func TestFire( t *testing.T) {
	// given
	gameId := "1"
	preconditions := []core.GameEventPdu{
		core.InvitedForGame{
			GameId:    gameId,
			Initiator: "me",
			Invitee:   "you",
		}.ToPdu(),
		core.GameAccepted{
			GameId:    gameId,
		}.ToPdu(),
	}
	command := core.FireSalvo{
		GameId:  gameId,
		FiredBy: "mw",
		Targets: []core.Coordinate{
			{Row:    1, Column: 2},
		},
	}

	// when
	game, err := when(preconditions, command.GameId, func( sut core.Service) error{
		return sut.OnCommand(command.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, command.GameId, game.GameId)
}

func TestFired( t *testing.T) {
	// given

	// when

	// then

}

func TestAssessed( t *testing.T) {
	// given

	// when

	// then

}

func TestCompleted( t *testing.T) {
	// given

	// when

	// then

}

func TestQuited( t *testing.T) {
	// given

	// when

	// then

}


func when(preconditions []core.GameEventPdu, gameId string, testFunc func(core.Service) error) (*Game,error){
	repo := NewGameRepository(infra.NewBasicEventStore())
	sut := NewGameLogicService(repo)

	// force preconditions to be set
	for _, e := range preconditions {
		err := repo.StoreEvent(e)
		if err != nil {
			return nil, err
		}
	}

	err := testFunc(sut)
	if err != nil {
		return nil, err
	}

	game, err := repo.GetGameOnId(gameId)
	if err != nil {
		return nil, err
	}

	return game, nil
}
