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

// A sender can close a channel to indicate that no more values will be sent
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Printf("Cap: %v, Len: %v\n", cap(c), len(c))
		x, y = y, x+y
	}
	// Closing is only necessary when the receiver must be told there are no more values, e.g loop termination
	close(c)
}

// Channels are a typed conduit through for sending and receiving values with the channel operator: <-
// Unbuffered channels block receivers until data is available and senders until a receiver is available
func channels() {
	s := []int{7, 2, 8, -9, 4, 0}
	// Channels must be created before use
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

// Channels can be buffered
// Buffered blocks a sender once the buffer fills up
func bufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
}

// Receivers can test if a channel has been closed by assigning a second parameter to the receive expression: after
func closeChannel() {
	c := make(chan int, 10)
	go fibonacci(cap(c)*2, c)
	time.Sleep(1 * time.Second)
	for i := range c {
		fmt.Println(i)
	}
}

// Goroutines run in the same address space, so access to shared memory must be synchronized
func main() {
	basic()
	channels()
	bufferedChannels()
	closeChannel()
}
