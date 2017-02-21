package main

import (
	"fmt"
	"math"
)

// Takes a function as input
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Returns a function, creating a closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
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

// A closure is a function value that references variables from outside its body
func functionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*1),
		)
	}
}

func main() {
	functionValue()
	functionClosures()

	// Calling unassigned closure
	fmt.Println(func(i int) int { return 42 * i }(2))
}
