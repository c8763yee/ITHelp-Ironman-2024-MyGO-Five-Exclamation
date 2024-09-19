# Day10. 錯誤處理
今天來講在Go語言中如何進行錯誤處理。因為在Go語言中進行錯誤處理跟其他語言的差異很大，因此我需要單獨寫一篇文章介紹

## 簡介
當我們在開發過程中可能因為各種原因(如忘記分號，縮排錯誤或是邏輯錯誤)造成異常，這時程式語言通常會拋出各種錯誤，然後退出程式，這時就需要進行錯誤處理。
錯誤處理的主要目的是確保程式能夠在面臨預期外的情況時依然表現得穩健，而不會崩潰或產生不正確的結果。在大型專案中，良好的錯誤處理策略還能幫助快速定位問題，減少調試和維護的成本。

Go語言不像一些其他語言那樣有例外（exceptions）機制，而是採用了一種簡單而明確的錯誤處理方式，即通過返回值來傳遞錯誤。不過這對於從其他語言開始學習的人會非常不習慣，如`os.Open`這個用於開啟檔案的function它會回傳檔案物件與err(錯誤物件)，如果檔案能正常被打開則`err`為`nil`，反之則會將錯誤原因包為`Error`物件後回傳

## Usage
```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f

// 在 golang 中 error 型態底層是 interface，這個 error 裡面有 Error() 這個 function，並且會回傳 string
// type error interface {
//     Error() string
// }
```
如果要自訂Error物件的話則有以下幾種方式
1. 自己定義Error struct與interface
```go
// A Tour of Go，https://tour.golang.org/methods/19

// STEP 1：定義客製化的 error struct
type MyError struct {
	When time.Time
	What string
}

// STEP 2：定義能夠屬於 Error Interface 的方法
func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
```
然後在自己的Function中需要拋出錯誤的地方初始化struct並回傳就行

2. 使用`errors`套件中的`errors.New()`來建立error物件
```go
import (
  "fmt"
  "errors"
)

func checkUserNameExist(username string) (bool, error) {
	if username == "bar" {
		return true, errors.New("username bar is already exist")
	}

	return false, nil
}

func main() {
	if _, err := checkUserNameExist("bar"); err != nil {
		fmt.Println(err)
	}
}
```
3. 透過`fmt`套件的`err.Errorf`
```go
func checkUserNameExist(username string) (bool, error) {
	if username == "foo" {
    // 搭配 fmt.Errorf 來產生錯誤訊息
		return true, fmt.Errorf("username %s is already existed", username)
	}

	return false, nil
}
```

## 判斷錯誤類型
```go
// 可以透過`err.(Type) 檢查該錯誤的類別，常用於判斷錯誤是否為自訂錯誤
errObject, isMyError = err.(MyError)
```

--------------------
那麼今天的文章就到這告一段落，如果我的文章有任何地方有錯誤請在留言區反應
明天將會介紹Go語言的檔案處理
![time](https://i.imgur.com/Hk7po4w.gif)

# REF
- https://pjchender.dev/golang/error-handling/#%E4%BD%9C%E6%B3%95%E4%BD%BF%E7%94%A8-errorsnew