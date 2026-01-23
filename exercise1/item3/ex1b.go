package main

import (
	"fmt"
	"sync"
)

func incrementer(requests chan string, done chan struct{}) {

	for range 1000000 {
		requests <- "increment"
	}

	done <- struct{}{}
}

func decrementer(requests chan string, done chan struct{}) {

	for range 1000000 {
		requests <- "decrement"
	}

	done <- struct{}{}
}

func manager(requests chan string, done chan struct{}) {
	var i int = 0
	var active int = 2

	for {
		select {
		case req := <-requests:
			switch req {
			case "increment":
				i++
			case "decrement":
				i--
			}
		case <-done:
			active--
			if active == 0 {
				fmt.Println("Routines done! counter at ", i)

			}
		}
	}
}

func main() {

	var wg sync.WaitGroup

	requests := make(chan string)
	done := make(chan struct{})

	wg.Add(3)
	go manager(requests, done)
	go incrementer(requests, done)
	go decrementer(requests, done)

	wg.Wait()

}
