package main

import "fmt"

func play_haruhikage(condition bool) bool {
	return condition == true
}

func main() {
	for i := 1; i < 10; i++ {
		println(i)
	}
	itersObj := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range itersObj {
		fmt.Printf("Print value: %d\n", v)
	}
	for i, v := range itersObj {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	for index := range itersObj {
		fmt.Printf("Index: %d\n", index)
	}
	for {
		fmt.Println("Infinite loop")
		break
	}

	for play_haruhikage(true) {
		go func() {}()
	}
}
