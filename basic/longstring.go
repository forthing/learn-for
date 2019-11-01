package main

import "fmt"

func lengthsOfsearchlongstring (Str string) int {
	lastOccurered := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(Str) {
		if lastI, ok := lastOccurered[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurered[ch] = i

	}
	return maxlength
}
func main() {
	fmt.Println(lengthsOfsearchlongstring("abcbcabba"))
	fmt.Println(lengthsOfsearchlongstring("bbbbbbb"))
	fmt.Println(lengthsOfsearchlongstring("pwwkew"))


}
