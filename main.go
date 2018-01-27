package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
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
		runHost(connIP)
	} else {
		// index:   0       1
		// go run main.go <ip>
		connIP := os.Args[1]
		fmt.Println("Connecting to: ", connIP)
		runGuest(connIP)
	}
}

const port = "8080"

func runGuest(ip string) {

}

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("error: ", listenErr)
	}
	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Println("Message received: ", message)
}
