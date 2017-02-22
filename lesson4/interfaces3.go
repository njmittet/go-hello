package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	// The struct T is not instantiated, hence the underlying interface value is nil
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	var j I
	describe(j)
	// j.M() Causes runtime error since both the interface and the underlying value is nil
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
