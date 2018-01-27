package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const port = "8080"

// RunHost accepts an IP to start listening on (port 8080)
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("error: ", listenErr)
	}
	fmt.Println("Listening on ", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}
	fmt.Println("New connection accepted")

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Println("Message received: ", message)
}

// RunGuest accepts an IP address to connect to (using port 8080)
func RunGuest(ip string) {

}
