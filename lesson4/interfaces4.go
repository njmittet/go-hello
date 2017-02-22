package main

import "fmt"

func zeroValue() {
	// An empty interface may hold values of any type
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "Hello"
	describe(i)
}

func typeAssertion() {
	// A type assertion provides access to an interface value's underlying concrete value
	var i interface{} = "Hello"

	// Asserts that the interface value i holds a string
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Panics, since the underlying value is a string
	// f = i.(float64)
}

// A type switch is a construct that permits several type assertions in series
func typeSwitch(i interface{}) {
	// A type switch compares the type of the value held by the underlying interface
	switch v := i.(type) {
	case int:
		fmt.Printf("%v * 2 is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("Unknown type %T!\n", v)
	}

}

// An interface that specifies zero methods is known as an empty interface
func main() {
	zeroValue()
	typeAssertion()
	typeSwitch(21)
	typeSwitch("Hello")
	typeSwitch(true)
}

// Empty interfaces are used by code that handles values of unknown type
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
