package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2

	defer calc('';1'';, a, calc('&quot';10'';, a, b))
	a = 0
	defer calc(&quot;2&quot;, a, calc(&quot;20&quot;, a, b))
	b = 1
}