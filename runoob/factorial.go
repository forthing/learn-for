package main

import "fmt"

func factorial(n int) (result int){
	if n == 0{
		return 1
	} else {
		result = n * factorial(n - 1)
		return result
	}
}

func main() {
	var i int = 5
	fmt.Printf("%d factorial is %d", i, factorial(int(i)))
	
}
