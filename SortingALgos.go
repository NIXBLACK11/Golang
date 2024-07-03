package main

import (
	"fmt"
	"time"
)

func measureExecutionTime(name string, f func(arr []int, n int), arr []int, n int) {
	start := time.Now()
	f(arr, n)
	elapsed := time.Since(start)
	fmt.Printf("\n%s execution took %s\n", name, elapsed)
}

func printArr(arr []int, n int) {
	fmt.Println("The sorted array:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println("\n")
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func bubbleSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

func selectionSort(arr []int, n int) {
	for i:=0; i<n; i++ {
		min:= i
		for j:=i; j<n; j++ {
			if arr[min]>arr[j] {
				min = j
			}
		}
		swap(arr, i, min)
	}
}

func main() {
	var n , algo int
	fmt.Println("Enter the size of the array:")
	fmt.Scan(&n)

	fmt.Println("Enter the values:")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	
	for {
		fmt.Println("Enter the sorting algorithm to apply:\n1 for bubble sort\n2 for selection sort\n10 to exit!")
		fmt.Scan(&algo)
		switch algo {
		case 1:
			measureExecutionTime("Bubble Sort", bubbleSort, arr, n)
		case 2:
			measureExecutionTime("Selection Sort", selectionSort, arr, n)
		case 10:
			break
		default:
			fmt.Println("Invalid input")
		}
		printArr(arr, n)
	}
}
