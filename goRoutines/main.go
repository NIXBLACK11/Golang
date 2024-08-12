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
}