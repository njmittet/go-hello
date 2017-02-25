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

// A goroutine is a lightweight thread managed by the Go runtime
func basic() {
	go say("world")
	say("hello")
}

func main() {
	basic()
}
