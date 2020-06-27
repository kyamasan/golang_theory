### IF

2 通りの書き方が可能。result2 は IF 文の外では使えないので注意。

```go
func main() {
	result := 10
	if result == 10 {
		fmt.Println("10")
	}
	if result2 := 20; result2 == 20 {
		fmt.Println("20")
	}
}
```

### for

複数の書き方が可能

```go
func main() {
  //通常のfor
	for i := 0; i < 10; i++ {

  }

  //条件付きfor
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
  }

  //無限ループ
	for {
	}
}
```

### range

スライスや map の要素数分だけループを行う際に、len()を使わずに range でループ処理を書くことができる。

```go
func main() {
	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
  }

	//iを使用しない場合(vはlの要素なので、l[i]のような書き方はしない)
	for _, v := range l {
		fmt.Println(v)
  }

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//kを使用しない場合
	for _, v := range m {
		fmt.Println(v)
	}

	//vを使用しない場合
	for k := range m {
		fmt.Println(k)
	}
}
```

### switch

if 文のように変数を switch 文の外で使わないなら一文で書ける。
switch に条件式で使う変数を書かなくてもよい。

```go
func main() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("MAC!")
	case "windows":
		fmt.Println("WINDOWS!")
	default:
		fmt.Println("OTHER!")
	}
}

func main() {
	switch os := "mac"; os {
	case "mac":
		fmt.Println("MAC!")
	case "windows":
		fmt.Println("WINDOWS!")
	default:
		fmt.Println("OTHER!")
	}
}

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 11:
		fmt.Println("Morning")
	default:
		fmt.Println("Afternoon")
	}
}
```

### defer

遅延実行という意味。main 関数の中で使用すると、「main 関数の処理が終わった後」に実行される。
stacking defer といって、先に書いた defer は以降の defer よりも後に実行される。

```go
func main() {
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
  fmt.Println("success")
/*
	file, _ := os.Open("./main.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
  fmt.Println(string(data))
*/
}
```

> run
> success
> 3
> 2
> 1

### log

fmt と同じく、Println()、Printf()、Fatalln()、Fatalf()などが使用される。

### エラーハンドリング

```go

同名のerrという変数に対して、file, err :=、count, err :=で2回初期化しているが、:=は変数一つでも使えることに注意。2回目では上書きしている。もし戻り値が一つしかなければ、err=...で上書きすればよい。

func main() {
	file, err := os.Open("./main.go")
	defer file.Close()
	if err != nil {
		log.Fatal("error")
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal("error2")
	}
	fmt.Println(count, string(data))
}
```

### panic(非推奨)、recover

panic()は、自分で例外を発生させる関数。golang ではきちんとエラーハンドリングを行うことが推奨されており、panic は非推奨。panic が発生しても defer で recover を実行すれば強制終了を回避できる。

defer には実行するものを渡す必要があるので、func{}()を付ける。

```go
func main() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	panic("panic")
}
```

### 演習

Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。

l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

Q2. 以下の果物の価格の合計を出力するコードを書いてください。

m := map[string]int{
"apple": 200,
"banana": 300,
"grapes": 150,
"orange": 80,
"papaya": 500,
"kiwi": 90,
}

```go
func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	m := l[0]
	for _, l := range l {
		if m > l {
			m = l
		}
	}
	fmt.Println(m)
}

func main() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	var max int
	for _, v := range m {
		max += v
	}
	fmt.Println(max)
}
```
