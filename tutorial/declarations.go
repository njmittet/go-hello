package main

import (
	"fmt"
)

var a, b, c bool
var d, e int = 1, 2

func main() {
	var i int
	f := 2
	fmt.Println(i, a, b, c)
	fmt.Println(d, e, f)

	x, y, z := true, false, "Test"
	fmt.Println(x, y, z)
}
