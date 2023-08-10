package main

import (
	"fmt"
	"sync"
)

// Define type
type ArrayOfStr []string

// Print array of string and the loop number
func (aos ArrayOfStr) doPrint(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	fmt.Println(aos, number)
}

// Create interface based on method
type PrintData interface {
	doPrint(wg *sync.WaitGroup, number int)
}

func main() {
	// Declare and initiate variables
	var interface1 PrintData = ArrayOfStr{"bisa1", "bisa2", "bisa3"}
	var interface2 PrintData = ArrayOfStr{"coba1", "coba2", "coba3"}

	// Wait group so the function will wait for the go routines to be executed
	var wg sync.WaitGroup

	// Loop 4x for each interface
	for i := 1; i < 5; i++ {
		wg.Add(2)
		go interface1.doPrint(&wg, i)
		go interface2.doPrint(&wg, i)
	}

	// Wait the wait group to be done before exiting program
	wg.Wait()
}
