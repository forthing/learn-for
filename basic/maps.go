package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "xuzhiyong",
		"site": "golanag",
	}
	fmt.Println(m)
	var m3 map[string]int
	fmt.Println(m3)
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("getting values")
	zyongname, ok := m["name"]
	fmt.Println(zyongname, ok)

	fmt.Println("deleting elements")
	name, ok := m["name"]

}
