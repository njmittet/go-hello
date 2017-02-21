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

func mapLiterals() {
	// Map literals are like struct literals, but the keys are required
	var m = map[string]Coordinate{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m["Google"])
}

func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("Value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("Value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("Value: ", m["Answer"])

	// Test that a key is present with a two-value assignment
	v, ok := m["Answer"]
	fmt.Println("Value: ", v, "Present: ", ok)
}

func main() {
	basicMap()
	mapLiterals()
	mutatingMaps()
}
