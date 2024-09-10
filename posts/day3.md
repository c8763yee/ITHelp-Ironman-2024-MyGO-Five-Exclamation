---
title: Day3. 安裝Go與設定開發環境
tags: [鐵人賽, Golang]

---

# Day3. 安裝Go與設定開發環境
如果只是想嘗試一下Go語言的語法的話，官方有提供[Playground(https://go.dev/play/)](https://go.dev/play/)，可以在這裡嘗試Go語言的語法。官方也有提供一些範例可以進行參考
## 安裝Go
首先，進到[Go官網的下載界面(https://go.dev/doc/install)](https://go.dev/doc/install)，點擊Download按鈕![Image](https://i.imgur.com/MBNbviE.png)再根據自己系統下載檔案![Image](https://i.imgur.com/cKMbdzK.png)

### 各系統的安裝方式
- Linux:
    通常都安裝在`/usr/local/go`，不過也可以自己指定
    > 這個路徑之後會被存在 叫`GOROOT`的環境變數中。
    1. 如果之前已經在\$GOROOT安裝過的話需要先刪除然後重新解壓至$GOROOT。因為直接覆蓋會破壞環境，
    > Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.
    ```bash
    # you may need to add `sudo` in front of `rm`
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
    ```
    > 順便說一下，如果你使用ArchLinux的話可以直接打`sudo pacman -S go`來安裝go

    2. 如果之前完全沒有安裝過的話需要將安裝路徑加入環境變數，可以透過在\$HOME/.profile或/etc/profile或\$HOME/.bashrc加入以下內容
    ```bash
    # $HOME/.profile
    export PATH=${PATH:+${PATH}:}/usr/local/go/bin
    ```
    然後再對於該檔案執行source指令以應用變更
    ```
    source $HOME/.profile
    ```
- Windows
    1. 打開剛才安裝的MSI檔然後根據提示安裝
    > 可以在這一步驟指定安裝路徑
    > ![Image](https://i.imgur.com/bzOqJfS.png)
    > 出現這個畫面就代表安裝完成了
    >![Image](https://i.imgur.com/UGUHAXc.png)

- Mac
    - 我沒用過，不過可以參考下載界面中的安裝教學

安裝完之後開啟Terminal然後執行`go`，如果有出現幫助文字的話就代表安裝成功了
## 設定開發環境
再來是選擇文字編輯器，基本上用自己習慣的就行。因為我自己比較習慣使用Visual Studio Code，因此接下來都是以Visual Studio Code為主。
下載完Vscode之後打開Extension界面，然後搜尋`Go`，之後選擇第一個套件安裝
> 前兩個都是官方的套件，差別是Nightly是測試版，可能會出現Bug。所以還是建議安裝第一個
![Image](https://i.imgur.com/yPo5som.png)

### 設定運行Go與偵錯
1. 首先在vscode同時按下<kbd>ctrl</kbd>+<kbd>shift</kbd>+<kbd>p</kbd>，會出現類似下圖的指令清單。
![Image](https://i.imgur.com/4yzSi11.png)
2. 輸入`Preferences: Open Default Settings(JSON)`然後點進去，會跳轉到一個json檔案，這個json檔案就是你的設定檔。
> `Preferences: Open Default Settings(UI)`可以透過圖形化介面進行一些設定，但偵錯相關的設定還是需要在json中編輯
1. 將下面的json複製到設定檔裡

```json
"launch": {
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${file}",
        }
    ]
}
```

設定完後可以按下F5來運行你的程式，如果有在下面的輸出看到`exited with status 0`的話代表你的程式能正常運行。反之則會出現紅字的錯誤訊息
- 成功執行
![Image](https://i.imgur.com/kIxIF3i.png)
- 程式出錯
![Image](https://i.imgur.com/GbNLVPl.png)


----------------------------------
那麼今天的文章就到這告一段落，如果我的文章有任何地方有錯誤請在留言區反應
明天將會實際撰寫Go語言的Code，並且介紹Go語言的語法與變數宣告。
![time](https://i.imgur.com/Hk7po4w.gif)

# Reference
- https://go.dev
- https://ithelp.ithome.com.tw/articles/10238252
- https://code.visualstudio.com/docs/languages/go