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

// A slice literal is like an array literal without the length
func literals() {
	q := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(s)
}

func main() {
	basic()
	slice()
	view()
	literals()
}
