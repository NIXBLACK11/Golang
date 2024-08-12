package main

import (
	"fmt"
)

func goFunc(s string) {
	fmt.Println(s)
}

func main() {
	go goFunc("hello")
	
	go func (s string) {
		fmt.Println(s)
	} ("second")
}