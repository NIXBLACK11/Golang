package main

import (
	"fmt"
)

type Stack []int

func (st *Stack) Push(val int) {
	*st = append(*st, val)
}

func (st *Stack) Pop() int {
	if st.IsEmpty() {
		return -1
	}

	val := (*st)[len(*st)-1]
	*st = (*st)[0:len(*st)-1]
	return val
}

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

// func main() {
// 	stack := Stack{}
// 	stack.Push(1)
// 	stack.Push(2)
// 	stack.Push(3)
// 	fmt.Println(stack.Pop())
// 	fmt.Println(stack.Pop())
// 	fmt.Println(stack.Pop())
// }