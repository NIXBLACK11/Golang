package main

import (
	"fmt"
)

func testFunc(val rune) {
	fmt.Printf("%#U\n", val)
}

func main() {
	s:="hello"
	for i, val := range s {
		fmt.Printf("index: %d, value: %#U\n", i, val)
	}

	s = "dsfkv dgffdgjfdg fd gfdgfdg"

	for _, val := range s {
		testFunc(val)
	}
}