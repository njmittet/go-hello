package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// A goroutine is a lightweight thread managed by the Go runtime.
func basic() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// Send and receive blocks until the other side is ready.
	// This allows goroutines to synchronize without explicit locks
	c <- sum
}

// Channels are a typed conduit through for sending and receiving values with the channel operator: <-
func channels() {
	s := []int{7, 2, 8, -9, 4, 0}
	// Channels must be created before use
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

// Goroutines run in the same address space, so access to shared memory must be synchronized
func main() {
	basic()
	channels()
}
