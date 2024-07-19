In Go, `make` is used to create slices, maps, and channels, but not arrays. When you use `make(chan string)`, you are creating a channel that can carry values of type `string`, not an array.

Here’s a quick explanation of the difference between arrays, slices, and channels in Go:

- **Array**: A fixed-size sequence of elements of a specific type.
- **Slice**: A dynamically-sized, flexible view into the elements of an array.
- **Channel**: A communication mechanism that allows goroutines to send and receive values of a specified type.

Here is how you use `make` to create each of these:

### Creating a Slice
```go
// Creates a slice of integers with length 3
mySlice := make([]int, 3)
mySlice[0] = 1
mySlice[1] = 2
mySlice[2] = 3
```

### Creating a Map
```go
// Creates a map with string keys and integer values
myMap := make(map[string]int)
myMap["one"] = 1
myMap["two"] = 2
```

### Creating a Channel
```go
// Creates a channel that carries string values
myChannel := make(chan string)
```

### Creating an Array
Arrays do not use `make`; instead, they are declared with a fixed size:
```go
// Creates an array of integers with length 3
var myArray [3]int
myArray[0] = 1
myArray[1] = 2
myArray[2] = 3
```

In the context of your original code:

```go
package main

import "fmt"

// Define a global channel
var myChannel chan string

func someFunc() {
    myChannel <- "data"
}

func main() {
    // Initialize the global channel
    myChannel = make(chan string)
    
    // Start a goroutine to send data to the channel
    go someFunc()
    
    // Receive data from the channel
    msg := <-myChannel
    
    fmt.Println(msg)
}
```

- `myChannel = make(chan string)` initializes a channel that can carry strings.
- This is different from arrays and slices, which store collections of elements.

So, `make(chan string)` creates a channel, not an array of type `chan string`.