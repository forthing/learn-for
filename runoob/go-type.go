package main

import (
	"fmt"
//	"strings"
)

//包的引入可以采用：项目名/包名
//方法的调用采用 包名.方法名（）


// 一行代表一个语句结束，多个语句之间用；隔开
// 注释不会被编译，单行注释用// ，多行用/* 。。。 */ 隔开
var b bool = true
// 标识符第一个字符必须是字母或下划线而不能是数字。
// 这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量

func main() {
	var a int
    a,b := 2, 3
//	var b float32 = 3.5
//	var c, d int
	fmt.Println("goolge"+"runnoob")
	fmt.Println(a,b)
	var i int
	var f float64
//	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
//