package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	OS_X  string = "darwin"
	LINUX string = "linux"
)

func basic() {
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

func weekend() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 2:
		fmt.Println("Tomorrow.")
	case today + 3:
		fmt.Println("In two day.")
	default:
		fmt.Println("To far away.")

	}
}

// Switch without a condition is the same as switch true.
func alwaysTrue() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	case t.Hour() < 17:
		fmt.Println("After noon")
	default:
		fmt.Println("Night")
	}
}

func main() {
	basic()
	weekend()
	alwaysTrue()
}
