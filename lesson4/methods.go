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


func main() {
	v := Coordinate{3, 4}
	fmt.Println(v.Calc())
	fmt.Println(Sqrt(v))
}
