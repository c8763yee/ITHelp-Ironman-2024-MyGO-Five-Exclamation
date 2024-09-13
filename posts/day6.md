---
title: 'Day6. if-else, switch case, goto'
tags: [鐵人賽, Golang]

---

# Day6. 條件與跳轉
今天來講Go語言如何處理條件與跳轉
## if-else
基本上使用跟C語言的if-else一樣，不過條件式外面可以不用加括號。
- 跟function還有之後會講到的for loop一樣，if-else的大括號不可換行
```go
x, y, z := 1, 2, 3
if x < y {
	return x
} else if (x > z) {
	return z
} else {
	return y
}
```

Go語言中if-else的其中一個特點是: 你可以在前面加上宣告變數(只能使用簡易宣告`:=`)，然後在這個if-else區塊中使用該變數(離開後就不可使用該變數)。
我自己觀察起來這個在if中宣告主要會用於錯誤處理(像是: `if (n, err:= fmt.Scanf(fmt_str, smth, smth2, ...))`)

```go
func did_something_happen() string {
	 return "nothing happened"
}
func main() {
    if (val:= did_something_happened(); val != "Nothing Happened"){
        println("SOMETHING HAPPENED")	
    }else{
        println("nothing")
    }
    // val 無法在它被宣告的if-else區塊外面使用
}

```
![Image](https://i.imgur.com/4SRlxSm.png)
# 蛤
![Image](https://i.imgur.com/5UZiDv8.gif)

經常寫C的人可能會覺得這個code沒有問題，但這樣的寫法在Go是大忌，原因在於:
如果每行最後是identifier、簡單表達式(數字或字串常數)、或是以下其中一個關鍵字
```go
break continue fallthrough return ++ -- ) }
```
那麼編譯時會在這些**表示結尾的符號**後面隱性加上分號。
> Like C, Go's formal grammar uses semicolons to terminate statements, but unlike in C, those semicolons do not appear in the source. Instead the lexer uses a simple rule to insert semicolons automatically as it scans, so the input text is mostly free of them.
> The rule is this. If the last token before a newline is an identifier (which includes words like int and float64), a basic literal such as a number or string constant, or one of the tokens
> the lexer always inserts a semicolon after the token. This could be summarized as, “if the newline comes after a token that could end a statement, insert a semicolon”.
> 
因此上面的code在編譯器會變成
```go
if (val:= did_something_happened(); val != "Nothing Happened");{
    // other code
}

// 或是
if (val:= did_something_happened(); val != "Nothing Happened");
{
    // other code
}
```
## Switch
- switch可以理解為另一種形式的if-else，跟C/C++不太一樣的地方在
    1. 你不需要使用break來跳出這次條件檢測
    2. 你甚至可以switch後面不加任何條件然後在case使用條件判斷
    3. 你可以同時檢查多個case
- 然後switch跟if-else一樣可以在switch那一行中宣告變數。
```go 
switch tag {
	default:
		s3()
	case 0, 1, 2, 3:
		s1()
	case 4, 5, 6, 7:
		s2()
	}

	switch x := f(); { // 後面沒加條件預設為 "true"
	case x < 0:
		return -x
	default:
		return x
	}

	switch {
	case x < y:
		f1()
	case x < z:
		f2()
	case x == 4:
		f3()
	}
```
接下來我要介紹兩個在switch-case中很重要的關鍵字: `break`跟`fallthrough`
### break
由於Go語言的特性為如果符合case時會進行case處理，結束後直接離開switch-case。因此你可以不用像寫C一樣使用break來確保不會走到其他case。也因此break在Go語言主要用來提早離開switch-case
```go
command := ReadCommand()
argv := strings.Fields(command)
switch argv[0] {
case "echo":
    fmt.Print(argv[1:]...)
case "cat":
    if len(argv) <= 1 {
        fmt.Println("Usage: cat <filename>")
        break
    }
    PrintFile(argv[1])
default:
    fmt.Println("Unknown command; try 'echo' or 'cat'")
}
```

### fallthrough
fallthrough會強制進入下一層case中運行該case的行為。不過fallthrough只能放在case區塊的最後一行
```go
v := 42
switch v {
case 100:
    fmt.Println(100)
    fallthrough
case 42:
    fmt.Println(42)
    fallthrough
case 1:
    fmt.Println(1)
    fallthrough
default:
    fmt.Println("default")
}
```
輸出的結果為
```
42
1
default
```

## 以下code為錯誤示範
```go

switch {
case f():
    if g() {
        fallthrough // Does not work!
    }
    h()
default:
    error()
}
```
## goto
在Go語言中，goto語句允許無條件跳轉到程序中的某一行，這在某些情況下可以提高代碼的可讀性或是簡化一些繁瑣的結構。不過需要注意的是，濫用goto會導致代碼難以理解和維護，因此在使用時應謹慎。

- 基本用法
goto 語句的基本用法包括定義一個標籤，然後使用 goto 跳轉到該標籤處的代碼。標籤的定義形式為一個簡單的標識符，後面跟著冒號。
```go
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

```
-----------------
那麼今天的文章就到這告一段落，如果我的文章有任何地方有錯誤請在留言區反應
明天將會介紹Go語言的迴圈操作
![time](https://i.imgur.com/Hk7po4w.gif)
# REF
- https://openhome.cc/Gossip/Go/StdOutInErr.html
- https://pkg.go.dev/fmt
- https://easonwang.gitbook.io/golang/ji-ben-yu-fa/fmt
- https://vocus.cc/article/6502d147fd89780001f4e54d
- https://go.dev/tour/flowcontrol/9
- https://ithelp.ithome.com.tw/articles/10331545
- https://go.dev/wiki/Switch