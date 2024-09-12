---
title: Day5. 輸入與輸出
tags: [鐵人賽, Golang]

---

# Day5. 輸入與輸出
今天來講一下怎麼讓你的Go程式透過`fmt`這個package進行輸入與輸出的格式化
> fmt 套件的主要功能是處理格式化的輸出和輸入，包括資料格式處理、字串解析和數值轉換等功能。通常用來記錄訊息、檔案處理等等

## 輸入(stdin)
如果你之前寫過C語言的話，你一定會用到scanf這個函式，如果沒有的話，簡單來說就是可以把你在Console中的輸入傳到指定的變數(如python的`input()`)。
不過輸入在Go語言中細分成了三個Function，分別為Scan, Scanln, Scanf，使用方式如下
- `func Scan(a ...any) (n int, err error)`
    - 以空格或換行分割，會一直讀取直到所有參數都成功賦值
    - 賦值期間出現錯誤，則結束賦值並出現錯誤(可忽略錯誤)
    - 換行等價於空格
- `func Scanf(format string, a ...any) (n int, err error)`
    - 透過格式化字串指定輸入格式，並在後面指定接受值的變數指標
- `func Scanln(a ...any) (n int, err error)`
    - 類似`Scan`，但是讀取到換行就結束讀取
```go
func main(){
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Scanf("a=%d b=%d", &a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Scanln(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}
```

## 輸出(stdout)
每個參數間輸出時以空格分隔(等價於python的`print(*a, sep=' ')`)
- `func Print(a ...any) (n int, err error)`
    - 輸出參數，結尾不換行
- `func Printf(format string, a ...any) (n int, err error)`
    - 設定格式化輸出後應用參數並輸出格式化後的字串
- `func Println(a ...any) (n int, err error)`
    - 類似`Print`，不過結尾預設換行


## 其他function
除了針對Console的輸入/輸出(stdin, stdout)以外，Go還支援直接對字串或檔案進行類似操作，只需要在前面指定你的目標檔案或字串，然後像用Scan或Print一樣就行
- 上面的`Scan`跟`Print`相關指定其實本質上都是透過`Fscan`或`Fprint`並指定檔案為`stdin`或`stdout`做到的
```go
// this is the source code of go

// src/fmt/print.go
func Print(a ...any) (n int, err error) {
	return Fprint(os.Stdout, a...)
}

func Println(a ...any) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

func Printf(format string, a ...any) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

// src/fmt/scan.go
func Scanf(format string, a ...any) (n int, err error) {
	return Fscanf(os.Stdin, format, a...)
}

func Scanln(a ...any) (n int, err error) {
	return Fscanln(os.Stdin, a...)
}

func Scan(a ...any) (n int, err error) {
	return Fscan(os.Stdin, a...)
}
```
### 檔案操作
#### Read
- `func Fscan(r io.Reader, a ...any) (n int, err error)`
- `func Fscanf(r io.Reader, format string, a ...any) (n int, err error)`
- `func Fscanln(r io.Reader, a ...any) (n int, err error)`
#### Write
- `func Fprint(w io.Writer, a ...any) (n int, err error)`
- `func Fprintf(w io.Writer, format string, a ...any) (n int, err error)`
- `func Fprintln(w io.Writer, a ...any) (n int, err error)`
### 字串操作
使用方式跟Scan或Print一樣，不過Sprint會回傳結果為字串
- `func Sscan(a ...any) (n int, err error)`
- `func Sscanf(format string, a ...any) (n int, err error)`
- `func Sscanln(a ...any) (n int, err error)`
#### Write
- `func Sprint(a ...any) string`
- `func Sprintf(format string, a ...any) string`
- `func Sprintln(a ...any) string`
## 參數說明
- `format(string)`: 指定輸入格式
- `n(int)`: 成功接收到的變數數量
- `...any`: 可以接受0~多個任意型態變數輸入，或將變數解包(可以理解成python的`*args`)
- `error`: 錯誤資訊 

## 格式化參數
在Go語言中稱為`verbs`，用於fmt包中的格式化函數（如`fmt.Printf`、`fmt.Sprintf`等）
對於格式化function`(Printf, Sprintf, and Fprintf)`，可以透過在format string中設定index決定使用第幾個參數
> For example,

```
fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
```
> will yield "22 11", while
```
fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
```
> equivalent to
```
fmt.Sprintf("%6.2f", 12.0)
```
> will yield " 12.00". Because an explicit index affects subsequent verbs, this notation can be used to print the same values multiple times by resetting the index for the first argument to be repeated:

```
fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
```
> will yield "16 17 0x10 0x11"[name=Golang Documentation]

### 基礎類型
- %v: 輸出該型態的預設格式
    - 對於`struct`則是輸出該struct中的所有fields的value
    - 對於`struct`使用`%+v`可以把輸出加上name
- %#v: 將值以Go語法表示
- %T: 將型別以Go語法表示
- %%: 百分符號(轉義後)
- %p: 記憶體位置(對於slice, pointer)
### 布林型態
- %t: 輸出true或false
### 整數型態
- %b	表示為二進位（不包含0b）
- %c	表示為Unicode字元
- %d	表示為十進位
- %o	表示為八進位
- %O	表示為八進位(包含0o)
- %q	類似%c, 不過會用單引號包住，且必要時會安全的轉義
- %x	表示為小寫十六進位
- %X	表示為大寫十六進位
- %U	表示為Unicode
### 浮點數與複數
- %b    轉換為IEEE754，Pow2的科學表示法
- %e    Pow10科學表示法，與%e只差在輸出的大小寫
- %f    表示為float

-----------------
那麼今天的文章就到這告一段落，如果我的文章有任何地方有錯誤請在留言區反應
明天將會介紹Go語言的條件與跳轉
![time](https://i.imgur.com/Hk7po4w.gif)
# REF
- https://openhome.cc/Gossip/Go/StdOutInErr.html
- https://pkg.go.dev/fmt
- https://easonwang.gitbook.io/golang/ji-ben-yu-fa/fmt
- https://vocus.cc/article/6502d147fd89780001f4e54d
- https://ithelp.ithome.com.tw/articles/10223934