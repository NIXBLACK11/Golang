package main

type car struct {
	Make string
	Height int
	Wheel wheel
}

type wheel struct {
	Radius int
}

// func main() {
// 	myCar := car{}
// 	myCar.Wheel.Radius = 5
// }