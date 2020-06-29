### スニペットの追加方法

vscode から、

> file → preferences → user snippets

を起動。go.json を選択。

```json
{
  // Place your snippets for go here. Each snippet is defined under a snippet name and has a prefix, body and
  // description. The prefix is what is used to trigger the snippet and the body will be expanded and inserted. Possible variables are:
  // $1, $2 for tab stops, $0 for the final cursor position, and ${1:label}, ${2:another} for placeholders. Placeholders with the
  // same ids are connected.
  // Example:
  // "Print to console": {
  // 	"prefix": "log",
  // 	"body": [
  // 		"console.log('$1');",
  // 		"$2"
  // 	],
  // 	"description": "Log output to console"
  // }
  "Println": {
    "prefix": "println",
    "body": ["fmt.Println()"],
    "description": "fmt.Println()"
  },
  "Printf": {
    "prefix": "printf",
    "body": ["fmt.Printf(\"%T %v\", foo, foo)"],
    "description": "fmt.Printf()"
  }
}
```

### hello world

```go
package main

import "fmt"

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Hello")
}
```

メソッドが大文字であれば Public なので呼び出し可能。
init()は main()よりも先に呼ばれる特別な関数。

### import

Golang の標準ドキュメント
https://golang.org/pkg/

https://golang.org/pkg/os/user/

```go
import (
	"fmt"
	"os/user"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(user.Current())
}
```

### 変数宣言

var を変数全てに書かなくても、()で括る書き方もある。
var を使わない x := ~ という書き方もある。(short variable declaretion)

var を使った変数宣言は関数の外でも実行可能。:=は関数内でのみ実行可能。
var を使えば変数の型を明示的に宣言できるというメリットもある。

fmt.Printf("%T", 変数) で変数の型をチェックできる。
fmt.Printf("%v", 変数) で変数の値をチェックできる。
fmt.Printf("%d", 変数) で 10 進数の値をチェックできる。
fmt.Printf("%s", 変数) で文字列の値をチェックできる。
※fmt.Printf("%T\n", 変数)と付ければ改行される。fmt.Println()は自動で改行される。

```go
func main() {
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)
	fmt.Printf("%T", f64)
}
```

//型変換

文字列からの数値への変換には strconv.Atoi を用いる。
バイト配列から文字列への変換には string(s[0])を用いる。
バイト配列は文字列と関わりが深いので、変換方法が異なる点に注意。

```go
var s string = "14"

i, _ := strconv.Atoi(s)
```

//配列、スライス
配列の場合にはサイズの変更が負荷。スライスは変更可能。

```go
var a [2]int

a[0] = 100
a[1] = 200
fmt.Println(a) //[100 200]

var b [2]int = [2]int{100, 200}

//スライスの場合にはサイズを書かない。
n := []int{1, 2, 3, 4, 5, 6}
fmt.Println(n)
fmt.Println(n[2])
fmt.Println(n[2:4]) //[3, 4]
fmt.Println(n[:2])
fmt.Println(n[2:])
fmt.Println(n[:])

//サイズを変更する際はappendを用いる。
n = append(n, 100)
fmt.Println(n)

```

//make,cap

スライスの　長さは len、キャパシティ(メモリに確保した値)は cap で取得可能。

```go
n := make([]int, 3, 5)
fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
```

```go
// c := make([]int, 5)
c := make([]int, 0, 5)

for i := 0; i < 5; i++ {
  c = append(c, i)
  fmt.Println(c)
}
```

len が 0、キャパシティが 5 になる。

> [0][0 1][0 1 2][0 1 2 3][0 1 2 3 4]

```go
c := make([]int, 5)
// c := make([]int, 0, 5)

for i := 0; i < 5; i++ {
  c = append(c, i)
  fmt.Println(c)
}
```

len とキャパシティが 5 になる。

> [0 0 0 0 0 0][0 0 0 0 0 0 1][0 0 0 0 0 0 1 2][0 0 0 0 0 0 1 2 3][0 0 0 0 0 0 1 2 3 4]

### map

```go
func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)
}
```

map に対して、指定したキーの value があるかどうかを調べるには、m["key"]で value と値があるかどうかの真偽値の 2 つの値で確かめられる。

```go
v, ok := m["apple"]
fmt.Println(v, ok)
```

### メソッド

関数の戻り値に明示的に result という変数を返すように名前を宣言することができる。
この書き方をすると、メソッドの中で result := ...という書き方ができない。

```go
func cal(price, item int) (result int) {
	result = price * item
	return result
}

func cal(price, item int) (int) {
	result := price * item
	return result
}

```

また、result という値を返すことが分かっているので、return result ではなく、return と省略して書くことができる。

```go
func cal(price, item int) (result int) {
	result = price * item
	return
}
```

また、関数の中で別の関数を定義することも可能。
呼び出し型が 2 通り存在する。

```go
func main() {
	f := func(x int) {
		fmt.Println("Inner Func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("Inner Func", x)
	}(1)
}
```

### クロージャ

無名関数と

関数の中で、x と別の関数を定義する書き方。

```go
func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment()) //1
	fmt.Println(increment()) //2
	fmt.Println(increment()) //3
}
```

関数の外で、x と別の関数を定義する書き方。

```go
func incrementGenerator() (func() int) {
	x := 0
	return func() int {
		x++
		return x
	}
}
func main() {
	//counterにはincrementGenerator()の戻り値であるfunc()が格納されている。counter()の呼び出し毎にxは0に初期化されない点に注意
	counter := incrementGenerator()
	fmt.Println(counter()) //1
	fmt.Println(counter()) //2
}
```

クロージャの仕様例(円の面積を求める関数)
circleArea(3.14)と書いた時点で、pi=3.14 の関数が戻り値にセットされ、戻り値の関数の中でもその 3.14 という変数を参照できる。

```go
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	c1 := circleArea(3.14)
	fmt.Println(c1(2)) //12.56

	c2 := circleArea(3)
	fmt.Println(c2(2)) //12
}
```

### 可変長引数

関数の引数の数が一定でない時は、(params ...引数の型)と書く。
また、スライスを可変長引数として関数に渡す時は、foo(s...)と書く。

```go
func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	foo(s...)
}
```

### 演習

Q1. 以下の 1.11 を int 型に変換して出力してください。

f := 1.11

Q2. コードを書かずに以下の出力結果を答えてください。

s := []int{1, 2, 5, 6, 2, 3, 1}
fmt.Println(s[2:4])
Q3. 以下のコードを実行した時に

fmt.Printf("%T %v", m, m)

以下のような出力結果となる m を作成してください。

map[string]int map[Mike:20 Nancy:24 Messi:30]

※Go1.12 より、map をキー順で表示するようになったので、
https://golang.org/doc/go1.12#fmt

> map[string]int map[Messi:30 Mike:20 Nancy:24]

と表示される。

```go
func main() {
	//Q1
	f := 1.11
	i := int(f)
	fmt.Printf("%T, %v\n", i, i)

	//Q2
	[5,6]

	//Q3
	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
```
