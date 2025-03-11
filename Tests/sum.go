package Tests

func Test(arr...int) {
	test(0, arr)
}

func test(idx int, arr []int) int {
	if idx==len(arr) {
		return 0
	}

	return arr[idx] + test(idx+1, arr)
}