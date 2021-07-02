package api

import "github.com/MarcGrol/zeeslag/core"

type Subscriber interface {
	OnEventPublished(topic string, event core.GameEventPdu) error
}

type Publisher interface {
	Publish(topic string, event core.GameEventPdu) error
}

type PubSub interface {
	Subscribe(topic string, subscriber Subscriber)
	Publisher
}
