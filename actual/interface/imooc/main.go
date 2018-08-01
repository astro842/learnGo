package main

import (
	"fmt"
	"learnGo/actual/interface/imooc/real"
	"learnGo/actual/interface/imooc/mock"
)

//接口 使用者 定义接口里的方法

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"ttt"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
