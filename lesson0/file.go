package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
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

func cmdRead() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("You wrote %s", text)
}

func cmdScan() {
	fmt.Println("Enter text: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	fmt.Printf("You wrote %s", scan.Text())
}

func main() {
	readFile()
	readBytes()
	cmdRead()
	cmdScan()
}
