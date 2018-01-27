package main

import (
	"flag"
	"fmt"
	"os"

	"./lib"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Is this machine the host or a client?")
	flag.Parse()
	if isHost {
		// index:   0       1       2
		// go run main.go --listen <ip>
		connIP := os.Args[2]
		fmt.Println("Host listening on: ", connIP)
		lib.RunHost(connIP)
	} else {
		// index:   0       1
		// go run main.go <ip>
		connIP := os.Args[1]
		fmt.Println("Connecting to: ", connIP)
		lib.RunGuest(connIP)
	}
}
