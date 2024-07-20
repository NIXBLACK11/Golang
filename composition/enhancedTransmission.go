package main

import "fmt"

type EnhancedTransmission struct {
}

func (et EnhancedTransmission) ShiftUp() {
	fmt.Println("quick shifting up...")
}

func (et EnhancedTransmission) ShiftDown() {
	fmt.Println("quick shifting down...")
}