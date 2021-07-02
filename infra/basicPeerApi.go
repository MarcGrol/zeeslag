package infra

import (
	"github.com/MarcGrol/zeeslag/api"
	"github.com/MarcGrol/zeeslag/core"
)

type basicPeer struct {
	baseUrl string
}

func NewBasicPeer(baseUrl string) api.Peerer {
	return &basicPeer{
		baseUrl: baseUrl,
	}
}

func (p *basicPeer) InformPeer(event core.GameEventPdu) error {
	return nil
}
