package ui

import (
	"github.com/MarcGrol/zeeslag/infra"
)

type UserInterface struct {
	toSelf *infra.ChannelsToSelf
}

func NewUserInterface(toSelf *infra.ChannelsToSelf) *UserInterface{
	return &UserInterface{
		toSelf: toSelf,
	}
}

func Listen() error {
	// TODO

	//_ := uuid.New().String()

	// waitfor ^C or command-line actions
	// readKeyboardInput
	// convert keyboard input into command
	// push command into service


	/*

		playerService.Command(
			core.InviteForGame{GameId: gameId}.ToPdu())
	 */
	return nil
}
