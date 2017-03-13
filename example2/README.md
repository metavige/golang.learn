
# 寫程式一定要會用 log

## 簡單的 op log

找到一個簡單的 logging 套件 [https://github.com/op/go-logging]()  
網站的 README 相當清楚，如果沒有什麼特別要設定的，只是輸出 console.log 的話，已經很好用了  

## rolling file

之前有一些專案，有用到 rolling file，剛剛找了一下，有個 [https://github.com/NYTimes/logrotate]() 看起來不錯  
因為有結合 SIGHUP 的信號，所以當有 `kill -HUP` 動作的時候  
就會有 rolling file 的動作，表示這可以跟 unix 的 logrotate 功能結合～  
不錯～  

如果搭配上面的 go-logging 套件，只需要設定一個 backend 就可以  

```go
logfile, err := logrotate.NewFile("sample.log")
if err != nil {
    log.Fatal(err)
}

backend1 := logging.SetBackend(logging.NewLogBackend(logfile, "", 0))
logging.SetBackend(backend1)
```

> 本來以為介面上 `logging.SetBackend` 說的是要用 `io.Writer`, 所以我本來還用了 `bufio.NewWriter` 去轉換 file
> 後來發現是多此一舉，因為 File 本來就有實作 `io.Writer`

## 比較強大的 log 

我有找到一個號稱做 "Abstract Logging" 的套件 - [https://github.com/birkirb/loggers]()  
光看他後面一個 [logrus](https://github.com/Sirupsen/logrus) 的實作，就覺得好強大！   
之後真的在比較大的專案裡面或許可以用得到  