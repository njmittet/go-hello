package main

import (
	"fmt"
	"math"
)

// Go does not have classes, but you can define methods on types
type Coordinate struct {
	X, Y float64
}

// A method is a function with a special Calc argument
func (c Coordinate) Calc() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
}

// Calc rewritten as normal function
func Sqrt(c Coordinate) float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
}

type MyFloat float64

// Methods can be declared on non-struct types
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Coordinate{3, 4}
	fmt.Println(v.Calc())
	fmt.Println(Sqrt(v))

	// Methods can only be declared on types defined in the same package as the method
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
