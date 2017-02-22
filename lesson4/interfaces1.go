package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures
type Abser interface {
	Abs() float64
}

type Val float64

type Vertex struct {
	X, Y float64
}

func (val Val) Abs() float64 {
	if val < 0 {
		return float64(-val)
	}
	return float64(val)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var abser Abser
	val := Val(-math.Sqrt2)
	vertex := Vertex{3, 4}

	// Val implements Abser
	abser = val
	fmt.Println(abser.Abs())

	// *Vertex implements Abser
	abser = &vertex
	fmt.Println(abser.Abs())
}
