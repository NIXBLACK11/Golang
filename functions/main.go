package main

import "fmt"

// first class functions
func add(x int, y int) int {
	return x+y
}

func mul(x int, y int) int {
	return x*y
}

// higher order functions
func aggregate(x int, y int, arithematic1 func(int, int) int, arithematic2 func(int, int) int) int {
	return arithematic1(x, arithematic2(x, y))
}

// A function that returns a function
func selfMath(operation func(int, int) int) func(int) int {
	return func(x int) int {
		return operation(x, x)
	}
}

// func main() {
// 	fmt.Println(aggregate(3, 4, add, mul))
// 	fmt.Println(aggregate(3, 4, mul, add))

// 	sqr := selfMath(mul)
// 	fmt.Println(sqr(2))
// }