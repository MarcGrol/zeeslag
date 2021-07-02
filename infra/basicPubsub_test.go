package infra

import (
	"fmt"
	"testing"

	"github.com/MarcGrol/zeeslag/core"
	"github.com/stretchr/testify/assert"
)

type eventSubsriptionHandler struct {
	publications []string
}

func (eh *eventSubsriptionHandler) OnEventPublished(topic string, event core.GameEventPdu) error {
	eh.publications = append(eh.publications, fmt.Sprintf("%s:%s", topic, event.EventType))
	return nil
}

func TestSubscribe(t *testing.T) {
	ps := NewBasicPubsub()
	eh := &eventSubsriptionHandler{
		publications: []string{},
	}
	err := ps.Subscribe("my_topic", eh)
	assert.NoError(t, err)

	err = ps.Publish("my_topic", core.GameInvitationSent{}.ToPdu())
	assert.NoError(t, err)

	err = ps.Publish("my_topic", core.GameAccepted{}.ToPdu())
	assert.NoError(t, err)

	err = ps.Publish("other_topic", core.GameInvitationSent{}.ToPdu())
	assert.NoError(t, err)

	assert.NoError(t, err)
	assert.Len(t, eh.publications, 2)
}
