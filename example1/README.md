
## defer

- 了解 `defer` 後面一定要接一個 func 才可以
- `defer` 就相當於 finally, 所以最後一定會執行，不管錯誤與否
- 不過這個 `defer` 是寫在前面，目的是有中斷流程的意義在裡面

## 怎樣把 Body 抓出來 (print)

找了[網路範例](https://dlintw.github.io/gobyexample/public/http-client.html])才知道，要用 `ioutil.ReadAll()`

> 後來發現[這個網站](https://dlintw.github.io/gobyexample/public/index.html)不錯～ 很多範例程式

## interface
 
以下是 response Body 所實作的介面，說明說到，這個 ReadCloser 是有 *groups* Read + Close 方法的介面  
```go
// ReadCloser is the interface that groups the basic Read and Close methods.
type ReadCloser interface {
	Reader
	Closer
}
```

而 `func ReadAll(r io.Reader) ([]byte, error)` 方法是這樣

所以～像是 ReadCloser 繼承了 Reader 這個介面，所以可以傳入 ReadAll 方法  

## byte[] 轉型成 string

用下面的方式就可以～  


```go
string(body)
```