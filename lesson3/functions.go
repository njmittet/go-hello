package main

import (
	"fmt"
	"math"
)

// Takes a function as input
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Functions are first class citizens
func functionValue() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func main() {
	functionValue()
}
