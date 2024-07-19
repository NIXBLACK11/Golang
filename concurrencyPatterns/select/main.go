package main

import (
	"fmt"
	"time"
)

var myChannel1 chan string
var myChannel2 chan string

func someFunc() {
	fmt.Println("Waiting for first!!")
	time.Sleep(time.Second * 2)
	myChannel1 <- "data1"
}

func anotherFunc() {
	fmt.Println("Waiting for second!!")
	time.Sleep(time.Second * 2)
	myChannel2 <- "data2"
}

func main() {
	myChannel1 = make(chan string)
	myChannel2 = make(chan string)

	go someFunc()
	go anotherFunc()
	

	// blocks until one of them works
	select {
	case msgMyChannel1 := <- myChannel1:
		fmt.Println(msgMyChannel1)
	case msgMyChannel2 := <- myChannel2:
		fmt.Println(msgMyChannel2)
	}
}