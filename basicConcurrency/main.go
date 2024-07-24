package main

import (
	"fmt"
	"time"
)

func testFunc(channel chan int) {
	time.Sleep(time.Second * 2)
	channel <- 69
}

// func main() {
// 	// ch := make(chan int)
// 	// ch <- 69
// 	// will wait until someone reads from the channel
// 	// v := <-ch

// 	channel := make(chan int)
// 	go testFunc(channel)
// 	value := <-channel
// 	fmt.Println(value)
// }