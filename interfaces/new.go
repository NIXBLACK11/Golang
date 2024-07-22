package main

import (
	"fmt"
	"math"
)


type shape interface {
	Area() float64
	Perimeter() float64
}

func printAreaPerimeter(sh shape) {
	msg := fmt.Sprintf("Area:%v\nPerimeter:%v", sh.Area(), sh.Perimeter())
	fmt.Println(msg)
}

type rect struct {
	Width int
	Height int
}

type circle struct {
	Radius int
}

func (r rect) Area() float64 {
	return float64(r.Height*r.Width)
}

func (r rect) Perimeter() float64 {
	return float64(2*(r.Height*r.Width))
}


func (c circle) Area() float64 {
	return float64(math.Pi * math.Pow(float64(c.Radius), 2.0))
}

func (c circle) Perimeter() float64 {
	return float64(2 * math.Pi * float64(c.Radius))
}

func main() {
	r := rect{
		10,
		10,
	}

	c := circle{
		10,
	}
	
	printAreaPerimeter(r)
	printAreaPerimeter(c)
}