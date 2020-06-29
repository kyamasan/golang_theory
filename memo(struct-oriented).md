### メソッド(値レシーバ)

Go にはクラスが存在しないため、通常のオブジェクト指向型言語のようなアーキテクチャはできない。
その為、Vertex という struct が持つメソッドを実行する、といった事が出来ず、以下のように関数が Vertex を引数として受け取るような構成になってしまう。

```go
type Vertex struct {
	X, Y int
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
}
```

メソッドとして書く為には、Area 関数の前に(v Vertex)を宣言し、Vertex という struct にメソッドを紐づけることができる。

```go
func (v Vertex) Area() int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Area())
}
```

### ポインタレシーバ

ポインタを参照することで、オブジェクトの実体を書き換えることができる。その為、オブジェクトの実体を書き換えたい場合には、ポインタに紐づくレシーバを書く。
また、この時 struct のポインタを(&v).Scale(10)のように書かなくても暗黙的にコンパイラが変換を行ってくれる。(もちろん、書いても良い)

```go
func (v Vertex) Area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) { //値レシーバ(v Vertex) Scaleを書いても実体のvの値は変わらない。
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) //(&v).Scale(10)と同じ意味
	fmt.Println(v.Area())
}
```

### コンストラクタ

golang は python 等の言語とは異なり、クラスが存在しない。
その為、コンストラクタの書き方は以下のようになる。

①struct の要素(X,Y)を小文字にすることで、他のパッケージからは x, y が見えなくなる。
他のパッケージからは New()を呼び出すことで、Vertex の struct を生成することが多い。

```go
type Vertex struct {
  // X, Y int
  x, y int
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
  v := New(3, 4)
  ...
}
```

### embedded

Go には継承という機能はないが、Embedded という機能はある。

```go
type Vertex struct {
	x, y int
}

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())
}
```

### Non-struct メソッド

type で型に名称を付けられるのは、struct だけではない。
レシーバを付けることも可能。

```go
type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}
```

### インターフェース

Human インターフェースの中に入れる struct は Say()を実装している必要がある。

```go

type Human interface {
	Say()
}

type Person struct {
	Name string
}

func (p Person) Say() {
	fmt.Println(p.Name)
}
func main() {
	var mike Human = Person{"Mike"} //Person{Name:"Mike"}
	mike.Say()
}
```

Say()の中で、p の実態を書き換えたい場合は、ポインタ型を渡す。

```go
func (p *Person) Say() {
	p.Name = "Hi Mr." + p.Name
	fmt.Println(p.Name)
}

func main() {
	var mike Human = &Person{"Mike"}
	mike.Say()
}
```

### ダックタイピング

もしもそれがアヒルのように歩き、アヒルのように鳴くのなら、それはアヒルである。
(明示的に実装しなくてもインタフェースが定義するメソッドをすべて実装していれば、そのインタフェースを実装していることになる)
https://qiita.com/tenntenn/items/e04441a40aeb9c31dbaf#%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%95%E3%82%A7%E3%83%BC%E3%82%B9

DriveCar()では、引数に Say()を実装している必要がある為、強制的に Say()を持つ interface を定義する。

```go
type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func DriveCar(h Human) {
	if h.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get Out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	DriveCar(mike)

	var x Human = &Person{"X"}
	DriveCar(x)

  //cannot use dog (type Dog) as type Human in argument to DriveCar: Dog does not implement Human (missing Say method)
	// var dog Dog = &Dog{"Hachi"}
	// DriveCar(dog)
}
```

### タイプアサーション

関数 do()の中で、int や string など複数の型を引数に取りたい場合は、interface{}(何の型でも引数に取れる)を引数にセットし、do()の中で引数の型に応じて処理を変える。

空インタフェース型 (interface{}) の変数（型の定まっていない変数）を、特定の型 T の変数に代入するには、型アサーション (Type Assertion) という仕組みを使用する必要があります。
switch-type 文と組み合わせることが可能。
https://tour.golang.org/methods/16

```go
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	do(10)
	do("Mike")
	do(true)
}
```

### Stringer

struct で String()を実装すれば、fmt.Println()で表示される内容を自由に変更できる。
例えば、以下のように String()を記述しない場合、fmt.Println()では Name と Age の情報が全て出力される。

```go
type Person struct {
	Name string
	Age  int
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike) //{Mike 22}
}
```

String()を実装することで、Age の情報を出力しない、などの制御が可能。

```go
func (p Person) String() string {
	return fmt.Sprintf("My Name is %v", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike) //My Name is Mike
}
```

### カスタムエラー

User が存在しない場合に、カスタムエラーを作成する。
この時、Error()はポインタレシーバにしておくことが望ましい。
値レシーバだと同一のエラー内容を比較すると true が返される為、エラー種類によって処理を変化させることができなくなる。

```go
type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.UserName)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{UserName: "Mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
```
