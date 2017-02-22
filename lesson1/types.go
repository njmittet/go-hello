package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Variables can be declared in blocks
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

const Pi = 3.14

func types() {
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v\n", z, z)

	// Type inference
	t := "Test"
	fmt.Printf("Type of t is %T\n", t)
}

// Undeclared variables is given their default zero value
func defaultValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// Assignment between different type requires explicit conversion
func typeConversion() {
	var i, j int = 3, 4
	var f float64 = math.Sqrt(float64(i*i + j*j))
	var u uint = uint(f)
	fmt.Println(i, f, u)
}

func toInt(x int) int {
	return x*10 + 1
}

func toFloat(x float64) float64 {
	return x * 0.1
}

func constants() {
	// Constants are declared using the const keyword, and can not use the := syntax
	fmt.Println("Happy", Pi, "Day")

	// Numeric constants are high-precision values
	fmt.Println(Small)
	fmt.Println(toInt(Small))
	fmt.Println(toFloat(Small))
	fmt.Println(toFloat(Big))
}

func main() {
	types()
	defaultValues()
	typeConversion()
	constants()
}
