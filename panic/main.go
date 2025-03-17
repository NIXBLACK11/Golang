package main

import (
	"errors"
	"fmt"
)

func divideWithPanic(a, b int) float32 {
	if b == 0 {
		panic("Divide by zero non-recoverable")
	}
	return float32(a) / float32(b)
}

func divideWithError(a, b int) (float32, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero error")
	}
	return float32(a) / float32(b), nil
}

func main() {
	// Not a good practice 
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	result, err := divideWithError(4, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result from divideWithError:", result)
	}

	result, err = divideWithError(4, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result from divideWithError:", result)
	}

	result = divideWithPanic(4, 2)
	fmt.Println("Result from divideWithPanic:", result)

	result = divideWithPanic(4, 0)
	fmt.Println("Result from divideWithPanic:", result)
}