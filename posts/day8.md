# Day8. 內建資料結構
今天醒來發現眾志成城獎沒了，在最後一天突然有一團35個人的team參賽
![Image](https://i.imgur.com/ZnUEfF7.gif)
事已至此還是繼續寫文吧，今天來講Go語言的內建資料結構，比較常見的有array, slice, map
## 陣列(array)
- 宣告方式
```go
// var <valName> [<length>]<type>
// length: array長度
// type: array類型

var byteArray [32]byte

// 宣告多維Array
var multipleDimensionArray [312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312][312]int

// 也能在宣告同時賦值
array := [...]float64{7.0, 8.5, 9.1}
```
- 為一個在固定長度(記憶體大小)中儲存相同類別物件的序列，若沒有在宣告時賦值則Array裡面的值為該類別的預設值
- Array在Go語言中有幾個特點
    1. Array的長度也算在Array Type的一部份且不能為負數，換句話說`[123]int`跟`[456]int`為不同的類別
    2. Array是一個值而不是指標，因此當你將一個Array賦值給另一個Array或是傳到function時，實際上是將所有元素逐一複製，而不是複製指向同一個記憶體位置的指標
    ```go
    // 使用[...]讓Go編譯器自己算長度
    originalArray := [...]int{1, 2, 3, 4, 5}
    after := originalArray

    // some of modify to original array
    after[3] = 7

    fmt.Printf("original Array: %#v\n", originalArray) // 1, 2, 3, 4, 5
    fmt.Printf("Array after modified: %#v\n", after) // 1, 2, 3, 7, 5
    ```
    可以注意到就算修改了after中的值，originalArray還是沒有改變。
    如果想要直接對原Array修改、或是使用參考(reference)的話可以使用指標(pointer)來提昇效能
    ```go
    var originalArray [4]int
    var after *[4]int
	after = &originalArray
    after[3] = 7

    fmt.Printf("original Array: %#v\n", originalArray) // 1, 2, 3, 7, 5
    fmt.Printf("Array after modified: %#v\n", *after) // 1, 2, 3, 7, 5
    ```

## slice
在介紹Array時說到可以使用pointer來提昇效能，不過Go官方比較建議使用slice進行操作。在實際開發中也比較少使用Array

Slice本質上算是一個對於它自己底層的Array的window，所以在function中修改slice參數的值時也會同時影響到原本底層的array
> slice 只存了能夠連結到底層 array 的 pointer，在思考的時候，要把 "slice" 當成一個遮罩或視窗（window），透過它過濾原本 array 的內容後，看到的內容才是 slice 的結果。

順帶一提你其實可以把slice當成一個struct，並且slice初始化後如果沒有賦值的話預設值為`nil`
```go
type slice struct {
    array unsafe.Pointer //指向底層array的pointer
    len   int            //長度(目前slice中元素數量)
    cap   int            //容量(目前slice中做多能存多少)
}

//
lengthOfSlice := len(slice_)
sliceOfSlice := cap(slice_)
```

- 宣告方式
```go
var slices []int
pi := []int{3,1,4,1,5,9,2,6} // 中括號內不放任何物件才是slice

// 從現有array 或slice建立slice
newSlice = oldSlice
newSlice = arr[:] // 以arr作為底層`array`
newSlice = arr[4:7] // 以arr作為底層`array`，並選取其中的index (4~7)進行遮罩

anotherNewSlice = newSlice // 兩個變數都是指向同一個記憶體位置
// 也可以透過make([]T, length, capacity)來建立(如果你已經確定slice大小)
madeSlice := make([]int, 0, 5)
```
補充一下，不管是slice還是array，在獲取元素(`[k]`)時index千萬不可為負數或超出底層array的capacity

- `func append(s []T, vs ...T) []T`
前面講到，slice的長度是可變的，而在go語言中我們將使用進行append操作
注意: 如果slice在append後超出該底層Array的capacity，那麼將會重新分配一個新的Array作為底層
```go
newSlice = append(oldSlice, newElement1, newElement2) // append兩個新元素
newSlice = append(oldSlice, appendSlice...) // append整個slice(等同python的list.extend())
```
- `func copy(dst, src []T) int`
如果只需要對array(slice)的一小部分進行操作的話，建議把那一小部分copy到新的變數中，然後配合append建立新slice與底層array
copy回傳的值為有多少個元素被複製過去
```go
size = high-low
cloneSlice = make([]T, size)
copy(cloneSlice, oldSlice[low: high])
```
> `cloneSlice` 的元素數量如果「多於」被複製進去的元素時，會用 zero value 去補。例如，當 `cloneSlice` 的長度是 4，但只複製 3 個元素進去時，最後位置多出來的元素會補預設值。
> `cloneSlice` 的元素數量如果「少於」被複製進去的元素時，超過的元素不會被複製進去。例如，當 cloneSlice 的長度是 1，但卻複製了 3 個元素進去時，只會有 1 個元素被複製進去。

- 對slice再次擷取(re-slicing)
你可以從一個slice中再次進行擷取以得到新的slice, 而這兩個slice共用同一個底層Array
```go
newSlice = oldSlice[start: end]
```
> 須要留意的是，在 golang 中使用 : 來切割 slice 時，並不會複製原本的 slice 中的資料，而是建立一個新的 slice，但實際上還是指稱到相同位址的底層 array（並沒有複製新的），因此還是會改到原本的元素：
```go
func main() {
	scores := []int{1, 2, 3, 4, 5}
	newSlice := scores[2:4]
	fmt.Println(newSlice)  // 3, 4
	newSlice[0] = 999      // 把原本 scores 中 index 值為 3 的元素改成 999
	fmt.Println(scores)    // 1, 2, 999, 4, 5
}
```
- 擴展slice
對於re-slicing的slice，如果在re-slicing時需要擴展slice 大小時，可以在end index指定比原本slice時還大的index(不可超過capacity)來達到擴展slice的效果
```go
func main() {
    s := []int{1, 2, 3, 4, 5}
    newSlice := s[1:3]               // [2,3], len = 5, cap = 5

  	// 沒有指定 end index 並不會 expand slice
    nonExpandedSlice := newSlice[0:] // [2, 3]

    // 如果想要 expanded slice 需要指定 end index 才行
    expandedSlice := newSlice[0:4] // [2, 3, 4, 5]
}
```
## Map
基本上等同於python的`dict`或是C++的`std::unordered_map`，不過定義完後需要透過`make`來初始化才能使用。
```go
var m map[string]func() int // 對，你可以把function作為value

func a() int {
	return 1
}

func main() {
	m = make(map[string]func() int)
    // 設定指定key的值
	m["a"] = a

    // 或是使用賦值宣告
    var m = map[string]func() int{
        "a": a,
    }

    // 或是直接簡單宣告
    m := map[string]func() int{
		"a": a,
	}
}
```
至於要獲取map中指定的值的話除了一般的`val = map[key]`以外也能透過`val, ok = map[key]`來初步檢查該key是否在map中，至於刪除元素則是透過`delete(map, key)`來刪除元素


--------------------
那麼今天的文章就到這告一段落，如果我的文章有任何地方有錯誤請在留言區反應
明天將會介紹Go語言的函式與log
![time](https://i.imgur.com/Hk7po4w.gif)
# REF
- https://go.dev/doc/effective_go#arrays
- https://hsinyu.gitbooks.io/golang_note/content/slice.html
- https://pjchender.dev/golang/slice-and-array/
- https://coolshell.cn/articles/21128.html
- https://dev.to/gholami1313/re-slicing-in-golang-bp0