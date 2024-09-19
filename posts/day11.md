# Day11. 檔案處理

在Go語言中，檔案處理是一個非常常見的操作，無論是讀取檔案內容還是將資料寫入檔案，Go都提供了強大的標準庫來幫助我們完成這些任務。今天我們將來介紹如何在Go語言中進行檔案的讀寫操作。
## 開啟檔案
你可以使用`os.Open`或`os.OpenFile`來開啟檔案，不過正常開發情況下使用`os.Open`來讀取檔案或是`os.Create`來創建檔案就足夠了(雖然其實`os.Open`跟`os.Create`都是靠`os.OpenFile`來實現的)

至於`os.OpenFile`，它是一個非常靈活的檔案打開方法，它能允許我們指定檔案的打開模式（如讀取、寫入、追加等）以及檔案權限
```go
/*
func Open(name string) (*File, error)
func OpenFile(name string, flag int, perm FileMode) (*File, error)
*/
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("example.txt") 
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
```


## 讀取檔案內容
Go語言提供了多種方法來讀取檔案內容，其中最常用的是使用`io/ioutil` 包中的`ReadFile`或是`os.ReadFile`。它接受檔案路徑作為參數，並以`[]byte`形式回傳檔案內容
> 附註： 1.16之後ioutil這個package已被棄用，其中的Function將會被改名並移動到`os`或`io`這兩個Package
> Deprecated: As of Go 1.16, the same functionality is now provided by package io or package os, and those implementations should be preferred in new code. See the specific function documentation for details. [ref](https://pkg.go.dev/io/ioutil#WriteFile)

你也可以先透過`os.Open`打開檔案後再使用`ReadAll(file)`來讀取檔案內容
```go
package main

import (
	"fmt"
	// "io/ioutil"
    "io"
	"log"
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
```


## 寫入內容至檔案

寫入檔案的操作與讀取相似，開啟檔案後可以透過該檔案物件的 `Write` 或 `WriteString` 方法來將資料寫入檔案。

```go
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
```


-----------------------------------------
那麼今天的文章就到此告一段落，如果我的文章有任何地方有錯誤請在留言區反應。
明天將會講在Go語言中如何作到類似物件導向的行為
![time](https://i.imgur.com/Hk7po4w.gif)

# 參考資料
- [如何在 Go 中操作檔案](https://medium.com/@alandev9751210/%E5%A6%82%E4%BD%95%E5%9C%A8-go-%E4%B8%AD%E6%93%8D%E4%BD%9C%E6%AA%94%E6%A1%88-50897de15e0d)
- [Go標準庫 - os包](https://pkg.go.dev/os)
- [Go標準庫 - ioutil包](https://pkg.go.dev/io/ioutil)
