package main

import "fmt"

func main() {
	var line int  = 10
	for i := 0; i < line; i++ {
		for j := 0; j < line - i - 1; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < 2 * i + 1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	
}
