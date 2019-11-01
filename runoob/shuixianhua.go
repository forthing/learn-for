package main


import (
	"fmt"
	"math"
)

func main() {
	for num := 100; num <= 999; num++ {
		var i = num / 100
		var j = num / 10 % 10
		var k = num % 10
		if math.Pow(float64(i), 3) + math.Pow(float64(j), 3) + math.Pow(float64(k), 3) == float64(num) {
			fmt.Println(num)
		}
	}
}
