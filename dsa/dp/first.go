package main

import (
	"fmt"
)

func main() {
	var fib func(n int) int
	var n int
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &n)

	// recursive
    fib = func(n int) int {
        if n <= 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }

	fmt.Println(fib(n))

	// memoised
	dp := make(map[int]int)
	fib = func (n int) int {
		if n<=2 { return n }
		if val, exists := dp[n]; exists { return val }
		dp[n] = fib(n-1) + fib(n-2)
		return dp[n]
	}

	fmt.Println(fib(n))

	// iterative
	a := 0
	b := 1
	c := 0
	for i:=1; i<=n; i++ {
		c = a+b
		a = b
		b = c
	}

	fmt.Println(c)
}