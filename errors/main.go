package main

import (
	"errors"
	"fmt"
)

// Go does not have try catch block

func div(a float64, b float64) (float64, error) {
	if b==0 {
		return 0, errors.New("Divide by zero")
	}

	return a/b, nil
}

// func main() {
// 	res, err := div(2, 1)
// 	if err!=nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Printf("Result is: %v", res)
// }