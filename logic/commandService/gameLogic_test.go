package commandService

import (
	"testing"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/infra"
	"github.com/MarcGrol/zeeslag/logic/repo"
	"github.com/MarcGrol/zeeslag/model"
	"github.com/stretchr/testify/assert"
)

func TestInvite(t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{}
	event := core.InviteForGame{
		GameId:    "1",
		Initiator: "me",
		Invitee:   "you",
	}

	// when
	game, err := when(preconditions, event.GameId, func(sut *CommandService) error {
		return sut.OnCommand(event.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, event.GameId, game.GameId)
	assert.Equal(t, event.Initiator, game.Initiator)
	assert.Equal(t, event.Invitee, game.Invitee)
}

func TestAccept(t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{
		core.InvitedForGame{
			GameId:    "1",
			Initiator: "me",
			Invitee:   "you",
		}.ToPdu(),
	}
	command := core.AcceptGame{
		GameId:  "1",
		Starter: "me",
	}

	// when
	game, err := when(preconditions, command.GameId, func(sut *CommandService) error {
		return sut.OnCommand(command.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, command.GameId, game.GameId)
	assert.Equal(t, command.Starter, game.Starter)
}

func TestFire(t *testing.T) {
	// given
	gameId := "1"
	preconditions := []core.GameEventPdu{
		core.InvitedForGame{
			GameId:    gameId,
			Initiator: "me",
			Invitee:   "you",
		}.ToPdu(),
		core.GameAccepted{
			GameId: gameId,
		}.ToPdu(),
	}
	command := core.FireSalvo{
		GameId:  gameId,
		FiredBy: "mw",
		Targets: []core.Coordinate{
			{Row: 1, Column: 2},
		},
	}

	// when
	game, err := when(preconditions, command.GameId, func(sut *CommandService) error {
		return sut.OnCommand(command.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, command.GameId, game.GameId)
}

func when(preconditions []core.GameEventPdu, gameId string, testFunc func(service *CommandService) error) (*model.Game, error) {
	repo := repo.NewGameRepository(infra.NewBasicEventStore(), infra.NewBasicPubsub())
	peerer := infra.NewBasicPeer("")
	sut := NewCommandService(repo, peerer)

	// force preconditions to be set
	err := repo.StoreEvents(preconditions)
	if err != nil {
		return nil, err
	}

	err = testFunc(sut)
	if err != nil {
		return nil, err
	}

	game, err := repo.GetGameOnId(gameId)
	if err != nil {
		return nil, err
	}

	return game, nil
}
