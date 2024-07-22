package main

func main() {
	// Example array
	/*
		Fixed size contiguous memory locations
	*/
	var arr1 [10]int

	arr2 := [6]int{1, 2, 3, 4, 5, 6}

	// Example slice
	/*
		Built on top of slices and are dynamic
	*/
	var slic1 []int
	slic2 := arr2[1:2]
	slic3 := []int{1, 2, 3, 4}

	a1 := make([]int, 5)
	a2 := []int{}
}