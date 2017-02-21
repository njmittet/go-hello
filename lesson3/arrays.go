package main

import (
	"fmt"
	"strings"
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
	q := []int{2, 3, 5, 7, 11, 13}
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

//When slicing, the high or low bounds can be omitted
func defaults() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s[1:4])
	fmt.Println(s[:4])
	fmt.Println(s[1:])
}

// A slice has both a length and a capacity
func lencap() {
	s := []int{2, 3, 5, 7, 11, 13}

	// Length is the number of elements the slice contains
	fmt.Println(len(s))

	// Capacity is the number of elements in the underlying array, counting from the first element in the slice
	fmt.Println(cap(s))

	a := s[:0]
	printSlice(a)

	b := s[:4]
	printSlice(b)

	c := s[2:]
	printSlice(c)
}

// The zero value of a slice is nil
func nilSlice() {
	// A nil slice has a length and capacity of 0 and has no underlying array
	var s []int
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Slices can be made dynamically using the make function
func makeSlice() {
	// Allocates a zeroed array and returns a slice
	a := make([]int, 5)
	printSlice(a)

	// Pass a third argument to make to specify capacity
	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

// Slices can contain any type, including other slices
func twoDimensional() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func extend() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	// The backing array is exended if capasity is to small
	s = append(s, 2, 3, 4)
	printSlice(s)
}

// The range form of the for loop iterates over a slice or map
func ranges() {
	s := []int{2, 3, 5, 7, 11, 13}

	// Two values is returned, the index and (a copy of) the value
	for i, v := range s {
		fmt.Printf("%d: %d\n", i, v)
	}

	// If you only want the index, drop the ", value" entirely.
	for i := range s {
		s[i] = 1 << uint(i)
	}

	// The index or value can be skipped by assigning to _
	for _, n := range s {
		fmt.Printf("%d\n", n)
	}
}

func main() {
	basic()
	slice()
	view()
	literals()
	defaults()
	lencap()
	nilSlice()
	makeSlice()
	twoDimensional()
	extend()
	ranges()
}
