package main

import "fmt"

type transmission interface {
	ShiftUp()
	ShiftDown()
}

type Convertable struct {
	Engine
	transmission
	SteeringWheel
}

func (t Convertable) ConvertTop() {
	fmt.Println("Converting top...")
}

func (t Convertable) ConvertDown() {
	fmt.Println("Converting down...")
}