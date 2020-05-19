package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	fmt.Println("Hello, world.")
	server, err := InjectServer(wait)
	if err != nil {
		log.Fatal(err)
	}
	server.Run("127.0.0.1", "3000")
}
