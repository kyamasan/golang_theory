package main

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

/*  */
/*  */
/*  */
/*  */
/*  */
/*  */
/*  */
/*  */
/*  */
/*  */
