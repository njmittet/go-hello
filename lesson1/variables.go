package main

import (
	"fmt"
)

// Declares a global list of variables
var a, b bool

// Variables with initializers
var d, e int = 1, 2

func main() {
	var i int

	// Short assignment statement with implicit type
	f := 2
	fmt.Println(i, a, b)
	fmt.Println(d, e, f)

	// If an initializer is present, the type can be omitted
	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)

	// If an initializer is present, the type can be omitted
	var x, y, z = true, false, "Test"
	fmt.Println(x, y, z)
}
