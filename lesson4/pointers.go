package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	X, Y float64
}

// With a value receiver, the method operates on a copy of the original struct
func (c Coordinate) Abs() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
}

// Methods can be pointer receivers instead of value receivers
func (c *Coordinate) Scale(f float64) {
	c.X = c.X * f
	c.Y = c.Y * f
}

// Scale rewritten as function
func functionScale(c *Coordinate, f float64) {
	c.X = c.X * f
	c.Y = c.Y * f
}

func main() {
	c := Coordinate{3, 4}
	fmt.Println(c.Abs())

	// Methods with pointer receivers can modify the value of the struct they point to
	c.Scale(10)
	fmt.Println(c.Abs())

	functionScale(&c, 2)
	fmt.Println(c.Abs())
}
