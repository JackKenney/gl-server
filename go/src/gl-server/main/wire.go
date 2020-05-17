package main

import (
	"github.com/google/wire"
)

func InitializeServer() Server {
	wire.Build()
	return Server{}
}
