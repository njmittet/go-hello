package main

import (
	"fmt"
)

// A pointer holds the memory address of a value.
func main() {
	i, j := 42, 2701
	p := &i         // Implicit pointer declaration
	fmt.Println(*p) // Dereferencing the pointer
	fmt.Println(p)  // Prints pointers memory address
	*p = 21
	fmt.Println(i) // Variable i has new value

	// Declare pointer to an int. The zero value is nil
	var q *int
	fmt.Println(q)
	q = &j
	fmt.Println(q)
}
