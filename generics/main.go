package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func sum[T Number](numbers []T) T {
	var result T

	for i := range numbers {
		result += numbers[i]
	} 

	return result
}

func getLast[T any](numbers []T) T {
	if len(numbers) == 0 {
		var zeroVal T
		return zeroVal
	}

	return numbers[len(numbers)-1]
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int32{1, 2, 3, 4, 5}
	c := []int64{1, 2, 3, 4, 5}
	d := []float32{1.1, 2.2, 3.3}
	e := []float64{1.1, 2.2, 3.3}

	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(sum(c))
	fmt.Println(sum(d))
	fmt.Println(sum(e))
}