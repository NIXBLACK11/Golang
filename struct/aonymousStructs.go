package main

import "fmt"

type rect struct {
	Height int
	Width int
}

func (r rect) Area() int {
	return r.Height*r.Width
}

func main() {
	// myCar := struct {
	// 	Height string
	// 	Model string
	// } {
	// 	Height: "150cm",
	// 	Model: "X",
	// }
	// myCar.Height = "160cm"

	r := rect{
		Width: 150,
		Height: 160,
	}

	fmt.Println(r.Area())
}