package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i ++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}

}
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	/*
	file, err := os.Open("abc.txt")
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError.Err)
		} else {
			fmt.Println("Unknown error", err)
		}
	}

	 */
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
//func main() {
	//tryDefer()
	//writeFile("fib.txt")
	
//}


