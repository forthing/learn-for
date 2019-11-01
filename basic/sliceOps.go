package main

import "fmt"

func printSlice ( s []int ) {
	fmt.Printf("%v, len = %d, cap = %d\n",s, len(s), cap(s))
}

func main() {

	fmt.Println("creating slice")
	var s []int //zero value for slice is nil
	fmt.Println(s)
	for i := 0; i < 10; i++ {
//		printSlice(s)
		s = append(s, 2 * i + 1)
	}
//	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("deleting elements from slice")
	//s2[:3] + s2[4:]
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Poping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(front, tail)
	fmt.Println(s2)
	printSlice(s2)

}
