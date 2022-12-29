package test

// "sync"

// func test(x int, y int) (result int) {
// 	// var x int = 1
// 	// var y int = 2
// 	// x := 1
// 	// y := 2
// 	// const x int = 1
// 	// const y int = 2
// 	result = x + y
// 	return result
// }

// func incrementGenerator() func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// type Vertex struct{
// 	x int
// 	y int
// }
// //継承みたいなもの
// type Vertex3D struct{
// 	Vertex
// 	z int
// }

// type Person struct {
// 	Name string
// }

// func (p *Person) say() {
// 	fmt.Println(p.Name)
// }

// type Human interface {
// 	say()
// }

// func (v Vertex) Area() int {
// 	return v.X * v.Y
// }

// *をつけることでポインタのアドレスの中身の実態を見に行って書き換えることができる
// func (v *Vertex) Scale(i int) {
// 	v.X = v.X * i
// 	v.Y = v.Y * i
// }

// func New(x,y int) *Vertex {
// 	return &Vertex{x,y}
// }

// タイプアサーション
// i.(int)
// switch v:= i.(type) {
// case int:
// case string:
// case default:

// }

// func goroutine(j int, wait *sync.WaitGroup) {
// 	for i := 0; i < j; i++ {
// 		fmt.Println("並列:" , i)
// 	}
// 	wait.Done()
// }

// func noGoroutine(j int) {
// 	for i := 0; i < j; i++ {
// 		fmt.Println("直列:" , i)
// 	}
// }

// func plus1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// 	close(c)
// }

// func plus2(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func producer(ch chan int, i int) {
// 	ch <- i + 3
// }

// func customer(ch chan int, wg *sync.WaitGroup) {
// 	//closeしないと次のchanを待っている
// 	for i := range ch {
// 		fmt.Println("出力", i)
// 		wg.Done()
// 	}
// }

// func producer(ch chan<- int) {
// 	defer close(ch)
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		ch <- i
// 	}
// }

// func select1(ch chan string) {
// 	for {
// 		ch <- "aaaaaa"
// 		time.Sleep(3 * time.Second)
// 	}
// }
// func select2(ch chan string) {
// 	for {
// 		ch <- "bbbbb"
// 		time.Sleep(1 * time.Second)
// 	}
// }

// const (
// 	c1 = iota
// 	c2 = iota
// 	c3 = iota
// )

func main() {

	// resp, _ := http.Get("https://google.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()
	// ioutil.ReadFile()
	// c1 := 2
	// fmt.Println(c1, c2, c3)
	// i := []int{1, 5, 4, 6, 7, 2, 32}
	// sort.Ints(i)
	// fmt.Println(i)
	// p := []struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	{"aaa", 20},
	// 	{"ccc", 30},
	// 	{"baa", 20},
	// }
	// sort.Slice(p, func(i, j int) bool {
	// 	return p[i].Name < p[j].Name
	// })
	// fmt.Println(p)
	// match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	// fmt.Println(match)

	// r := regexp.MustCompile("a([a-z]+)e")
	// ms := r.MatchString("apple")
	// fmt.Println(ms)

	// r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	// fs := r2.FindString("/view/test")
	// fmt.Println(fs)

	// fss := r2.FindStringSubmatch("/view/test")
	// fmt.Println(fss, fss[0], fss[1])

	// t := time.Now()
	// tFormat := t.Format(time.RFC3339)
	// fmt.Println(tFormat)
	// 	c1 := make(chan string)
	// 	c2 := make(chan string)

	// 	go select1(c1)
	// 	go select2(c2)

	// loopbreak:
	// 	for {
	// 		select {
	// 		case msg1 := <-c1:
	// 			fmt.Println(msg1)
	// 			break loopbreak
	// 		case msg2 := <-c2:
	// 			fmt.Println(msg2)
	// 		}
	// 	}
	// var wg sync.WaitGroup
	// ch := make(chan int)

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go producer(ch, i)
	// }

	// go customer(ch, &wg)
	// wg.Wait()
	// close(ch)
	// first := make(chan int)
	// second := make(chan int)
	// third := make(chan int)
	// go producer(first)
	// s1 := []int{1, 2, 3, 4, 5}
	// s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// c := make(chan int, 2)

	// go plus1(s1, c)
	// go plus2(s2, c)
	//　この行はcが入ってくるまで待っている
	// x := <-c
	// fmt.Println(x)
	//　この行はcが入ってくるまで待っている
	// y := <-c
	// fmt.Println(y)
	// fmt.Println(c)

	// for i := range c {
	// 	fmt.Println(i)
	// }
	// var wait sync.WaitGroup
	// wait.Add(1)
	// go goroutine(5, &wait)
	// noGoroutine(5)
	// wait.Wait()
	// v := Vertex{3,4}
	// fmt.Println(v.Scale(100))
	// v := New{3,4}
	// v.Scale(100)
	// fmt.Println(v.Area())

	// var ohya Human = &Person{"ryota"}
	// ohya.say()
	// fmt.Println("Hello World")
	// result := test(1,2)
	// fmt.Println(result)

	// increment := incrementGenerator()

	// fmt.Println(increment())
	// fmt.Println(increment())
	// fmt.Println(increment())
	// fmt.Println(increment())

	// x := map[string]int {
	// 	"aaa" : 20,
	// 	"bbb" : 20,
	// }
	// f, ok := x["ddd"]
	// fmt.Println(f,ok)
	// d, ok2 := x["aaa"]
	// fmt.Println(d,ok2)

	// y := make([]int, 3, 5)
	// fmt.Println(len(y), cap(y), y)

	// z := y[0]
	// for  z <10 {
	// 	z ++
	// 	fmt.Println("aaa")
	// }

	// func (x int) {
	// 	fmt.Println(x)
	// }(1)

	// l := []string{"a", "b", "c"}

	// for _, v := range l {
	// 	fmt.Println(v)
	// }

	// switch x := test(); x {
	// case "a":

	// }
	// t := time.Now()
	// var x int = 1
	// var y *int = &x

	// fmt.Println(x)
	// fmt.Println(&*&*y)

	// defer fmt.Println("aaaaa")

	// fmt.Println("aaaaaaaaa")
	// x1 := Vertex{1,2}
	// log.Println(x1)
	// log.Println("aaa")
}
