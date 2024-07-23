package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["john"] = 11
	ages["roy"] = 12

	delete(ages, "john")
	_, ok := ages["john"]

	if ok==false { fmt.Println("john does not exist") }


	// This results in extra code
	twoD := make(map[string]map[string]int)

	twoD["sidd"] = make(map[string]int)

	twoD["sidd"]["height"] = 120

	fmt.Println(twoD["sidd"]["height"])

	// Instead use structs
	type Key struct {
		Name string
		Param string
	}

	keyMap := make(map[Key]int)

	keyMap[Key{Name: "sidd", Param: "Height"}] = 120

	fmt.Println(keyMap[Key{"sidd", "Height"}])
}