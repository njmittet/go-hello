package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("The time is %v\n", time.Now())
	fmt.Printf("Random number %v\n", rand.Intn(10))
	fmt.Printf("We have %g problems\n", math.Sqrt(7))
	fmt.Printf("Pi is %g\n", math.Pi)
}
