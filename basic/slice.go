package main

import "fmt"

func upadeteslice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println("array[2:6]", s)
	fmt.Println("arr[:]", arr[:])

	s1 := arr[2:]
	upadeteslice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	s2 := arr[:]
	upadeteslice(s2)
	fmt.Println(s2)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("extending slice")
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr = ", arr)
	s1 = arr[2:6]

	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	s2 = s1[3:5]
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))
}
