package main

import (
	"fmt"
)

// Deferred function calls are pushed onto a stack and executed in LIFO order when function returns
func counting() {
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")

}

// A defer statement defers the execution of a function until the surrounding function returns
func main() {
	// The deferred call's arguments are evaluated immediately
	defer fmt.Println("World!")
	fmt.Println("Hello")

	counting()
}
