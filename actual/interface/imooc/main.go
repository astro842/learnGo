package main

import (
	"fmt"
	"learnGo/actual/interface/imooc/mock"
	"learnGo/actual/interface/imooc/real"
)

//定义接口 和接口里的方法
type Retriever interface {
	Get(url string) string
}
//接口 使用者
func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.AAA{"tt12t"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
