package main

import (
	"fmt"
	"io/ioutil"
)

func eval (a ,b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)

	}
	return result
}
func main() {
	const filename = "abc.txt"
	//  if条件里可以赋值
	if contens, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contens)
	}
	//fmt.Println(contents)
	fmt.Println(grade(0), grade(59), grade(60), grade(101))
}

numbers := [6]int{1, 2, 3, 4, 5, 6}
for i = 0; i < lens(numbers); i++ {
	fmt.Println(numbers[i])
}
sum := 0
for _, v := range numbers {
	sum += v
}

func grade (score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}