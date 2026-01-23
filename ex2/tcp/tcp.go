package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func inviter(address string) {

	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := []byte("Connect to:10.24.208.228:8080\x00")

	buffer := make([]byte, 1024)

	for {

		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Error writign message: ", err)
			continue
		}

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error readign response", err)
			return
		}

		fmt.Printf("Response: %s\n", buffer[:n])

		time.Sleep(time.Second)
	}

}

func sendOnce(address, msg string) {

	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msgToSend := []byte(msg + "\x00")

	_, err = conn.Write(msgToSend)
	if err != nil {
		fmt.Println("Error writing message: ", err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error readign response", err)
		return
	}

	fmt.Printf("Response: %s\n", buffer[:n])
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected! ", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	for {

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading message from connection: ", err)
			return
		}

		received := strings.TrimSpace(string(buffer[:n]))

		fmt.Printf("Received from %s: %s\n", conn.RemoteAddr(), received)
	}

}

func accepter(port string) {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("TCP server listening on " + port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}

		go handleConnection(conn)

	}

}

func main() {

	ip, remotePort, myPort := "10.24.208.228", ":34933", ":8080"

	var wg sync.WaitGroup

	wg.Add(1)
	go accepter(myPort)
	time.Sleep(time.Second)

	sendOnce(ip+remotePort, "Connect to:"+ip+myPort)

	//wg.Add(1)
	//go inviter(ip + remotePort)

	wg.Wait()
}
