package main

import (
	"fmt"
	"io"
	"strings"
)

// The io.Reader interface, which represents the read end of a stream of data
func main() {
	// The io.Reader interface has a Read method, implemented by the Reader
	r := strings.NewReader("Hello, Reader")
	b := make([]byte, 8)
	for {
		// Read populates the byte slice with data and returns the number of bytes populated and an error value
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
