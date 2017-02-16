package main

import (
	"fmt"
	"math"
)

// The for loop is the only looping construct
func loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Printf("Sum is %v\n", sum)
}

// The init and post statement are optional
func compact() {
	sum := 1
	for ; sum < 500; {
		sum += sum
	}
	fmt.Printf("Sum is %v\n", sum)
}

// Dropping the semicolons creates a while loop
func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Printf("Sum is %v\n", sum)
}

// Basic if statement
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// The if statement can start with a short statement to execute before the condition.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// Variables declared inside if short statements are also available in the else block
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// v is not available here
	return lim
}

func main() {
	loop()
	compact()
	while()
	fmt.Println(
		sqrt(2),
		sqrt(-4),
	)
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
