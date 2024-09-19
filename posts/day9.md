# Day9. 函式
## 函式(Function)
### 基礎概念
function可以接受輸入參數並返回輸出，可用於整理code成一個區塊或模組，讓你可以重複使用且不需要重新撰寫相同邏輯。
最常見的函數是`main()`，代表程式的入口(須配合`package main`使用才會呼叫到裡面的main)，在Python中的類似語句為
```python=
if __name__ == '__init__':
    main()
```
### 函式命名風格
- 函式名稱必須以字母開始
- 函式名稱只能包含字母、數字或下划線(`_`)
- 函式名稱區分大小寫
- 如果函式名稱由多個單詞組成，從第二個單詞開始的每個單詞首字母應大寫，例如`haruHiKage`
- 如果第一個字母為大寫，則可以被其他套件(package)引用，反之只能在同一套件內部使用
    - 可以理解為public method與private method
### 使用與定義方式
- 使用 `func` 關鍵字來定義函式
```go
func <funcName>(arg1 type, arg2 <type>, ...) <returnType>{
    // some of function
    return <returnValue>
}
```

- 如果想輸入任意長度的參數可以透過解包`(args...)`來處理，有點像python的`*args`，不過需要注意類別，而且必須作為最後一個參數
```go
func sumArgs(args... int) int { // arg: []int
    var sum_arg int = 0
    for _, arg := range args{
        sum_arg += arg
    }
    return sum_arg
}
```

- 如果想要回傳多個值需要使用括號包起來
順便一提，在Go語言中你甚至可以在宣告函式時就對回傳參數命名，這樣只需要執行後將結果賦值給目標變數然後return就行
```go
// 回傳多參數
func divMod(n, d int) (int, int){
    return n/d, n%d
}

// 命名回傳參數
func namedReturnArgs(args ...int) (returnValue int) {
	returnValue = len(args)
	return 
}
```

- 由於在Go語言中函式也是一個值，因此你也可以把函式作為參數丟入或者是直接回傳函式，其中後者常在閉包(Closure)中使用
```go
 // function as argument
type Comparator func (int, int) bool 

func smaller(a, b int) bool {
    return a < b
}
func customCompare(a,b int, comparator Comparator){
    if comparator(a, b){
        fmt.Println("a is smaller than b")
    }
}

// return function (Closure)
type intFunc = func() int

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() intFunc {
    first, second := 0, 1
    return func() int {
        ret := first
        first, second = second, first+second
        return ret
    }
}
```
### 匿名函式（lambda）
匿名函數是一種宣告時不需要命名的函數。類似於python的`lambda`或是C/C++的`[](){}`或是JS的Arrow function
就像普通函數一樣，匿名函數可以接受輸入並返回輸出。
```go
type returnType func(int, int) int

func returnGcdFunction() (gcd returnType){
    gcd = func(a, b int) int{
        if b==0{
            return a
        }
        return gcd(b, a%b)
    }
}
```
### 閉包(Closure)
閉包（Closure）是函式以及該函式被宣告時所在的作用域環境（lexical environment）的組合，它將變數的參考存放在作用域中而不是變數的值。常見於JavaScript，在 Golang 中，閉包是匿名函數的一種特殊情形

在通常情況下，函式內的區域變數只能活在它所在的函式中。也就是當函式結束後，這個變數就無法繼續在使用直到下次呼叫
```go
func tempVal(){
    x := 100
    println(x)
}
func main(){
    tempVal()
    println(x) // you can't use `x` in here
}
```
因此為了延長區域變數的生命週期，因此我們會使用閉包
```go
type Counter func()
type Setter func(int)
type Getter func() int

func closureCounter() (counter Counter, adder, setter Setter, getter Getter){
    count := 0
    
    // assign function as value to variable
    counter = func(){
        count += 1
    }
    adder = func(val int){
        count += val
    }
    setter = func(num int){
        count = num
    }
    getter = func() int{
        return count
    }
    
    return counter, adder, setter, getter
}

func main(){
    cCounter, cAdder, cSetter, cGetter := closureCounter()
    // counter process
}
```
以上面的範例來說，只要在`main`執行的期間，`closureCounter`中的count會一直存在，就算離開`closureCounter`也不會影響

--------------------
那麼今天的文章就到這告一段落，如果我的文章有任何地方有錯誤請在留言區反應
明天將會介紹Go語言的錯誤處理
![time](https://i.imgur.com/Hk7po4w.gif)
# REF
- https://vocus.cc/article/64edb58cfd897800015bcda1
- https://go.dev/ref/spec#Function_types
- https://go.dev/tour/moretypes
- https://hackmd.io/@upk1997/go-function
- https://developer.mozilla.org/zh-TW/docs/Web/JavaScript/Closures