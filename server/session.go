package server

import "mekanicum/pkg/protocol"

type Session interface {
	ID() string
	Initialize()
	Initialized() bool
	Notifications() chan<- protocol.McpMessage
}
