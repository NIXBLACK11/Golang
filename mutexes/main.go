package main

import (
	"fmt"
	"sync"
	"time"
)

var mux sync.Mutex

/*
There are two ways in which we can unlock
1)
mux.Lock()
pperation
mux.Unlock()

2)
mux.Lock()
defer mux.Unlock()
operation
*/


func protected(callerId string) {
	mux.Lock()
	defer mux.Unlock()
	fmt.Println("Called by :"+ callerId + "\n")
	time.Sleep(time.Second*2)
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		protected("a")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		protected("b")
// 	}()

// 	wg.Wait()
// }