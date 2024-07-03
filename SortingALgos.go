package main

import (
	"fmt"
)

func printArr(arr[] int, n int) {
	for i:=0; i<n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}

func swap(arr[] int, i int, j int) {
	temp:= arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	var n int
	fmt.Println("Enter the size of the array:")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&arr[i])
	}

	bubbleSort
	printArr(arr, n)

}