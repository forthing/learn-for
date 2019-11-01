> :memo:该笔记主要记录GO语言相应难点及重点学习，初步的语法不做相应介绍。
***

#### 1、逐行读取文件内容

``` go
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
```
`:=` 可以对函数内变量进行定义和赋值，go语言支持多返回值，一版函数加`err`进行判断
#### 2、函数调用函数示例
```go
// define apply fuction
func apply(op func(int, int) int, a, b int ) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args " + "(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func main() {
	fmt.Println(apply(
		func (a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
}
```
将函数作为参数进行调用，并读取了函数的名称。
#### 3、slice重要操作
```go
	fmt.Println("extending slice")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
```
slice的实现有ptr、len、cap三个参数决定，尽管超过范围了，依然能够cap参数向后扩展取到相应值。
#### 4、rune相当于go的char
* 使用`range`可以遍历pos, rune对
* 使用utf8.RuneCountInString 获得字符数量
* len获得是字节长度，[]byte 获得是字节，[]rune 获得是字符
#### 5、结构体的创建
```go
root := TreeNode{value: 3}
root.Left = &TreeNode{}
root.Right = &TreeNode(nil, nil, 5)
root.Right.Left = new(TreeNode)
```
为结构体定义方法，调用时编译器判别指针或者值
> func (node *TreeNode) print( ) {}
> *值/指针接受者*  均可以接受值/指针
#### 6、删除`map`中的key-value对，并加以判断
```go
	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
```
#### 7、函数和闭包
 > 函数参数、变量、返回值都可以是函数
 > 函数体有局部变量和自由变量
 > 函数返回不仅是一个函数，是一个闭包
```go
func adder() func(int) int {
	sum := 0
	return func( v int) int {
		sum += v
		return sum
	}
}
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
```
#### 8、往文件中写入斐波那契数列
```go
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
```
> `defer`让我们不用担心文件关闭问题，用**bufio**写可以加快速度
#### 9、错误处理
```go
	file, err := os.Open("abc.txt")
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError.Err)
		} else {
			fmt.Println("Unknown error", err)
		}
	}
```
#### 10、goroutine
```go
func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello from " + "goroutine %d\n", i)
			}
		}(i) // 匿名函数
	}
	time.Sleep(time.Millisecond)
}
```
> 1毫秒内1000个协程不一定都能执行
> 协程是非抢占式多任务处理，由协程主动交出控制权。
> go里面是编译器/解释器/虚拟机层面的多任务。调度器决定多个协程在一个或多个线程中执行
> runtime.Gosched() 可以交出控制权
> 可以使用-race检测数据冲突，I/O，select，channel都可以进行切换。