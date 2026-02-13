package main

import (
	"fmt"
	"net"
)

func main() {

	serverAddr := "10.100.23.11:34933"

	fmt.Println("Connected to ", serverAddr)

	//establishing listener connection
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening to ", serverAddr)

	for {
		_, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting, ", err)
			continue
		}
	}
}
