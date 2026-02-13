package main

import (
	"fmt"
	"net"
)

func main() {

	serverAddr := "10.100.23.11:34933"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to ", serverAddr)

	message := "hello by tcp!\\0"

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sendinr: ", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}

	fmt.Printf("Response: %s\n", buffer[:n])

}
