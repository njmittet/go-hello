package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

// M method for pointer to struct T
func (t *T) M() {
	fmt.Println(t.S)
}

// M method for float F
func (f F) M() {
	fmt.Println(f)
}

// Under the covers, interface values can be thought of as a tuple of a value and a concrete type
// An interface value holds a value of a specific underlying concrete type
func main() {
	// Calling a method on an interface executes the method of the same name on its underlying type
	var i I

	// Instantiate  a struct T and assign the address to a variable directly
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
