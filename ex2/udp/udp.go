package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func listener(port string) {
	conn, err := net.ListenPacket("udp", port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Listening on port ", port)

	buffer := make([]byte, 1024)

	for {

		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading message: ", err)
			continue
		}

		fmt.Printf("Received from %s : %s\n", addr, buffer[:n])

	}

}

func sender(addr string) {

	conn, err := net.Dial("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := "Hello using udp"

	for {

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message: ", err)
		}

		fmt.Println("Sent message!")

		time.Sleep(time.Second)
	}

}

func main() {

	var wg sync.WaitGroup

	ip, port := "10.24.208.228", ":20001"

	wg.Add(1)
	go listener(port)
	time.Sleep(time.Second)

	wg.Add(1)
	go sender(ip + port)

	wg.Wait()
}
