package main

import (
	"fmt"
)

/* hello world */

// func init() {
// 	fmt.Println("init")
// }

// func main() {
// 	fmt.Println("Hello")
// }

/* import */
// func main() {
// 	fmt.Println("Hello")
// 	fmt.Println(user.Current())
// }

/* 変数宣言 */
// func main() {
// 	var (
// 		i    int     = 1
// 		f64  float64 = 1.2
// 		s    string  = "test"
// 		t, f bool    = true, false
// 	)
// 	fmt.Println(i, f64, s, t, f)
// 	fmt.Printf("%v\n", f64)
// }

/*
func main() {

	var a [2]int

	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n =append(n, 100)
}
*/
/*  */

// func main() {
// 	// n := make([]int, 3, 5)
// 	// fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)

// 	// c := make([]int, 5)
// 	c := make([]int, 0, 5)

// 	for i := 0; i < 5; i++ {
// 		c = append(c, i)
// 		fmt.Println(c)
// 	}

// }

/* map */

// func main() {
// 	m := map[string]int{"apple": 100, "banana": 200}
// 	fmt.Println(m)
// 	fmt.Println(m["apple"])
// 	m["banana"] = 300
// 	fmt.Println(m)
// 	m["new"] = 500
// 	fmt.Println(m)

// 	v, ok := m["apple"]
// 	fmt.Println(v, ok)

// 	v2, ok2 := m["nothing"]
// 	fmt.Println(v2, ok2)
// }

/* メソッド */
// func cal(price, item int) (result int) {
// 	result = price * item
// 	return
// }

// func main() {
// 	r := cal(100, 2)
// 	fmt.Println(r)

// 	f := func(x int) {
// 		fmt.Println("Inner Func", x)
// 	}
// 	f(1)

// 	func(x int) {
// 		fmt.Println("Inner Func", x)
// 	}(1)
// }

/* クロージャ */
// func incrementGenerator() func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }
// func circleArea(pi float64) func(radius float64) float64 {
// 	return func(radius float64) float64 {
// 		return pi * radius * radius
// 	}
// }

// func main() {
// 	counter := incrementGenerator()
// 	fmt.Println(counter())
// 	fmt.Println(counter())

// 	c1 := circleArea(3.14)
// 	fmt.Println(c1(2))

// 	c2 := circleArea(3)
// 	fmt.Println(c2(2))
// }

/* 可変長引数 */

// func foo(params ...int) {
// 	fmt.Println(len(params), params)
// 	for _, param := range params {
// 		fmt.Println(param)
// 	}
// }

// func main() {
// 	foo(10, 20, 30)

// 	s := []int{1, 2, 3}
// 	foo(s...)
// }

// func main() {
// 	//Q1
// 	f := 1.11
// 	i := int(f)
// 	fmt.Printf("%T, %v\n", i, i)

// 	//Q2
// 	[5,6]

// 	//Q3
// 	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
// 	fmt.Printf("%T %v", m, m)
// }

/* IF */

// func main() {
// 	result := 10
// 	if result == 10 {
// 		fmt.Println("10")
// 	}
// 	if result2 := 20; result2 == 20 {
// 		fmt.Println("20")
// 	}
// }

/* For */

// func main() {
// 	l := []string{"python", "go", "java"}
// 	for i := 0; i < len(l); i++ {
// 		fmt.Println(i, l[i])
// 	}

// 	for i, v := range l {
// 		fmt.Println(i, v)
// 	}

// 	//iを使用しない場合
// 	for _, v := range l {
// 		fmt.Println(v)
// 	}

// 	m := map[string]int{"apple": 100, "banana": 200}

// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}

// 	//kを使用しない場合
// 	for _, v := range m {
// 		fmt.Println(v)
// 	}

// 	//vを使用しない場合
// 	for k := range m {
// 		fmt.Println(k)
// 	}
// }

/* switch */

// func main() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 11:
// 		fmt.Println("Morning")
// 	default:
// 		fmt.Println("Afternoon")
// 	}
// }

/* defer */

// func main() {
// 	fmt.Println("run")
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3)
// 	fmt.Println("success")
// }

/* err */
// func main() {
// 	file, err := os.Open("./main.go")
// 	defer file.Close()
// 	if err != nil {
// 		log.Fatal("error")
// 	}
// 	data := make([]byte, 100)
// 	count, err := file.Read(data)
// 	if err != nil {
// 		log.Fatal("error2")
// 	}
// 	fmt.Println(count, string(data))
// }

/* panic recover */

// func main() {
// 	defer func() {
// 		s := recover()
// 		fmt.Println(s)
// 	}()
// 	panic("panic")
// }

/* 演習 */

// func main() {
// 	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

