package service

import (
	"fmt"
	"github.com/MarcGrol/zeeslag/cmd"
	"log"
)

type ChannelBasedPlayerService struct {
	playerName     string
	commandChannel chan cmd.GameCommandPdu
	cmdService cmd.CommandService
}

func NewPlayerService(playerName string, cmdService cmd.CommandService) *ChannelBasedPlayerService {
	return &ChannelBasedPlayerService{
		playerName:     playerName,
		commandChannel: make(chan cmd.GameCommandPdu),
		cmdService: cmdService,
	}
}

func (p *ChannelBasedPlayerService) ListenInBackground() {
	go func() {
		// Listen for commmands
		for {
			c := <-p.commandChannel
			err := p.onCommand(c)
			if err != nil {
				log.Printf("Error processing command: %+v", err)
			}
		}
	}()
}

func (p *ChannelBasedPlayerService) Command(command cmd.GameCommandPdu) {
	p.commandChannel <- command
}

func (p *ChannelBasedPlayerService)onCommand(command cmd.GameCommandPdu) error {
	switch command.CommandType {
	case cmd.CommandType_Start:
		return p.cmdService.OnStart(*command.Start)
	case cmd.CommandType_Accept:
		return p.cmdService.OnAccept(*command.Accept)
	case cmd.CommandType_Fire:
		return p.cmdService.OnFire(*command.Fire)
	default:
		err := fmt.Errorf("Unrcognized command: %+v", command)
		log.Fatal(err.Error())
		return err
	}
}

