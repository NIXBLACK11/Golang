package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMux sync.RWMutex

func readProtected(callerId string) {
	rwMux.RLock()
	defer rwMux.RUnlock()
	fmt.Println("Read lock acquired by: " + callerId)
	time.Sleep(2 * time.Second)
	fmt.Println("Read lock released by: " + callerId)
}

func writeProtected(callerId string) {
	rwMux.Lock()
	defer rwMux.Unlock()
	fmt.Println("Write lock acquired by: " + callerId)
	time.Sleep(2 * time.Second)
	fmt.Println("Write lock released by: " + callerId)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// Start a reader goroutine
	go func() {
		defer wg.Done()
		readProtected("Reader 1")
	}()

	// Start another reader goroutine
	go func() {
		defer wg.Done()
		readProtected("Reader 2")
	}()

	// Start a writer goroutine
	go func() {
		defer wg.Done()
		writeProtected("Writer 1")
	}()

	wg.Wait()
}
