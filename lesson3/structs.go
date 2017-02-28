package main

import (
	"fmt"
	"sync"
)

type Vertex struct {
	X int
	Y int
}

type Square struct {
	X, Y int
}

// A struct literal denotes a newly allocated struct value by using names fields
var (
	s1 = Square{2, 4}  // Has as type Square
	p1 = &Square{1, 2} // Has as type *Square

	s2 = Square{X: 1} // Y: 0
	s3 = Square{}     // X: 0, Y: 0
)

// Declare and initialize struct members using literals
var Fetched = struct {
	Items map[string]error
}{Items: make(map[string]error)}

// A struct is a collection of fields
func main() {
	fmt.Println(Vertex{12, 24})
	v := Vertex{21, 42}
	fmt.Println(v.X, v.Y)

	// Struct fields can be accessed through a struct pointer
	p := &v
	p.X = 1e9
	fmt.Println(v.X)

	fmt.Println(s1, p1, s2, s3)
	t := *p1
	fmt.Println(t.X)
}
