package main

import (
	"fmt"
)

type Queue []int

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	val := (*q)[0]
	*q = (*q)[1:len(*q)]
	return val
}

func (q *Queue) IsEmpty() bool {
	return len(*q)==0
}

func main() {
	queue := Queue{}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}