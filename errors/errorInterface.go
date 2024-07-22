package main

import "fmt"

type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("Error in user: %s", e.name)
}

func getData() error{
	return userError{name: "John"}
}