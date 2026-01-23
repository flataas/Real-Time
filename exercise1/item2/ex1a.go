package main

import (
	"fmt"
	"sync"
)

var i int

func increaser(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1000000; j > 0; j-- {
		i++
	}
}

func decreaser(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1000000; j > 0; j-- {
		i--
	}
}

func main() {

	var wg sync.WaitGroup

	i = 0
	fmt.Println("Counter starts at", i)

	wg.Add(1)
	go increaser(1, &wg)
	wg.Add(1)
	go decreaser(2, &wg)

	wg.Wait()
	fmt.Println("All workers finnished.\nCounter ends at", i)

}
