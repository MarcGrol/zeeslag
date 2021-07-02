package infra

import (
	"github.com/MarcGrol/zeeslag/core"
	"log"
)

const channelBufferSize = 10

type ChannelBasedService struct {
	playerName      string
	channelsToSelf  *ChannelsToSelf
	service         core.Service
}

func NewPlayerService(playerName string, channelsToSelf *ChannelsToSelf,
	service core.Service) *ChannelBasedService {
	return &ChannelBasedService{
		playerName:      playerName,
		channelsToSelf:  channelsToSelf,
		service:         service,
	}
}

type ChannelsToSelf struct {
	FromPlayer chan core.GameCommandPdu
	ToPlayer   chan core.GameEventPdu
}

func NewChannelsToSelf() *ChannelsToSelf {
	return &ChannelsToSelf{
		FromPlayer: make(chan core.GameCommandPdu, channelBufferSize),
		ToPlayer:   make(chan core.GameEventPdu, channelBufferSize),
	}
}

func (p *ChannelBasedService) Listen() {
	// Listen for commands from user and events from other service
	for {
		select {

		case c := <-p.channelsToSelf.FromPlayer:
			err := p.service.OnCommand(c)
			if err != nil {
				log.Printf("Error processing command from user: %+v", err)
			}
		}
	}
}

func (p *ChannelBasedService) Command(command core.GameCommandPdu) {
	p.channelsToSelf.FromPlayer <- command
}
