
## 自動產生測試案例，連懶惰的理由都沒有

[gotests][]  這個套件，可以很簡單地幫你產生 go 的測試案例，只需要定義好測試案例所需要的資料跟結果就好  
一切都是 [gotests][] 程式碼幫你產生好的  
我自己搭配 vscode，就可以很快速地產生我要的測試案例程式～  
我連懶惰的理由都沒有！  

我的一個簡單方法

```go
func multiple(x int, y int) int {
	return x * y
}
```

產生出來的測試程式，真心不騙，太容易了 ～～～  

```go
func Test_multiple(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// test case
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiple(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

vscode 本身如果有安裝 [go 套件](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go) 的話，可以直接使用它提供的 command  
`Go: Generates unit tests (file)` 就可以替你的程式碼產生測試功能   
其它這個套件提供的功能還真是不少，改天再來研究看看    

## 測試方法注意事情

- 測試方法的檔案是 原檔案名稱_test.go 結尾  
- 一定要 import testing package
- 測試方法 TestXxx() 的參數是 testing.T（測試方法 "Test" 後的第一個字必須不可以是小寫！ ）

老實說，除非有必要，真的用產生的還蠻快的，連格式都不會錯誤  
但是我自己是覺得，如果你自己寫的程式太複雜，你覺的這樣沒辦法測試，你該思考的是，你是不是寫太大包了？  
是不是應該要把方法再細化一點？  
因為測試案例其實不複雜，複雜的是人～  

## stucts 初始化～

我以上唯一搞不太懂的是怎麼寫測試資料，就是下面那一段：  

```go
type args struct {
	x int
	y int
}
tests := []struct {
	name string
	args args
	want int
}{
	// test case
}
```

後來找了一下 golang struct initialize 的部分才找到範例  
基本上這種語法也不用太傷腦筋，背起來就是  
(其實 [gotests][] 這個套件的 README 已經教你怎麼寫了～ )

[gotests]: https://github.com/cweill/gotests