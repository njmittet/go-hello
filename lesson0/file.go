package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	File = "files/example.txt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() {
	content, err := ioutil.ReadFile(File)
	check(err)
	fmt.Println(string(content))
}

func readBytes() {
	file, err := os.Open(File)
	defer file.Close()
	check(err)
	content := make([]byte, 5)
	count, err := file.Read(content)
	check(err)
	fmt.Printf("Read %d content: %s\n", count, string(content))
}

func main() {
	readFile()
	readBytes()

}
