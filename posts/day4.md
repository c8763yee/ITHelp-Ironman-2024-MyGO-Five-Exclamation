# Day4. 介紹語法、變數宣告與型態
今天來介紹Go語言的語法、變數宣告與型態

首先我們要了解一個概念，Go 程式是透過 package 來組織的。`package <pkgName>`告訴我們當前檔案屬於哪個套件(main 則告訴我們它是一個可獨立執行的套件，它在編譯後會產生可執行檔案)。除了 main 套件之外，其它的套件最後都會產生*.a檔案（也就是套件檔案）並放置在`$GOPATH/pkg/$GOOS_$GOARCH`中

所以我們需要先建立一個資料夾，然後進去那個資料夾之後運行`go mod init <dir_name>`
```bash
mkdir go_practice && cd go_practice
go mod init go_practice
```


## Hello World
我們從國際範例的Hello World開始撰寫，首先在剛才建立的資料夾中新增main.go
```go
package main

import "fmt"

func main()
{
    fmt.Println("Hello, world or 你好，世界 or καλημ ́ρα κóσμ");
}
```
首先我們在第一行指定package名稱，如果是要將這個檔案獨立執行的話通常會設定成main。再導入fmt套件，這個套件提供了輸入與輸出的功能，最後定義main function並且將文字輸出。

