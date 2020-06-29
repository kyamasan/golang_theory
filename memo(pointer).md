### ポインタ

&n には n の値を格納しているメモリのアドレスの場所が格納されている。&n の型はポインタ型(\*int)

```go
func main() {
	var n int = 100
	fmt.Println(n) //100
	fmt.Println(&n) //100の格納されたメモリのアドレス

	var p *int = &n //*intはintのポインタ型
	fmt.Println(p) //
	fmt.Println(*p)
}
```

foo()では n の値を直接書き換えていない為、n の値は 100 から変わらない。
bar()では n の値を指すアドレスを渡し、そこから n の実体を書き換えている為、n の値は 1 に変わる。
(後で出てくる struct でも同様)

```go
func foo(x int) {
	x = 1
}

func bar(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	foo(n)
	fmt.Println(n)
	bar(&n)
	fmt.Println(n)
}
```

### new, make

n に何も入れない状態で、メモリに領域を確保する為には new を用いる。

```go
func main() {
	var p *int = new(int)
	fmt.Println(p) //0xc000014090
	var p2 *int
	fmt.Println(p2) //<nil>
}
```

ポインタを返り値とする場合には new(ポインタ、struct)、
それ以外の場合には make(スライス、マップ、チャネル) を使う。

```go
//スライス
s := make([]int, 0)
fmt.Printf("%T\n", s)

//マップ
m := make(map[string]int)
fmt.Printf("%T\n", m)

//ポインタ
var p *int = new(int)
fmt.Printf("%T\n", p)
```

### struct

struct の要素(X,Y,S)は基本的には全て大文字で書く。小文字だと private の意味になる。
要素は{X:.., Y:..}のように指定することもできるが、指定がない場合は要素の順にセットされる。

```go
type Vertex struct {
	X、Y int
	S string
}

func main() {
	v1 := Vertex{X: 1, Y: 2}
	fmt.Println(v1)
	v2 := Vertex{1, 2, "test"}
	fmt.Println(v2)
}
```

v3、v4 のような呼び出し方が存在する。

```go
func main() {
	v3 := Vertex{}
	fmt.Println(v3)

	var v4 Vertex
	fmt.Println(v4)
}
```

ポインタ型を返す方法も存在する。
struct のポインタを返したい場合には、new で宣言するよりも&を宣言する方がポインタ型が返されるという事が分かりやすいので推奨されている。
※スライスや map の場合は、逆に make を使う方法が推奨されている。

```go
func main() {
	v5 := new(Vertex)
	fmt.Println(v5)

	v6 := &Vertex{}
  fmt.Println(v6)

	m1 := make([]int, 0)
  fmt.Printf("%T %v\n", m1, m1)

	m2 := []int{}
	fmt.Printf("%T %v\n", m2, m2)
}
```

また、struct のポインタ型は、\*を付けて実体を参照しなくても自動で実体を見に行くことができる。
int の時とは異なるので注意。

```go
v1 := &Vertex{X: 1, Y: 2}
v1.X = 1000 //(*v1).X=1000とも書ける。
fmt.Println(v1)
```
