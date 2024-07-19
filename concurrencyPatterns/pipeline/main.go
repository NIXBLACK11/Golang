package main

import "fmt"

func sliceToChannel(nums []int)  <-chan int {
	dataChannel := make(chan int)

	go func() {
		for _, n := range nums {
			dataChannel <- n
		}
		close(dataChannel)
	}()

	return dataChannel
}

func sq(in <-chan int) <-chan int {
	finalChannel := make(chan int)
	go func ()  {
		for n := range in {
			finalChannel<-n*n
		}
		close(finalChannel)
	}()

	return finalChannel
}

func main() {
	// each stage is synchronized that helps us range over a unbuffered channel, which is a single value channel

	// input
	nums := []int{2, 3, 4, 7 ,1}

	// stage 1
	dataChannel := sliceToChannel(nums)

	//stage2
	finalChannel := sq(dataChannel)

	//stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}