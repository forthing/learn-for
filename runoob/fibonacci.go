package main

import	"fmt"

const NUM  = 10

func fibonacci(n int) int{
	if n< 2{
		return n
	}
	return fibonacci(n - 2) + fibonacci(n - 1)
}
func main() {
	var i int
	for i = 0; i < NUM; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	
}
