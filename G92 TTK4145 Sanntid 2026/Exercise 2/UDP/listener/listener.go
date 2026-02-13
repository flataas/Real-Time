package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	var listeningPort string = ":20022"

	// Establishing listenin connection for udp on port 30000
	conn, err := net.ListenPacket("udp", listeningPort)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected!Listening to port", listeningPort)

	//buffer for receiving messages
	buffer := make([]byte, 1024)

	for {
		// assigning number of characters and sender address, and writing received message to buffer
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading", err)
			continue
		}

		//Printing received message
		fmt.Printf("Received from %s: %s\n", addr, buffer[:n])

		time.Sleep(time.Second)

	}

}
