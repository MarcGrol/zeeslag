package logic

import (
	"github.com/MarcGrol/zeeslag/core"
	"github.com/MarcGrol/zeeslag/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPopulated( t *testing.T) {
	// given
	gameId := "1"

	// when
	repo := NewGameRepository(infra.NewBasicEventStore())
	sut := NewGameLogicService(repo)
	err := sut.OnEvent(core.GridPopulated{
		GameId: gameId,
	}.ToPdu())

	// then
	assert.NoError(t, err)

	game, err := repo.GetGameOnId(gameId)
	assert.NoError(t, err)
	assert.Equal(t, gameId, game.GameId)
}

func TestInvite( t *testing.T) {
	// given

	// when

	// then

}

func TestInvited( t *testing.T) {
	// given

	// when

	// then

}

func TestAccept( t *testing.T) {
	// given

	// when

	// then

}

func TestAccepted( t *testing.T) {
	// given

	// when

	// then

}

func TestFire( t *testing.T) {
	// given

	// when

	// then

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
