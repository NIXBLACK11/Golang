package main

import (
	"fmt"
	"time"
)

func asyncWork(message chan string) {
	for i:=1; i<=10; i++{
		time.Sleep(time.Second*1)
		donePercent := float64(i)/10*100
		fmt.Println("Task done: ", donePercent, "%")
	}
	message<-"done"
}

func messageOperate(messages chan string, message string) {
	time.Sleep(time.Second*1)
	messages<-message
	fmt.Println(message, " message sent!!")
}

func main() {
	/*
	// Basics of channel
	message:=make(chan string)

	go asyncWork(message)

	msg:=<-message
	fmt.Println(msg)

	// Buffered channels
	messages := make(chan string, 2)
	go messageOperate(messages, "First")
	go messageOperate(messages, "Second")
	go messageOperate(messages, "Third")
	go messageOperate(messages, "Fourth")
	go messageOperate(messages, "Fifth")
	go messageOperate(messages, "Sixth")

	fmt.Println(<-messages, " message received!!")
	fmt.Println(<-messages, " message received!!")
	fmt.Println(<-messages, " message received!!")

	// Select statement in channels
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1<-"one"
	} ()

	go func() {
		time.Sleep(1 * time.Second)
		c2<-"two"
	} ()

	// waits for one of the channels to return
	select {
	case msg1 := <- c1:
		fmt.Println("received", msg1)
	case msg2 := <- c2:
		fmt.Println("received", msg2)
	}

	// to get from both
	for i:=0;i<2;i++ {
		select {
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}

	//simply
	fmt.Println("received", <-c1)
	fmt.Println("received", <-c2)
	*/
}