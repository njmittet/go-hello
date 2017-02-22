package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures
type Abser interface {
	Abs() float64
}

// Interfaces are implemented implicitly
type Printer interface {
	Print()
}

type Value struct {
	V string
}

type T struct {
	S string
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

// This method means type MyStruct implements Printer
func (value Value) Print() {
	fmt.Println(value.V)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	// A type implements an interface by implementing its methods
	var printer Printer = Value{"Hello"}
	printer.Print()

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// MyFloat implements Abser
	a = f
	fmt.Println(a.Abs())

	// *Vertex implements Abser
	a = &v
	fmt.Println(a.Abs())
}
