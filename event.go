package common

import (
	"context"
	"io"
)

type EventPayload interface{}
type EventType int

const (
	EpochStarted   EventType = iota
	Voted          EventType = iota
	Proposed       EventType = iota
	BlockAdded     EventType = iota
	ViewChanged    EventType = iota
	Committed      EventType = iota
	BalanceUpdated EventType = iota
)

type Event struct {
	T EventType
	//proto message
	Payload EventPayload
}

type EventBus interface {
	FireEvent(event *Event)
	Run(ctx context.Context)
	Subscribe(id string, writer io.Writer)
	Unsubscribe(id string)
}

type NullBus struct {
}

func (n *NullBus) Subscribe(id string, writer io.Writer) {
}

func (n *NullBus) Unsubscribe(id string) {
}

func (n *NullBus) FireEvent(event *Event) {
	//log.Debug("Suppress event")
}
func (n *NullBus) Run(ctx context.Context) {
}
