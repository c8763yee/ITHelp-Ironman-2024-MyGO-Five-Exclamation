package main

import "fmt"

func main() {
	i := 0

Here: // 標籤定義
	fmt.Println(i)
	i++
	if i < 5 {
		goto Here // 跳轉到標籤Here
	}
}
