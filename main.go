package main

import (
	"flag"
	"fmt"
)

func main() {
	var isHost bool
	flag.BoolVar(&isHost, "listen", false, "Is this machine the host or a client?")
	flag.Parse()
	if isHost {
		fmt.Println("This is the host")
	} else {
		fmt.Println("This is not the host")
	}
}
