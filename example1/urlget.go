package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := getUrlResponse("https://golang.org")
	// finally=> defer
	// 最後一定會執行的動作，不過會先寫在最前面
	// 關閉 body
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("some error happen: ", err)
	} else {
		fmt.Println(resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read body error:", err)
		} else {
			fmt.Println("this is body: ", string(body), 3)
		}
	}
}

func getUrlResponse(path string) (resp *http.Response, err error) {
	resp, err = http.Get(path)
	if err != nil {
		//panic(err)
		return resp, err
	}
	return
}
