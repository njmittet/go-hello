package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func times(x, y int) int {
	return x * y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Functions can return any number of values
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Printf("The time is %v\n", time.Now())
	fmt.Printf("Random number %v\n", rand.Intn(10))
	fmt.Printf("We have %g problems\n", math.Sqrt(7))
	fmt.Printf("Pi is %g\n", math.Pi)

	fmt.Println(add(2, 3))
	fmt.Println(times(2, 3))
	fmt.Println(split(22))

	a, b := swap("First", "Last")
	fmt.Println(a, b)
}
