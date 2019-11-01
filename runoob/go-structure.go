package main

fmt
// . represent simply import ,can use Println than fmt.Println

const PI = 3.14

var name = "gopher"

type newType int

type gopher struct{}

type golang interface {

}

func main() {
	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	fmt.Println("a", "b", 1, 2, 3, "c", "d")
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)
	// ab1 2 3cd
	// a b 1 2 3 c d
	// ab 1 2 3 cd
}
//