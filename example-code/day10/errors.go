package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}
	println(f.Name())
}

// do something with the open *File f}

// 在 golang 中 error 型態底層是 interface，這個 error 裡面有 Error() 這個 function，並且會回傳 string
// type error interface {
// 	Error() string
// }
