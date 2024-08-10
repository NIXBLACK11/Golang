package main

import (
	"fmt"
)

func test(nums ...int) {
	fmt.Print(nums)
}

func main() {
	test(1, 2, 3, 4)
	arr := []int{1, 2, 3, 4, 5}
	test(arr...)
}