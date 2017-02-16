package main

import (
	"fmt"
	"runtime"
)

const (
	OS_X  string = "darwin"
	LINUX string = "linux"
)

func main() {
	fmt.Print("OS is ")
	switch os := runtime.GOOS; os {
	case OS_X:
		fmt.Println("macOS")
	case LINUX:
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}
