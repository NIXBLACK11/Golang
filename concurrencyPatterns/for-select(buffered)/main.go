package main

// import "fmt"

// func main() {
// 	// This creates a buffered channel

// 	// The communication is asynchronous

// 	charChannel := make(chan string, 3)
// 	chars := []string{"a", "b", "c"}


// 	for _, char := range chars {
// 		select {
// 		case charChannel <- char:
// 		}
// 	}

// 	close(charChannel)

// 	for result := range charChannel {
// 		fmt.Println(result)
// 	}
// }