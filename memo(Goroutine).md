### Goroutine

軽量のスレッド(並列処理)の事。
マルチプロセス、マルチスレッド、イベント駆動などの処理を暗黙的に Golang が実行してくれる。
並列化したい処理の前に go と付けるだけでよい。

```go
func goroutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go goroutine("world")
	normal("Hello")
}
```

> Hello
> world
> Hello
> world
> Hello
> world
> Hello
> world
> Hello

time.Sleep()を行わなかった場合

```go

func goroutine(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go goroutine("world")
	normal("Hello")
}
```

並列化した処理で出力されるはずの"world"が出力されなくなる。並列化したスレッドを生成して実行する前に処理が終わってしまった為。

> Hello
> Hello
> Hello
> Hello
> Hello

### sync.WaitGroup

goroutine の処理完了を待たずにプログラムが終了してしまうことを避ける為に使用する。wg.Add(1)で WaitGroup に 1 つの並列処理があることを伝える。
並列処理には WaitGroup(wg) のアドレスを伝え、処理が終わったら、wg.Done()を実行。
main()の最後には wg.Wait()を実行し、wg.Done()が完了するまでプログラムの終了をストップする。

```go
func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)
	normal("Hello")
	wg.Wait()
}
```

### channel

main()と並列化した処理(goroutine())はそれぞれ別々に実行されているため、通常のやり方ではデータのやり取りができない。
そこで channel を用いる。

goroutine1 で sum を計算した後、c <- sum で channel に sum を送信する。
sum が送信されるまで、<-c では処理を受信待ち状態でブロッキングしている。

```go
func gorouitne1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go gorouitne1(s, c)
	x := <-c
	fmt.Println(x)
}
```

### Buffered Channels

ch := make(chan int, 2)のように書くことで、channel に入れられる値の上限を決めることができる。
上限 2 の Buffered Channel に 3 つ目の値を入れようとすると、エラーが発生する。

```go
func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	// ch <- 300
	// fmt.Println(len(ch))
}
```

上限の channel から値を取り出すには、x := <-ch のように受信処理を行う。

```go
func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	x := <-ch
	fmt.Println(x)

	ch <- 300
	fmt.Println(len(ch))
}
```

ループ処理で channel から値を取り出す場合には、x := <-ch ではなく、以下のように行う。

この時、channel を先にクローズしておかないと、ループが終わらずにエラーとなる為、range で channel を呼び出す際には、close()で channel の終了を宣言する必要がある。

```go
func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
```

channel のクローズは、並列処理内で実行してもよい。

```go

func gorouitne1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go gorouitne1(s, c)
	for i := range c {
		fmt.Println(i)
	}
}
```

### producer,comsumer

goroutine の producer の処理で実行した結果を channel に入れて、goroutine の consumer で受け取って処理をする。

先ほど、main()の中で range でチャネルから値を取り出す前に、close()を書いたが、今回は goroutine である consumer()の内部で、10 回のループが終わった後も、外側からチャネルが渡された時に備えて range から channel を取ろうとスタンバイしている状態。

```go
func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i*1000)
		wg.Done()
	}
	fmt.Println("#####################")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	//Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	//Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
```

close(ch)を書くことで、consumer()での range での待機待ち状態が終了し、 fmt.Println("#####################")が実行される。
その 2 秒後に fmt.Println("Done")が実行される。

もし、close(ch)を書かなかった場合は、consumer()での range での待機待ち状態が終了しないため、fmt.Println("#####################")が実行されず、
2 秒後に fmt.Println("Done")だけが実行される。

> process 0
> process 2000
> process 4000
> process 6000
> process 8000
> process 10000
> process 12000
> process 14000
> process 16000
> process 18000
> #####################
> Done

また、consumer()内の処理である fmt.Println("process", i\*1000)が正常に終了しなかった場合に、wg.Done()も実行されないリスクが存在する。
その為、range ループ内で inner function を使い、wg.Done()を遅延実行することでリスクを防ぐことができる。

```go
func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
	fmt.Println("#####################")
}
```

### fan-out,fan-in

始めの goroutine が処理をして channel に入れたものに対して、別の goroutine で再度処理をして、また別の channel に入れることを、fan-out、fan-in という。

multi2()では、first を受信、second に送信を行っているが、
multi2(first <-chan int, second chan<- int)のように<-chan、chan<-を使った書き方もできる。

```go
func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for v := range first {
		second <- v * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for v := range second {
		third <- v * 4
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi4(second, third)

	for result := range third {
		fmt.Println(result)
	}
}
```

### select を用いた channel の受信

main()で異なる goroutine から channel を通じて異なる情報を取得する。
select を使うことで、複数の goroutine から同時に値を受け取れる。

```go
func gorouitne1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func gorouitne2(ch chan int) {
	for {
		ch <- 100
		time.Sleep(3 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan int)
	go gorouitne1(c1)
	go gorouitne2(c2)
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
```

### Default Selection

https://tour.golang.org/concurrency/6

for ループが完了した後に完了メッセージ("##############")を出したいが、BOOM!後の return では main()を抜けてしまう。
return を break に変えると、for ループが終了せずに無限にループしてしまう。

```go

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("##############") //unreachable code
}
```

> ..tick. ..tick. ..tick. ..tick. ..tick. BOOM!

悪いやり方(フラグで管理する。)

```go
	for {
		isBreak := false
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			isBreak = true
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
		if isBreak {
			break
		}
	}
	fmt.Println("##############")
```

良いやり方

```go
OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("##############")
```

### sync.Mutex

複数の goroutine から一つの channel の値を同時に書き換えようとすると、何度か実行した結果エラーが発生する。

fatal error: concurrent map writes

```go
func main() {
	c := make(map[string]int)
	go func() {
		for i := 0; i < 100000; i++ {
			c["key"] += 1
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			c["key"] += 1
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c["key"])
}

```

goroutine 間で channel の受け渡しが出来ればよいが、それができない場合には、sync.Mutex を使用する。

c.mux.Lock()を行うことで、複数の goroutine が同時に channel への書き込みを行う事ができなくなる。

```go
type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

//実体を書き換えたいのでポインタレシーバ
func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	// c := make(map[string]int)
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 100000; i++ {
			// c["key"] += 1
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			// c["key"] += 1
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, (&c).Value("key"))
}
```
