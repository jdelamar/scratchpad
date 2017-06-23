package main

import (
	"fmt"
	"sync"
	"time"
)

var waiter sync.WaitGroup

func main() {
	// This is another change
	RunStuff()
}

func RunStuff() {
	fmt.Printf("Oh well, that started\n")

	waiter.Add(1)

	go doSomething()
	time.Sleep(10 * time.Second)

	fmt.Printf("Ah! it waited 10 seconds\n")

	waiter.Wait()

}

func doSomething() {

	defer waiter.Done()

	tick := 0

	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		tick++
		fmt.Printf("Ticked at %s\n", t.String())

		if tick > 5 {
			break
		}
	}
}
