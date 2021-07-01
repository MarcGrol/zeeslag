package service

import (
	"github.com/MarcGrol/zeeslag/cmd"
)

type ChannelBasedPlayerService struct {
	playerName     string
	games          map[string]ZeeslagService
	commandChannel chan cmd.GameCommandPdu
}

func NewPlayerService(playerName string) *ChannelBasedPlayerService {
	return &ChannelBasedPlayerService{
		playerName:     playerName,
		games:          map[string]ZeeslagService{},
		commandChannel: make(chan cmd.GameCommandPdu),
	}
}

func  (p *ChannelBasedPlayerService)getGameOnUid(id string) (ZeeslagService, bool) {
	game, found := p.games[id]
	return game, found
}

func (p *ChannelBasedPlayerService) ListenInBackground() {
	go func() {
		// Listen for commmands
		for {
			<-p.commandChannel
		}
	}()
}

func (p *ChannelBasedPlayerService) Command(command cmd.GameCommandPdu) {
	p.commandChannel <- command
}

func (p *ChannelBasedPlayerService)onCommand(command	 cmd.GameCommandPdu) {
	// TODO delegate to service

}

