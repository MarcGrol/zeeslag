package infra

import (
	"github.com/MarcGrol/zeeslag/api"
	"github.com/MarcGrol/zeeslag/core"
)

type basicPubsub struct {
	subscriptions map[string][]api.Subscriber
}

func NewBasicPubsub() api.Publisher {
	return &basicPubsub{
		subscriptions: map[string][]api.Subscriber{},
	}
}

func (p *basicPubsub) Subscribe(topic string, subscriber api.Subscriber) error {
	subscriptions, found := p.subscriptions[topic]
	if !found {
		subscriptions = []api.Subscriber{}
	}

	subscriptions = append(subscriptions, subscriber)
	p.subscriptions[topic] = subscriptions

	return nil
}

func (p *basicPubsub) Publish(topic string, event core.GameEventPdu) error {
	for topic, subscriptions := range p.subscriptions {
		for _, subscription := range subscriptions {
			err := subscription.OnEventPublished(topic, event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
