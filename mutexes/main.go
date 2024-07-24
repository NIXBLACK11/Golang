package main

import (
	"fmt"
	"sync"
	"time"
)

var mux sync.Mutex

func protected(callerId string) {
	mux.Lock()
	fmt.Println("Called by :"+ callerId + "\n")
	time.Sleep(time.Second*2)
	defer mux.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		protected("a")
	}()

	go func() {
		defer wg.Done()
		protected("b")
	}()

	wg.Wait()
}