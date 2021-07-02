package api

import "github.com/MarcGrol/zeeslag/core"

type Peerer interface {
	InformPeer(event core.GameEventPdu) error
}
