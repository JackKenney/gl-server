package main

import (
	"fmt"

	"github.com/google/wire"
	server "github.com/jackkenney/gl-server/api"
)

func main() {
	fmt.Println("Hello, world.")
	s := initializeServer()
	fmt.Println("Server struct", s)
}

func initializeServer() server.GLServer {
	wire.Build()
	return server.Server{}
}
