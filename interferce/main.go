package main

import (
	"fmt"
	"learngo/interferce/mock"
	real2 "learngo/interferce/real"
)

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retiever
	r = mock.Retiever{"this is a fake imooc.com"}
	r = real2.Retiever{}
	fmt.Println(download(r))
}
