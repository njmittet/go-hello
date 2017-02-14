package main

import (
	"fmt"
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

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(times(2, 3))
	fmt.Println(split(22))

	a, b := swap("First", "Last")
	fmt.Println(a, b)
}
