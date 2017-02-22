package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// Implements the Error method of the error interface defined by "fmt"
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"Did not work",
	}
}

// Go programs express error state with error values
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
