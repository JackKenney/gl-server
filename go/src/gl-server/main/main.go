package main

import (
	"fmt"

	"gl-server/server"
)

func main() {
	fmt.Println("Hello, world.")
	s := server.InitializeServer()

	s.Start()
}
