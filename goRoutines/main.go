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

func main() {
	message:=make(chan string)

	go asyncWork(message)

	msg:=<-message
	fmt.Println(msg)
}