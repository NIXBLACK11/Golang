package main

import "fmt"

type Transmission struct {

}

func (t Transmission) ShiftUp() {
	fmt.Println("Shift up")
}

func (t Transmission) ShiftDown() {
	fmt.Println("Shifting down")
}