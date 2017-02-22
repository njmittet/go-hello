package main

import (
	"fmt"
)

// Interfaces are implemented implicitly
type Printer interface {
	Print()
}

type Value struct {
	V string
}

// This method means type Value implements Printer
func (value Value) Print() {
	fmt.Println(value.V)
}

func main() {

	// A type implements an interface by implementing its methods
	var printer Printer = Value{"Hello"}
	printer.Print()
}
