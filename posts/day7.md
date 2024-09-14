# Day7. 迴圈操作
今天來講Go語言如何處理條件與跳轉

## 基礎指令(for loop)
基本上跟C/C++差不多，除了不能在最外面加括號以外
```go
// Like a C for
for init; condition; post { exec }

for i:=1; i<10; i++{
    println(i);
}
```
如果是想要做到類似其他語言直接對array-like物件做for loop(`for obj in smth`)的話可以在物件前面加上`range`關鍵字
```go
// 類似於python的enumerate(itersObj)
for index, value:= range itersObj{
    // do something
}


// 只需要index的情況(類似於python的 for index in range(len(itersObj)))
for index := range itersObj{
    // do something
}
```
## 條件迴圈(while)
![Image](https://i.imgur.com/IopKpri.png)
在Go語言只有for關鍵字，不過使用for也能做到while的事
```go
// while true in Go
for {
    // do something
}

// while condition
func play_haruhikage(condition bool) bool {
	return condition == true
}

for play_haruhikage() {
    xd()
}
```
![Image](https://i.imgur.com/xToE7Ng.gif)

----------------------
那麼今天的文章就到這告一段落，如果我的文章有任何地方有錯誤請在留言區反應
今天的文章比較簡短，不過明天將會介紹Go語言的陣列，Map，slice等內建資料結構
![time](https://i.imgur.com/Hk7po4w.gif)
# REF
- https://go.dev/doc/effective_go#for