package eventService

import (
	"testing"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/infra"
	"github.com/MarcGrol/zeeslag/logic/repo"
	"github.com/MarcGrol/zeeslag/model"
	"github.com/stretchr/testify/assert"
)

func TestPopulated(t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{}
	cmd := core.GridPopulated{
		GameId: "1",
	}

	// when
	game, err := when(preconditions, cmd.GameId, func(sut *EventService) error {
		return sut.OnEvent(cmd.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, cmd.GameId, game.GameId)
}

func NoTestInvited(t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{
		core.GridPopulated{
			GameId: "1",
		}.ToPdu(),
	}
	event := core.InvitedForGame{
		GameId:    "1",
		Initiator: "me",
		Invitee:   "you",
	}

	// when
	game, err := when(preconditions, event.GameId, func(sut *EventService) error {
		return sut.OnEvent(event.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, event.GameId, game.GameId)
	assert.Equal(t, event.Initiator, game.Initiator)
	assert.Equal(t, event.Invitee, game.Invitee)
}

func TestAccepted(t *testing.T) {
	// given

	// when

	// then

}

func TestFired(t *testing.T) {
	// given

	// when

	// then

}

func TestAssessed(t *testing.T) {
	// given

	// when

	// then

}

func TestCompleted(t *testing.T) {
	// given

	// when

	// then

}

func TestQuited(t *testing.T) {
	// given

	// when

	// then

}

func when(preconditions []core.GameEventPdu, gameId string, testFunc func(service *EventService) error) (*model.Game, error) {
	repo := repo.NewGameRepository(infra.NewBasicEventStore(), infra.NewBasicPubsub())
	sut := NewEventService(repo)

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
