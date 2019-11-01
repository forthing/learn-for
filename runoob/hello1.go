package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var bb = true


func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue(){
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableShorter(){
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler( ){
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi) + 1)
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt( float64(a * a + b * b)))
	fmt.Println(c)


}

func consts (){
	const(
		filename = "abc.txt"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt( a * a + b * b))
	fmt.Println(c)
	fmt.Println(filename, c)
}

func enums() {
	const(
		cpp = iota
		_
		java
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript , python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
// { 不能单独放一行
func main() {
	// fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n
//	fmt.Println("hello,world")
	//标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
	// 那么使用这种形式的标识符的对象就可以被外部包的代码所使用
//	variableZeroValue()
//	variableInitialValue()
//	variableShorter()
//	euler()
//	triangle()
	consts()
	enums()
	fmt.Println(1 << 10)
}