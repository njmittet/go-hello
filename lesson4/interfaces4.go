package main

import "fmt"

// An interface that specifies zero methods is known as an empty interface
func main() {

	// An empty interface may hold values of any type
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "Hello"
	describe(i)
}

// Empty interfaces are used by code that handles values of unknown type
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
