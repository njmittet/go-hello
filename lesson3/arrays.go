package main

import (
	"fmt"
)

// An array has a fixed size.
func basic() {
	var a [2]string
	a[0] = "Go"
	a[1] = "Hello"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// A slice is dynamically-sized,
func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var slice []int = primes[1:4]
	fmt.Println(slice)
}

// A slice just describes a section of an underlying array
func view() {
	names := []string{
		"John",
		"Ringo",
		"Paul",
		"George",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// Changing the elements of a slice modifies underlying array
	b[0] = "Unknown"
	fmt.Println(a, b)
	fmt.Println(names)
}

func main() {
	basic()
	slice()
	view()
}
