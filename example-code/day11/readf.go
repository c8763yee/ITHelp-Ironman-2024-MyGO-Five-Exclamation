package main

import (
	"fmt"
	// "io/ioutil"
	"io"
	"os"
)

func main() {
	file, err := os.Open("notebook.txt")
	// handle error
	defer file.Close()

	// # before Go 1.16
	// content, err := ioutil.ReadAll(file)
	// # after Go 1.16
	content, err := io.ReadAll(file)
	// handle error

	// read file via os.ReadFile
	anotherContent, err := os.ReadFile("notebook.txt")
	fmt.Println(string(content) == string(anotherContent)) // true
}
