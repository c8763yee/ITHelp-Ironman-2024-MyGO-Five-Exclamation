package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Go file handling!")
	if err != nil {
		log.Fatal(err)
	}
}
