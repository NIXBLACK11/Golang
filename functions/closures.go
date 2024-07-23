package main

import "fmt"

func returnsConcatter() func(string) string {
	sent := ""
	return func(word string) string {
		sent += word + " "
		return sent
	}
}

// Just testing

func fullImplement() (func(string), func() string) {
	sent := ""

	return func(word string) {
		sent += word + " "
	}, func() string {
		return sent
	}
}

func main() {
	concatter := returnsConcatter()

	concatter("Hello")
	concatter("How")
	concatter("are")
	concatter("you")
	concatter("?")
	concatter("Hey")
	fmt.Println(concatter("sdvfsdvf"))

	concat, getSent := fullImplement()
	concat("Hello")
	concat("Sir")
	fmt.Println(getSent())
}