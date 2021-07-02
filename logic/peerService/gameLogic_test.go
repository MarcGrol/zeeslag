package peerService

import (
	"testing"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/infra"
	"github.com/MarcGrol/zeeslag/logic/repo"
	"github.com/MarcGrol/zeeslag/model"
	"github.com/stretchr/testify/assert"
)

func TestInvited(t *testing.T) {
	// given
	preconditions := []core.GameEventPdu{}

	event := core.InvitedForGame{
		GameId:    "1",
		Initiator: "me",
		Invitee:   "you",
	}

	// when
	game, err := when(preconditions, event.GameId, func(sut *PeerService) error {
		return sut.OnEvent(event.ToPdu())
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, event.GameId, game.GameId)
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

func when(preconditions []core.GameEventPdu, gameId string, testFunc func(service *PeerService) error) (*model.Game, error) {
	repo := repo.NewGameRepository(infra.NewBasicEventStore(), infra.NewBasicPubsub())
	sut := NewPeerService(repo)

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