// 	m := l[0]
// 	for _, l := range l {
// 		if m > l {
// 			m = l
// 		}
// 	}
// 	fmt.Println(m)
// }

// func main() {
// 	m := map[string]int{
// 		"apple":  200,
// 		"banana": 300,
// 		"grapes": 150,
// 		"orange": 80,
// 		"papaya": 500,
// 		"kiwi":   90,
// 	}

// 	var max int
// 	for _, v := range m {
// 		max += v
// 	}
// 	fmt.Println(max)
// }

/* ポインタ */

// func foo(x int) {
// 	x = 1
// }

// func bar(x *int) {
// 	*x = 1
// }

// func main() {
// 	var n int = 100
// 	foo(n)
// 	fmt.Println(n)
// 	bar(&n)
// 	fmt.Println(n)
// }

// func main() {
// 	var n int = 100
// 	fmt.Println(n)
// 	fmt.Println(&n)

// 	var p *int = &n
// 	fmt.Println(p)
// 	fmt.Println(*p)
// }

/* new */

// func main() {
// 	// var p *int = new(int)
// 	// fmt.Println(p) //0xc000014090
// 	// var p2 *int
// 	// fmt.Println(p2) //<nil>

// 	//スライス
// 	s := make([]int, 0)
// 	fmt.Printf("%T\n", s)

// 	//マップ
// 	m := make(map[string]int)
// 	fmt.Printf("%T\n", m)

// 	//ポインタ
// 	var p *int = new(int)
// 	fmt.Printf("%T\n", p)
// }

/* struct */
// type Vertex struct {
// 	X int
// 	Y int
// 	S string
// }

// func main() {
// 	// v1 := Vertex{X: 1, Y: 2}
// 	// fmt.Println(v1)
// 	// v2 := Vertex{1, 2, "test"}
// 	// fmt.Println(v2)

// 	// v3 := Vertex{}
// 	// fmt.Println(v3)

// 	// var v4 Vertex
// 	// fmt.Println(v4)

// 	// v5 := new(Vertex)
// 	// fmt.Println(v5)
// 	// fmt.Printf("%T %v\n", v5, v5)

// 	// v6 := &Vertex{}
// 	// fmt.Println(v6)
// 	// fmt.Printf("%T %v\n", v6, v6)

// 	m1 := make([]int, 0)
// 	fmt.Printf("%T %v\n", m1, m1)
// 	m2 := []int{}
// 	fmt.Printf("%T %v\n", m2, m2)
// }

/* レシーバ */
// type Vertex struct {
// 	X, Y int
// }

// // func Area(v Vertex) int {
// // 	return v.X * v.Y
// // }

// func (v Vertex) Area() int {
// 	return v.X * v.Y
// }

// func (v *Vertex) Scale(i int) {
// 	v.X = v.X * i
// 	v.Y = v.Y * i
// }

// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

// func main() {
// 	// v := Vertex{3, 4}
// 	v := New(3, 4)
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }

// /* Embedded */
// type Vertex struct {
// 	x, y int
// }

// func (v Vertex) Area() int {
// 	return v.x * v.y
// }

// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// type Vertex3D struct {
// 	Vertex
// 	z int
// }

// func (v Vertex3D) Area3D() int {
// 	return v.x * v.y * v.z
// }

// func (v *Vertex3D) Scale3D(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// 	v.z = v.z * i
// }

// func New(x, y, z int) *Vertex3D {
// 	return &Vertex3D{Vertex{x, y}, z}
// }

// func main() {
// 	v := New(3, 4, 5)
// 	v.Scale3D(10)
// 	// fmt.Println(v.Area())
// 	fmt.Println(v.Area3D())
// }

/* nonstruct */

// type MyInt int

// func (i MyInt) Double() int {
// 	return int(i * 2)
// }

/* interface */

// type Human interface {
// 	Say() string
// }

// type Dog struct {
// 	Name string
// }

// type Person struct {
// 	Name string
// }

// func (p *Person) Say() string {
// 	p.Name = "Mr." + p.Name
// 	fmt.Println(p.Name)
// 	return p.Name
// }

// func DriveCar(h Human) {
// 	if h.Say() == "Mr.Mike" {
// 		fmt.Println("Run")
// 	} else {
// 		fmt.Println("Get Out")
// 	}
// }

// func main() {
// 	var mike Human = &Person{"Mike"}
// 	DriveCar(mike)

// 	var x Human = &Person{"X"}
// 	DriveCar(x)

// 	//cannot use dog (type Dog) as type Human in argument to DriveCar
// 	// var dog Dog = &Dog{"Hachi"}
// 	// DriveCar(dog)
// }

/* assertion */

// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println(v * 2)
// 	case string:
// 		fmt.Println(v + "!")
// 	default:
// 		fmt.Printf("I don't know %T\n", v)
// 	}
// }

// func main() {
// 	do(10)
// 	do("Mike")
// 	do(true)
// }

/* stringer */
// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("My Name is %v", p.Name)
// }

// func main() {
// 	mike := Person{"Mike", 22}
// 	fmt.Println(mike)
// }

/* Error */

// type UserNotFound struct {
// 	UserName string
// }

// func (e *UserNotFound) Error() string {
// 	return fmt.Sprintf("User not found: %v", e.UserName)
// }

// func myFunc() error {
// 	ok := false
// 	if ok {
// 		return nil
// 	}
// 	return &UserNotFound{UserName: "Mike"}
// }

// func main() {
// 	if err := myFunc(); err != nil {
// 		fmt.Println(err)
// 	}
// }

/*
Q1. 以下に、7と表示されるメソッドを作成してください。
*/

// type Vertex struct {
// 	X, Y int
// }

// func (v Vertex) Plus() int {
// 	return v.X + v.Y
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Plus())
// }

/*
Q2 X is 3! Y is 4! と表示されるStringerを作成してください。
*/

// type Vertex struct {
// 	X, Y int
// }

// func (v Vertex) String() string {
// 	return fmt.Sprintf("X is %v!, Y is %v!", v.X, v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v)
// }

/* GoRoutine */
// func goroutine(s string, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {
// 		//time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// 	wg.Done()
// }

// func normal(s string) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go goroutine("world", &wg)
// 	normal("Hello")
// 	wg.Wait()
// }

/* channel */

// func gorouitne1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 		c <- sum
// 	}
// 	close(c)
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int)
// 	go gorouitne1(s, c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

/* Buffered Channels */

// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 100
// 	fmt.Println(len(ch))
// 	ch <- 200
// 	fmt.Println(len(ch))

// 	close(ch)
// 	for c := range ch {
// 		fmt.Println(c)
// 	}
// }

/* producer, consumer */

// func producer(ch chan int, i int) {
// 	ch <- i * 2
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		func() {
// 			defer wg.Done()
// 			fmt.Println("process", i*1000)
// 		}()
// 	}
// 	fmt.Println("#####################")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	//Producer
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go producer(ch, i)
// 	}

// 	//Consumer
// 	go consumer(ch, &wg)
// 	wg.Wait()
// 	close(ch)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Done")
// }

/* fan-out fan-in */
// func producer(first chan int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }

// func multi2(first <-chan int, second chan<- int) {
// 	defer close(second)
// 	for v := range first {
// 		second <- v * 2
// 	}
// }

// func multi4(second <-chan int, third chan<- int) {
// 	defer close(third)
// 	for v := range second {
// 		third <- v * 4
// 	}
// }

// func main() {
// 	first := make(chan int)
// 	second := make(chan int)
// 	third := make(chan int)

// 	go producer(first)
// 	go multi2(first, second)
// 	go multi4(second, third)

// 	for result := range third {
// 		fmt.Println(result)
// 	}
// }

/* select */

// func gorouitne1(ch chan string) {
// 	for {
// 		ch <- "packet from 1"
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func gorouitne2(ch chan int) {
// 	for {
// 		ch <- 100
// 		time.Sleep(3 * time.Second)
// 	}
// }

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan int)
// 	go gorouitne1(c1)
// 	go gorouitne2(c2)
// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

/* Default Selection */

// func main() {
// 	tick := time.Tick(100 * time.Millisecond)
// 	boom := time.After(500 * time.Millisecond)
// OuterLoop:
// 	for {
// 		select {
// 		case <-tick:
// 			fmt.Println("tick.")
// 		case <-boom:
// 			fmt.Println("BOOM!")
// 			break OuterLoop
// 		default:
// 			fmt.Println("    .")
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// 	fmt.Println("##############")
// }

/* sync.Mutex */
// type Counter struct {
// 	v   map[string]int
// 	mux sync.Mutex
// }

// //実体を書き換えたいのでポインタレシーバ
// func (c *Counter) Inc(key string) {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	c.v[key]++
// }

// func (c *Counter) Value(key string) int {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	return c.v[key]
// }

// func main() {
// 	// c := make(map[string]int)
// 	c := Counter{v: make(map[string]int)}
// 	go func() {
// 		for i := 0; i < 100000; i++ {
// 			// c["key"] += 1
// 			c.Inc("key")
// 		}
// 	}()
// 	go func() {
// 		for i := 0; i < 100000; i++ {
// 			// c["key"] += 1
// 			c.Inc("key")
// 		}
// 	}()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(c, (&c).Value("key"))
// }

/* 演習 */

func goroutine(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}

/*  */
