package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	var receiverAddr string = "10.100.23.11:20022"

	// Establishing connection for sending udp packets to specified address
	conn, err := net.Dial("udp", receiverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to address", receiverAddr)

	message := "hello!"

	for {

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending", err)
			return
		}

		fmt.Println("Sent: ", message)

		/*
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response", err)
			return
		}

		fmt.Println("Received response: ", buffer[:n])
		*/
		
		time.Sleep(time.Second)
	}

}