撰寫完Code後我們可以在Terminal執行`go run <file-path>`來運行這個檔案
![Image](https://i.imgur.com/g6tBzxo.png)

# 蛤
![Image](https://i.imgur.com/5UZiDv8.gif)

你可能會好奇，為什麼簡單的Hello World也能報錯（尤其是習慣寫C/C++或JAVA的人)。其實是因為Go語言對於預設會將分號隱性插入每行最後面，因此這個大括號的結構就會被破壞。因此正確的語法應該是

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
> 其他相關語法錯誤可以到[這裡](https://www.kancloud.cn/uvohp5na133/golang/934228)參考
![Image](https://i.imgur.com/DgtQEPP.png)
改完後就能正常執行了



## 變數宣告
- 變數名稱不可與關鍵字衝突
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

Go語言可以使用多種方式來宣告與定義變數。其中`var`因為能在函式外部使用，因此常用於定義全域變數
```go
// 你可以這樣(定義指定type變數)
// var <variableName> [<type>] [= <value>]

var variableName string = "string"
var autoType = 213879 // 自動指定類別

// 或這樣（定義多個變數）
// var <variableNames>... [<type>] [= <value>]
var var1, var2, var3 int = a, b, c
// 針對不同型別
var c, python, java = true, false, "no!"


// 如果你在function裡面的話甚至可以這樣宣告(簡短宣告，忽略型態)
// <variableName1>, <variableName2>, ... := <val1>, <val2>, ...
func f(){
    int32_max, int32_min := 1 << 31 - 1, -(1 << 31)
}

// 或是用括號包起來
var (
    frequency = 16 / 10
)
```
然後如果變數不需要使用的話需要去除該變數或是指定到下划線變數`_`，否則無法編譯
```go
// Wrong
x, y := 1, 2
fmt.Println(x)

// correct
x := 1
//     or
x, _ := 1, 2
```
如果要定義常數則使用`const` 關鍵字
對於遞增常數，可以配合iota關鍵字簡化定義，Golang 會自動把它從const中的第一行以0開始每行遞增（忽略註解或空行）
```go
// const <constantName> [<value>] = value
const PI float64 = 3.14159265358979323846264

const (
    // 可自訂iota表達式，如果這行沒設定iota表達式則沿用上一行
    NoneVal = 1
    Cat = iota // 0(iota=0)
    Dog = 1 << (iota * 10) // 1024(iota=1) -> 1 << 10
        // skipped(iota=1)
    /*
        These Line will be skipped

    */ //skipped(iota=1)


    Bird // 1048576(iota=2) -> 1 << 20
    Pig = iota // 3(iota=3)
)
```

你還能透過別名(alias)為型態命名，如`byte`其實就是`uint8`的別名
```go
type myint = int
type myint int // 無法與int互通
```


## 變數型態
需要注意的一點是，數值型別的變數之間不允許互相賦值或操作，不然會在編譯時引起編譯器報錯。
如下的程式碼會產生錯誤：invalid operation: a + b (mismatched types int8 and int32)
```go
var a int8
var b int32
c:=a + b
```
另外，儘管 int 的長度是 32 bit, 但 int 與 int32 不可混用。


### 布林型態
- `bool`: 就true跟false，不過不能拿來跟數值比較或運算(也就是不能當數字的0或1)，只能當成是與否

### 數值型態
| Type                      | Size         | Default Value | note                                                                                    |
| ------------------------- | ------------ | ------------- | --------------------------------------------------------------------------------------- |
| int                       | 4/8 bytes    | 0             | 大小根據系統位元數決定(在32bit系統為4bytes, 64bit系統為8bytes)                          |
| int8, int16, int32, int64 | 1~8 bytes    | 0             | 直接定義好位數的int，其中後面的數字代表bit數，不過除非有特定需求不然建議直接使用int     |
| uint                      | 4/8 bytes    | 0             | 類似int，但沒有負號                                                                     |
| int8, int16, int32, int64 | 1~8 bytes    | 0             | 指直接定義好位數的uint，其中後面的數字代表bit數，不過除非有特定需求不然建議直接使用uint |
| float32                   | 4 bytes      | 0.0           | IEEE-754 32位元浮點數                                                                   |
| float64                   | 8 bytes      | 0.0           | IEEE-754 64位元浮點數                                                                   |
| complex64                 | 8 bytes      | (0+0i)        | 實部和虛部為 float32 的複數                                                             |
| complex128                | 16 bytes     | (0+0i)        | 實部和虛部為 float64 的複數                                                             |
| uintptr                   | 4 或 8 bytes | 0             | 用於存儲指標值的未解釋位元                                                              |

### 字串型態
#### rune
int32的別名，表示該字元為Unicode，單引號包起來的字元(char)預設為這個類別

#### byte
uint8的別名，表示該字元為Raw data

#### string
- 本身為不可變，需轉換為[]rune或[]byte才能修改
- 預設值為空字串
- 跟C不一樣，字串不以`\0`結尾

```go
func main(){
// byte
    var va1 byte = 123

// rune
	runeval := 'a'

// 單行字串
    s:= "single line string"

// 多行字串(不會處理反斜線)
    haruhikage := `
悴んだ心 ふるえる眼差し世界で
僕は ひとりぼっちだった
散ることしか知らない春は
毎年 冷たくあしらう
暗がりの中 一方通行に ただただ
言葉を書き殴って 期待するだけ
むなしいと分かっていても
救いを求め続けた
（せつなくて いとおしい）
今ならば 分かる気がする
（しあわせで くるおしい）
あの日泣けなかった僕を
光は やさしく連れ立つよ
雲間をぬって きらりきらり
心満たしては 溢れ
いつしか頬を きらりきらり
熱く 熱く濡らしてゆく
君の手は どうしてこんなにも温かいの？
ねぇお願い どうかこのまま 離さないでいて
縁を結んでは ほどきほどかれ
誰しもがそれを喜び悲しみながら
愛を数えてゆく
鼓動を確かめるように
（うれしくて さびしくて）
今だから 分かる気がした
（たいせつで こわくって）
あの日泣けなかった僕を
光は やさしく抱きしめた
照らされた世界 咲き誇る大切な人
あたたかさを知った春は 僕のため 君のための
涙を流すよ
Ah なんて眩しいんだろう
Ah なんて美しいんだろう
雲間をぬって きらりきらり
心満たしては 溢れ
いつしか頬を きらりきらり
熱く 熱く濡らしてゆく
君の手は どうしてこんなにも温かいの？
ねぇお願い どうかこのまま 離さないでいて
ずっと ずっと 離さないでいて
`
    /*
###################################################*****************************
#######################################******#########**#***********************
###############################%*=-::::--===-======+##**#***********************
############################*=========---:------------=*###*###*****************
#######################%*=========+=---::-=----------=---+####******************                なんで『春日影』やったの
###################*+*#+========+=---------==+=----------==+#####***************
###############*+*###++++=====+=--------==-=*=---------=-=-=-*##*###************
###########*+######*++=+=++=*=----=-----=-++-::--------==+====*#####************
########**######*+=+***+*=++=---==+----=-++---=------===#=-=-==+######**********
######+######+=+**+*+=++=+=-----+====-==*+=====----=+-+%+=-=====*#**####********
####+#####*=+##==+*+=*+++-==--=++=+===-++-==+++==+*==*-#+--+-===-##########*****
##=#####+=%##*=+**==*=+=-----=*==*--=-#*+=--==#*#==#-:-*+-=+-+=---###########***
#*####++####*=+**=*+=*=---=-=*+*+-===##+===--+++*+:...-*+-+=+==++-*############*
####*+#####*=***+*==+=----==#**==--*-:..-:++**=.......:*+=*=#==*+=+#############
##%*######+=**#+==++-----==**=-==++:.-::--:-...........*+=*+*+=*+=+#############
##*######++*#====++=----=+#+-===+--++@%####%%+:........:*=*+#+=*+==#############
%%#######+===++=*==---==*+=-+**=:.%@+:......-#**+==-....=+==#+=*+==#############
##***==+++**+==*=----==+-=+*++:...:#.....=##-.-..:...=-=:+=++*=++==#############
**##%##*#====*%-----=+===#++=:......-....#++*:.......::%@@*+-*+++=+#############
%#++*++==+*##*==-==-+==*#=*=:..........-===::.........::.:*%+**+*=*#############
*+===+*##*+*+--=*==#-=++=+#:...............:-:.........##-.+#*+=+=##############
+*#+**#++**==++==+#-+=+=+--...........................:=+=.:%#+*=*##############
*****=++*+=++=-=*#++=+==---.........................::=+=:.-%#+++###############
#+=+*+**+*=--=*#*#+**=*----:................:--:.....=..:--+%+*=*###############
+****#*+---=*#*#****+**#**#*..............=*####-=:.......:#++++################
+**#+==--=*##*****#**#***#*#=............#########-=:.....+*=*=#################
++-=--=+###*******##***####*#*..........*###########=....:#+=**#################
=====####*******###***#*##**#%=.........==----=+####=....*++#=##################
-+*##********##****#######***%=-:........:--+=---=**....-**+*+##################
#********###*****##*####***##*===-:.........:-:...-....+*+#**+########+*########
******###%%#*++*%#*####***%*%--=====+:...............***+***#+#####*=+#+=#######
*###*=:..........:+##***####=----=====**-.......:-+#****+***%*####+*####=#######
#+-..............-+*#*=*%@#-------=====-#######****#***#+*#*#%**+*######*#######
#++=+.....:---:........-=:==---------==*#%#**#*****#***#*+######################
:.:--=........:+=:.......:+--=+*+=----:%##*#*#*****#**+##*=#*%%%#***####%#######
--::-=+:...:=-.............:=.:-----=*==*##***#+-+#*#+**##*+++++++*#*+=*########
.....--==:-+=--==:...........-:=-::-=-:+---::.::*-:.-#*****##*****=-+##*########
.......:--=*+----==.........-.....=..:=====-....=...:=.:*********##+=###########
.............==---=+-......-......*..........-==......:-:.+###%#*****###########
..............:+*+=-=-....--.....-*.............+:.......=..#%###***############
.............-====..--+:.........-*.........:.....=:......:-.-##%*##############
...........:====*:....--*-......:=+.........--.....:+.......+.-%######%#########
.........:-====+-.......:-==-:..:==:.........--......==......-.*################
.......:-=====+-............:--=++*+====+-....=-......-=-.:=+==+################
    */
}
```
那麼今天的文章就到這告一段落，如果我的文章有任何地方有錯誤請在留言區反應
明天將會說明怎麼處理輸入與輸出(可以想成Python的`print`跟`input`或C的`scanf`與`printf`)。
![time](https://i.imgur.com/Hk7po4w.gif)

-----------------------------
# REF
- https://go.dev/doc/effective_go
- https://www.kancloud.cn/uvohp5na133/golang/934228
- https://willh.gitbook.io/build-web-application-with-golang-zhtw/02.0/02.2
- https://hackmd.io/@upk1997/go-datatype
- https://blog.wolfogre.com/posts/golang-iota/