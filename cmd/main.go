package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()

	fmt.Println("Hello, world.")
	server, err := InjectServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Run(":3000")
}
