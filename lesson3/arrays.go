package main

import (
	"fmt"
)

func basic() {
	var a [2]string
	a[0] = "Go"
	a[1] = "Hello"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func main() {
	basic()
}
