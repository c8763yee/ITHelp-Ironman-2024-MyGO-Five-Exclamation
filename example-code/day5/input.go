package main

import "fmt"

func fileProcess() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Scanf("a=%d b=%d", &a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Scanln(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}
func Example() string {
	a := fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	a = fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	a = fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
	return a
}
