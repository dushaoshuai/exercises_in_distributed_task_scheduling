//go:build wireinject

package main

import "github.com/google/wire"

func initializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
