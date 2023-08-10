package main

import (
	"fmt"
	"sync"
)

// Define struct
type ArrayOfStr struct {
	mu    sync.Mutex // To lock the object when being operated, so it can only be accessed by one goroutine
	aos   []string
	count int // The number of times the array is being looped
}

// Print array of string and the loop number
func (aos *ArrayOfStr) doPrint(wg *sync.WaitGroup) {
	// Announce that the function is done
	defer wg.Done()
	// Lock the aos resource so it cannot be accessed by another goroutine
	aos.mu.Lock()
	// Add the count
	aos.count++
	// Unlock the aos
	defer aos.mu.Unlock()
	// Print the string and the number of times the array is being printed
	fmt.Println(aos.aos, aos.count)
}

// Create interface based on method
type PrintData interface {
	doPrint(wg *sync.WaitGroup)
}

func main() {
	// Declare and initiate variables
	var interface1 PrintData = &ArrayOfStr{
		count: 0,
		aos:   []string{"bisa1", "bisa2", "bisa3"},
	}
	var interface2 PrintData = &ArrayOfStr{
		count: 0,
		aos:   []string{"coba1", "coba2", "coba3"},
	}

	// Wait group so the function will wait for the go routines to be executed
	var wg sync.WaitGroup

	// Loop 4x for each interface
	for i := 1; i < 5; i++ {
		wg.Add(2)
		go interface1.doPrint(&wg)
		go interface2.doPrint(&wg)
	}

	// Wait the wait group to be done before exiting program
	wg.Wait()
}
