package main

import (
	"fmt"
	"time"
)

var myChannel chan string

func someFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("Waiting for this to complete!!")
	myChannel <- "data"
}

func main() {
	
	// communication is synchronous

	// This would create a unbuffered channel

	// myChannel := make(chan string)
	// go func ()  {
	// 	myChannel <- "data"
	// }()

	myChannel = make(chan string)

	go someFunc()

	msg := <-myChannel

	fmt.Println(msg)
}