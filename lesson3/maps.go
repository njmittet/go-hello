package main

import (
	"fmt"
)

type Coordinate struct {
	Lat, Long float64
}

// A map maps keys to values
var m map[string]Coordinate

func basicMap() {
	// The make function returns a initialized map of the given type
	m = make(map[string]Coordinate)
	m["Bell Labs"] = Coordinate{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}

func main() {
	basicMap()
}
