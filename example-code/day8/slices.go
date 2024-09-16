package main

import "unsafe"

type slice struct {
	array unsafe.Pointer //指向底層array的pointer
	len   int            //長度(目前slice中元素數量)
	cap   int            //容量(目前slice中做多能存多少)
}

var slice_ []int = make([]int, 0, 10)
var lengthOfSlice int = len(slice_)
var sliceOfSlice int = cap(slice_)

func testSubSlice() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6} // 中括號內不放任何物件才是slice

	// 從現有array 或slice建立slice
	newSlice := arr
	newSlice = arr[:]   // 以arr作為底層`array`
	newSlice = arr[4:7] // 以arr作為底層`array`，並選取其中的index (4~7)進行遮罩

	anotherNewSlice := newSlice // 兩個變數都是指向同一個記憶體位置
	// 也可以透過make([]T, length, capacity)來建立(如果你已經確定slice大小)
	madeSlice := make([]int, 0, 5)
	madeSlice = append(madeSlice, 1, 2, 3, 4, 5)
	madeSlice = append(madeSlice, anotherNewSlice...)

	copiedSlice := make([]int, 4)
	copy(copiedSlice, arr[2:6]) // 將arr[2:6]複製到copiedSlice
}
