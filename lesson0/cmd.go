package main

// Go provides a flag package supporting basic command-line flag parsing
import (
	"fmt"
	"os"
	"flag"
)

func main() {
	args := os.Args[1:]
	fmt.Printf("There is %d arguments\n", len(args))
	for i, v := range args {
		fmt.Printf("Argument %d is %s\n", i, v)
	}

	// Omitted flags automatically uses default values
	n := flag.Int("n", 10, "Max elements")
	b := flag.Bool("v", false, "Verbose")

	// It is possible to declare an option that uses an existing var
	var s string
	flag.StringVar(&s, "s", "", "Command")

	flag.Parse()

	fmt.Println("n:", *n)
	fmt.Println("f:", *b)
	fmt.Println("s:", s)
	fmt.Println("Tail arguments:", flag.Args())
}
