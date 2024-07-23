package main

import "fmt"

func main() {
	x := 5
	y := 10
	z := &x

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*z)
} 