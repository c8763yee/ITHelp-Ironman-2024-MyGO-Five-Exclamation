package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// STEP 2：定義能夠屬於 Error Interface 的方法
func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func checkUserNameExistNew(username string) (bool, error) {
	if username == "bar" {
		return true, errors.New("username bar is already exist")
	}

	return false, nil
}
func checkUserNameExistFmt(username string) (bool, error) {
	if username == "foo" {
		// 搭配 fmt.Errorf 來產生錯誤訊息
		return true, fmt.Errorf("username %s is already existed", username)
	}

	return false, nil
}

func raiseError(err error) {
	if err == nil {
		return
	}
	switch t := err.(type) {
	case MyError:
		fmt.Println("MyError", t)
	default:
		fmt.Println("default", t)

	}
}

func main() {
	if _, err := checkUserNameExistNew("bar"); err != nil {
		fmt.Println(err)
	}
	if _, err := checkUserNameExistFmt("bar"); err != nil {
		fmt.Println(err)
	}
}
