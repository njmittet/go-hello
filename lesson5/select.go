package main

import (
	"fmt"
	"time"
)

// The select statement lets a goroutine wait on multiple communication operations
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// A select blocks until one of its cases can run
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func basicSelect() {
	c := make(chan int)
	quit := make(chan int)
	// Anonymous function
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// The default case in a select is run if no other case is ready
func defaultSelect() {
	tick := time.Tick(500 * time.Millisecond)
	quit := time.After(1500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-quit:
			fmt.Println("Quit")
			return
		default:
			fmt.Println(".")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func main() {
	basicSelect()
	defaultSelect()
}
