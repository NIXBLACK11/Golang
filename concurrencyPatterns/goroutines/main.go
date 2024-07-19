package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	// Asynchronous 

	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	time.Sleep(time.Second * 2)

	fmt.Println("hi")
}